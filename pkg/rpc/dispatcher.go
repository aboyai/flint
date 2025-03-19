package rpc

import (
	"context"
	"fmt"

	"flint/pkg/dag"
	"flint/pkg/workerpool"
	"flint/proto" // Replace with actual proto package path
)

func DispatchTask(t *dag.Task, w *workerpool.Worker) (*proto.TaskResult, error) {
	// Connect to worker via gRPC
	//address := fmt.Sprintf("%s:50051", w.Address) // Ensure worker gRPC port is correct
	/*conn, err := grpc.Dial(w.Address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("gRPC dial failed: %v", err)
	}
	defer conn.Close()*/

	client := proto.NewWorkerServiceClient(w.Conn)

	// Build request
	req := &proto.TaskRequest{
		TaskId:  t.ID,
		DagId:   t.DAGID, // Replace with t.DAGID if you have it
		Payload: t.Payload,
	}

	// Execute task via gRPC
	ctx := context.Background() // Wait indefinitely
	result, err := client.ExecuteTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("gRPC ExecuteTask failed: %v", err)
	}

	return result, nil
}

/*func DispatchTask(t *dag.Task, w *workerpool.Worker) error {
	fmt.Printf("Simulating execution of task %s on worker %s\n", t.ID, w.ID)
	time.Sleep(2 * time.Second)
	return nil // Force success
}*/

/*func DispatchTask(t *dag.Task, w *workerpool.Worker) error {
	conn, err := grpc.Dial(w.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewWorkerServiceClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	//ctx := context.Background()
	ctx, _ := context.WithCancel(context.Background())
	//defer cancel() // Call this when you want to cancel manually

	//defer cancel()

	_, err = client.ExecuteTask(ctx, &proto.TaskRequest{
		TaskId:  t.ID,
		DagId:   "dag1",
		Payload: t.Payload,
	})

	return err
}
*/
