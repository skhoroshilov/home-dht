package task

import (
	"context"
	"sync"
	"time"
)

var wg sync.WaitGroup

func WaitAll() {
	wg.Wait()
}

func Start(ctx context.Context, interval time.Duration, task func()) {
	wg.Add(1)

	go func() {
		defer wg.Done()

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
