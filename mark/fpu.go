package main

import (
	"time"
)

func BenchmarkFloat64(cfg *Config) {
	cfg = cfg.New()
	start := time.Now()
	done := make(chan struct{})

	for i := 0; i < cfg.Goroutines; i++ {
		go func() {
			it := 0
			for {
				sum := 0.0
				for k := 0; k < 10000; k++ {
					term := 1.0 / (2*float64(k) + 1)
					if k%2 == 1 {
						term = -term
					}
					sum += term
				}
				_ = sum * 4 // π approx
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
