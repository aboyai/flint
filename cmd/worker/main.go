package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flint/proto" // replace with your proto package
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type WorkerServer struct {
	proto.UnimplementedWorkerServiceServer
}

func runTaskAuto(scriptPath, venvPath string, payload map[string]interface{}) (string, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to serialize payload: %v", err)
	}

	// Determine file extension
	scriptPath = filepath.FromSlash(scriptPath)

	if scriptPath == "" {
		log.Fatal("Error: TASK_SCRIPT is empty")
	}

	ext := strings.ToLower(filepath.Ext(scriptPath))
	fmt.Println("Script Path:", scriptPath)
	fmt.Println("Detected Ext:", ext)

	var cmd *exec.Cmd

	switch ext {
	case ".py":
		// Python execution
		var pythonExe string
		if runtime.GOOS == "windows" {
			pythonExe = filepath.Join(venvPath, "Scripts", "python.exe")
		} else {
			pythonExe = filepath.Join(venvPath, "bin", "python")
		}
		println("Command: %s", cmd)
		cmd = exec.Command(pythonExe, scriptPath)

	case ".exe", "":
		// Native executable (Windows .exe or Linux ELF)
		cmd = exec.Command(scriptPath)

	default:
		return "", fmt.Errorf("unsupported script type: %s", ext)
	}
	cmd.Stdin = bytes.NewBuffer(payloadBytes)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	cmd.Env = os.Environ()

	err = cmd.Run()
	result := strings.TrimSpace(out.String())

	if err != nil {
		return result, fmt.Errorf("execution error [%s]: %v", ext, err)
	}
	return result, nil

}

// Custom Worker Logic
func (s *WorkerServer) ExecuteTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResult, error) {
	fmt.Printf("Executing task %s with payload %s\n", req.TaskId, req.Payload)
	// Load env variables
	_ = godotenv.Load(".env")

	scriptPath := os.Getenv("TASK_SCRIPT") // Can be .py or .exe
	venvPath := os.Getenv("PY_VENV")       // Used only if script is Python

	var payload map[string]interface{}
	err := json.Unmarshal([]byte(req.Payload), &payload)
	if err != nil {
		fmt.Printf("Invalid JSON payload: %v\n", err)
		return &proto.TaskResult{
			TaskId:     req.TaskId,
			DagId:      req.DagId,
			Status:     "failure",
			ResultData: "Invalid JSON payload",
		}, nil
	}

	result, err := runTaskAuto(scriptPath, venvPath, payload)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Task Result: %s\n", result)
	}

	fmt.Printf("Result: %s\n", result)

	return &proto.TaskResult{
		TaskId:     req.TaskId,
		DagId:      req.DagId,
		Status:     "completed",
		ResultData: result,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	fmt.Println("Listener:", lis)

	grpcServer := grpc.NewServer()
	proto.RegisterWorkerServiceServer(grpcServer, &WorkerServer{})
	grpcServer.Serve(lis)

}
