package main

import (
	"context"
	"encoding/json"
	"flint/proto" // replace with your proto package
	"fmt"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type WorkerServer struct {
	proto.UnimplementedWorkerServiceServer
}

func parseNumber(val interface{}) (float64, error) {
	switch v := val.(type) {
	case float64:
		return v, nil
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("invalid type")
	}
}

// Custom Worker Logic
func (s *WorkerServer) ExecuteTask(ctx context.Context, req *proto.TaskRequest) (*proto.TaskResult, error) {
	fmt.Printf("Executing task %s with payload %s\n", req.TaskId, req.Payload)

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

	aVal, err1 := parseNumber(payload["A"])
	bVal, err2 := parseNumber(payload["B"])
	if err1 != nil || err2 != nil {
		// handle parse error
	}
	sum := aVal + bVal

	result := fmt.Sprintf("Sum of %.2f and %.2f is %.2f", payload["A"], payload["B"], sum)

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
