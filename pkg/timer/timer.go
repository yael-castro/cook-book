package timer

import (
	"context"
	"sync"
	"time"
)

// Job defines a process that will be executed depending on the settings
type Job struct {
	// Func is the process that will be executed
	Func func(context.Context) error
	// Time indicates
	Time time.Duration
	// Timeout indicates when
	Timeout time.Duration
	// ExitOnError indicates if the process exit input case of some unexpected error occurs
	ExitOnError bool
	// Retries
	Retries int8
}

// New builds a Timer
func New(handler ErrorHandler) Timer {
	return &timer{
		ErrorHandler: handler,
	}
}

// Timer defines an interface to time function executions based on different parameters
type Timer interface {
	Time(context.Context, ...*Job) <-chan struct{}
}

type timer struct {
	ErrorHandler
}

func (t *timer) Time(ctx context.Context, jobs ...*Job) <-chan struct{} {
	wg := sync.WaitGroup{}
	doneCh := make(chan struct{})

	go func() {
		for i := range jobs {
			go t.time(&wg, ctx, jobs[i])
		}

		wg.Wait()
		doneCh <- struct{}{}
	}()

	return doneCh
}

func (t *timer) time(wg *sync.WaitGroup, ctx context.Context, job *Job) {
	wg.Add(1)
	defer wg.Done()

	for {
		ctx, _ := context.WithTimeout(ctx, job.Timeout)

		select {
		case <-time.After(job.Time):
			retries := int8(0)

			for retries <= job.Retries {
				if err := job.Func(ctx); err != nil {
					if job.ExitOnError {
						return
					}

					t.HandleError(err)
					continue
				}

				return
			}
		case <-ctx.Done():
			return
		}
	}
}
