package main

import (
	"flint/pkg/dag"
	"flint/pkg/rpc"
	"flint/pkg/workerpool"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	myDAG := dag.DAG{ID: "dag1", Tasks: loadTasks()}

	w1, err := workerpool.NewWorker("worker-1", "127.0.0.1", 50051)
	if err != nil {
		log.Fatalf("Failed to initialize worker: %v", err)
	}

	workers := []*workerpool.Worker{w1}

	if len(workers) == 0 {
		log.Fatal("[Startup] No workers hardcoded — exiting.")
	}

	workerPool := workerpool.New(workers)

	// Graceful shutdown handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("[Shutdown] Closing worker connections...")
		for _, w := range workers {
			w.Close()
		}
		os.Exit(0)
	}()

	var wg sync.WaitGroup

	for {
		readyTasks := myDAG.GetReadyTasks()
		if len(readyTasks) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}

		for _, task := range readyTasks {
			worker := workerPool.SelectWorker()
			if worker == nil {
				log.Println("[Dispatch] No available worker. Skipping task dispatch.")
				continue
			}

			wg.Add(1)
			go func(t *dag.Task, w *workerpool.Worker) {
				defer wg.Done()

				result, err := rpc.DispatchTask(t, w)
				if err != nil {
					t.Status = "failed"
					log.Printf("[Dispatch] Task %s failed ❌: %v", t.ID, err)
					return
				}

				t.Status = result.Status
				t.Result = result.ResultData

				if t.Status == "failed" {
					log.Printf("[Dispatch] Task %s status: %s ❌", t.ID, t.Status)
				} else {
					log.Printf("[Dispatch] Task %s status: %s ✅", t.ID, t.Status)
				}

			}(task, worker)

			time.Sleep(1 * time.Second) // Optional throttling
		}

		// Print all task statuses dynamically
		fmt.Println("Task statuses:")
		for id, task := range myDAG.Tasks {
			fmt.Printf(" - %s: %s\n", id, task.Status)
		}

		wg.Wait()
	}
}

func loadTasks() map[string]*dag.Task {
	return map[string]*dag.Task{
		"t1": {ID: "t1", DAGID: "dag1", Payload: `{"A": "100", "B": "90"}`, DependsOn: []string{}, Status: "pending"},
		"t2": {ID: "t2", DAGID: "dag1", Payload: `{"A": "500", "B": "20"}`, DependsOn: []string{}, Status: "pending"},
	}
}
