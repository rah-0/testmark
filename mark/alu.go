package main

import (
	"time"
)

func BenchmarkInt64(cfg *Config) {
	cfg = cfg.New()
	start := time.Now()
	done := make(chan struct{})

	for i := 0; i < cfg.Goroutines; i++ {
		go func() {
			var x int64 = 1
			it := 0
			for {
				x ^= (x << 1) + (x >> 3)
				x *= 3
				it++
				if cfg.ShouldStop(start, it) {
					break
				}
			}
			done <- struct{}{}
		}()
	}

	for i := 0; i < cfg.Goroutines; i++ {
		<-done
	}
}
