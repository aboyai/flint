package workerpool

import (
	"context"
	"fmt"
	"log"
	"time"

	"flint/proto" // Replace with your proto import

	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Worker struct {
	ID      string
	Address string // IP
	Port    int
	Client  proto.WorkerServiceClient
	Conn    *grpc.ClientConn
	Load    int
}

func DiscoverWorkers() []*Worker {
	client, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		log.Fatalf("Consul client error: %v", err)
	}

	services, _, err := client.Health().Service("worker", "", true, nil)
	if err != nil {
		log.Fatalf("Consul health query failed: %v", err)
	}

	var workers []*Worker
	for _, s := range services {
		address := s.Service.Address
		port := s.Service.Port
		log.Printf("[Consul] Discovered worker %s at %s:%d", s.Service.ID, address, port)

		w, err := NewWorker(s.Service.ID, address, port)
		if err != nil {
			log.Printf("[Warning] Failed to connect to worker %s: %v", s.Service.ID, err)
			continue
		}
		workers = append(workers, w)
	}
	return workers
}

// NewWorker creates a Worker and establishes a gRPC connection
func NewWorker(id, address string, port int) (*Worker, error) {
	target := fmt.Sprintf("%s:%d", address, port)

	// Add gRPC dial timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to dial worker %s: %w", id, err)
	}

	client := proto.NewWorkerServiceClient(conn)

	return &Worker{
		ID:      id,
		Address: address,
		Port:    port,
		Client:  client,
		Conn:    conn,
	}, nil
}

// Close cleans up gRPC connection
func (w *Worker) Close() {
	if w.Conn != nil {
		w.Conn.Close()
	}
}
