package workerpool

type WorkerPool struct {
	Workers []*Worker
}

func New(workers []*Worker) *WorkerPool {
	return &WorkerPool{
		Workers: workers,
	}
}

func (p *WorkerPool) SelectWorker() *Worker {
	// Round-robin or least-load logic
	selected := p.Workers[0]
	for _, w := range p.Workers {
		if w.Load < selected.Load {
			selected = w
		}
	}
	selected.Load++
	return selected
}
