// Package task provides functionality to run function periodically.
package task

import (
	"context"
	"sync"
	"time"
)

// TasksGroup is a group of periodically run tasks,
type TasksGroup struct {
	wg sync.WaitGroup
}

// NewTasksGroup creates new instance of TaskGroup type.
func NewTasksGroup() *TasksGroup {
	return &TasksGroup{}
}

// WaitAll waits all tasks to be finished.
func (tg *TasksGroup) WaitAll() {
	tg.wg.Wait()
}

// Start starts function to run periodically with specified interval.
func (tg *TasksGroup) Start(ctx context.Context, interval time.Duration, task func()) {
	tg.wg.Add(1)

	go func() {
		defer tg.wg.Done()

		task()

		for {
			ticker := time.NewTicker(interval)

			select {
			case <-ticker.C:
				task()
			case <-ctx.Done():
				return
			}
		}
	}()
}
