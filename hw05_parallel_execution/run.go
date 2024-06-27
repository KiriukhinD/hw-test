package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

var ErrTooManyTasks = errors.New("too many tasks for the given error limit")

type Task func() error

func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	if m > len(tasks) {
		return ErrTooManyTasks
	}

	var wg sync.WaitGroup
	var errorCount int32
	var taskCount int32

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				taskIndex := int(atomic.AddInt32(&taskCount, 1)) - 1
				if taskIndex >= len(tasks) {
					return
				}

				if atomic.LoadInt32(&errorCount) >= int32(m) {
					return
				}

				err := tasks[taskIndex]()
				if err != nil {
					atomic.AddInt32(&errorCount, 1)
					if atomic.LoadInt32(&errorCount) >= int32(m) {
						// If error limit exceeded, signal other goroutines to stop
						return
					}
				}
			}
		}()
	}

	wg.Wait()

	if atomic.LoadInt32(&errorCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
