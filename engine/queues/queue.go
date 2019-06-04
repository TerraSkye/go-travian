package queues

import (
	"fmt"
	ycq "github.com/jetbasrawi/go.cqrs"
	"time"
)

type Queue struct {
	dispatcher ycq.Dispatcher
	containers map[string]TaskContainer
}

type TaskContainer struct {
	tasks map[string]Task
	//can cancel all pending tasks.
	status chan bool
}

type Task struct {
	ContainerID string
	ID          string
	status      chan bool
	// time to process
	Ttp time.Duration
}

type Interval struct {
	Task
	Quantity int
}

// In memory Queue for handling actions
func NewQueue(dispatcher ycq.Dispatcher) *Queue {
	queue := Queue{
		dispatcher,
		map[string]TaskContainer{},
	}
	return &queue
}

func (queue *Queue) action(task Task, status chan bool) {
	select {
	case <-status:
		//pending task is canceld
		fmt.Print("Canceled building")
		break
	case <-time.After(task.Ttp * time.Millisecond):
		//task completed
		fmt.Print("Completed building")
		break

	}
}

func (queue *Queue) interval(task Interval, status chan bool, abort func(task2 Interval), complete func(task2 Interval), tick func(task2 Interval)) {
	ticker := time.NewTicker(task.Ttp * time.Millisecond)

	select {
	case <-status:
		abort(task)
		break
	case <-ticker.C:
		tick(task)
		//todo add quantity
	case <-time.After(task.Ttp * time.Millisecond):
		complete(task)
		break
	}

	fmt.Print("Task ended")
}
