package main

import (
	"crypto/sha256"
	"time"
)

func BenchmarkSha(cfg *Config, countFn func(int)) {
	cfg = cfg.New()
	start := time.Now()

	data := make([]byte, 1<<20)

	it := 0
	for {
		_ = sha256.Sum256(data)
		it++
		if cfg.ShouldStop(start, it) {
			break
		}
	}

	if countFn != nil {
		countFn(it)
	}
}
