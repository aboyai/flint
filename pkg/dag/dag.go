package dag

type Task struct {
	ID        string
	DAGID     string
	Payload   string
	DependsOn []string
	Status    string // pending, running, completed
	Result    string // Add this field
}

type DAG struct {
	ID    string
	Tasks map[string]*Task
}

func (d *DAG) GetReadyTasks() []*Task {
	ready := []*Task{}
	for _, task := range d.Tasks {
		if task.Status == "pending" && allDepsCompleted(d, task) {
			ready = append(ready, task)
		}
	}
	return ready
}

func allDepsCompleted(d *DAG, t *Task) bool {
	for _, depID := range t.DependsOn {
		depTask, exists := d.Tasks[depID]
		if !exists || depTask.Status != "completed" {
			return false
		}
	}
	return true
}
