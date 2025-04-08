package main

import (
	"time"
)

type Config struct {
	Duration   time.Duration
	Iterations int
	Goroutines int
	Args       any

	CounterFn func(int) // optional, for counting iterations
}

// ShouldUseDuration returns true if the benchmark should run until timeout
func (cfg *Config) ShouldUseDuration() bool {
	return cfg.Duration > 0
}

// ShouldUseIterations returns true if the benchmark should run a fixed number of iterations
func (cfg *Config) ShouldUseIterations() bool {
	return cfg.Iterations > 0 && cfg.Duration == 0
}

// New returns a new config with safe default values if unset
func (cfg *Config) New() *Config {
	if cfg.Duration <= 0 && cfg.Iterations <= 0 {
		cfg.Duration = 1 * time.Second
	}
	if cfg.Goroutines <= 0 {
		cfg.Goroutines = 1
	}
	return cfg
}

// ShouldStop tells the worker if it should stop based on config mode
func (cfg *Config) ShouldStop(start time.Time, currentIteration int) bool {
	switch {
	case cfg.ShouldUseDuration():
		return time.Since(start) >= cfg.Duration
	case cfg.ShouldUseIterations():
		return currentIteration >= cfg.Iterations
	default:
		return false
	}
}
