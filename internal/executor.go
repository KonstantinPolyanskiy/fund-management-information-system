package internal

import (
	"context"
	"sync"
)

type ExecutorProcess struct {
	Ctx       context.Context
	wg        *sync.WaitGroup
	Processes []func()
}

func (e *ExecutorProcess) Exec() {
	defer e.wg.Done()
	for _, p := range e.Processes {
		go p()
	}
}
func (e *ExecutorProcess) Add(process func()) *ExecutorProcess {
	e.wg.Add(1)
	e.Processes = append(e.Processes, func() {
		process()
	})
	return e
}

func NewExecutorProcess(wg *sync.WaitGroup) *ExecutorProcess {
	return &ExecutorProcess{
		Ctx:       context.Background(),
		wg:        wg,
		Processes: nil,
	}
}
