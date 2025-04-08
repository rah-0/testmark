package main

import (
	"fmt"
	"time"
)

func main() {
	cfg := &Config{
		Duration:   2 * time.Second,
		Goroutines: 4,
	}

	result := RunMark(cfg, BenchmarkSha)

	fmt.Printf("SHA256: %d ops in %v (%.2f ops/sec)\n",
		result.TotalIterations, result.Elapsed, result.ItersPerSec)
}
