package main

import (
	"math/rand"
	"time"
)

func BenchmarkMemory(cfg *Config) {
	cfg = cfg.New()
	start := time.Now()
	done := make(chan struct{})

	buf := make([]byte, 1<<27) // 128MB

	for i := 0; i < cfg.Goroutines; i++ {
		go func(x int) {
			r := rand.New(rand.NewSource(int64(x)))
			it := 0
			for {
				for j := 0; j < len(buf); j += 64 {
					buf[j] = byte(r.Intn(256))
				}
				it++
				if cfg.ShouldStop(start, it) {
					break
				}
			}
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < cfg.Goroutines; i++ {
		<-done
	}
}
