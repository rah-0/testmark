package main

import (
	"sync"
	"time"
)

type Result struct {
	TotalIterations int
	Elapsed         time.Duration
	ItersPerSec     float64
}

type Function func(cfg *Config, countFn func(int))

func RunMark(cfg *Config, fn Function) Result {
	cfg = cfg.New()

	var wg sync.WaitGroup
	var mu sync.Mutex
	total := 0
	start := time.Now()

	countFn := func(n int) {
		mu.Lock()
		total += n
		mu.Unlock()
	}

	for i := 0; i < cfg.Goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fn(cfg, countFn)
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)

	return Result{
		TotalIterations: total,
		Elapsed:         elapsed,
		ItersPerSec:     float64(total) / elapsed.Seconds(),
	}
}
