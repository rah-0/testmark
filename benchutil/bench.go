package benchutil

import (
	"runtime"
	"time"
)

// BenchResult stores the results of a benchmark run.
// It includes both the total measurements and the per-operation averages.
type BenchResult struct {
	// TimeNs is the total time taken by all benchmark runs in nanoseconds
	TimeNs     int64
	// Bytes is the total memory allocated during all benchmark runs
	Bytes      int64
	// NsPerOp is the average time per operation in nanoseconds
	NsPerOp    int64
	// BytesPerOp is the average memory allocated per operation in bytes
	BytesPerOp int64
}

// Bench is a lightweight benchmarking tool that can be used outside
// of the standard Go testing framework.
// It measures execution time and memory allocations for functions.
type Bench struct {
	// runsPerCase is the number of times each benchmark function will be executed
	runsPerCase int
}

// NewBench creates a new benchmark runner with default settings.
// By default, each benchmark function will be run 1000 times.
func NewBench() *Bench {
	return &Bench{
		runsPerCase: 1000,
	}
}

// SetRuns sets the number of times each benchmark function will be executed.
// Higher values give more accurate results but take longer to complete.
// Returns the Bench instance for method chaining.
func (t *Bench) SetRuns(n int) *Bench {
	t.runsPerCase = n
	return t
}

// Measure executes the provided function multiple times and measures its performance.
// It returns a BenchResult containing the timing and memory allocation statistics.
// Note: For memory measurements, this uses runtime.ReadMemStats which introduces some overhead
// and may be less precise than the built-in testing package's benchmarking.
func (t *Bench) Measure(fn func()) BenchResult {
	var totalNs, totalBytes int64
	for i := 0; i < t.runsPerCase; i++ {
		r := getMemInfo(fn)
		totalNs += r.timeNs
		totalBytes += r.bytes
	}

	avgNs := totalNs / int64(t.runsPerCase)
	avgBytes := totalBytes / int64(t.runsPerCase)

	res := BenchResult{
		TimeNs:     totalNs,
		Bytes:      totalBytes,
		NsPerOp:    avgNs,
		BytesPerOp: avgBytes,
	}

	return res
}

// DeltaNsPct calculates the percentage difference in execution time between this benchmark and the target.
// Returns a positive value if the target is slower (took more time), negative if it's faster.
// If the current benchmark's NsPerOp is 0, returns 0 to avoid division by zero.
func (a BenchResult) DeltaNsPct(target BenchResult) float64 {
	if a.NsPerOp == 0 {
		return 0
	}
	return float64(target.NsPerOp-a.NsPerOp) / float64(a.NsPerOp) * 100
}

// DeltaBPct calculates the percentage difference in memory allocation between this benchmark and the target.
// Returns a positive value if the target uses more memory, negative if it uses less.
// If the current benchmark's BytesPerOp is 0, returns 0 to avoid division by zero.
func (a BenchResult) DeltaBPct(target BenchResult) float64 {
	if a.BytesPerOp == 0 {
		return 0
	}
	return float64(target.BytesPerOp-a.BytesPerOp) / float64(a.BytesPerOp) * 100
}

// memInfo is an internal struct used to store timing and memory allocation information
// for a single run of a benchmark function.
type memInfo struct {
	timeNs int64 // execution time in nanoseconds
	bytes  int64 // memory allocated in bytes
}

// getMemInfo executes a function once and returns information about its execution time
// and memory usage. It uses runtime.ReadMemStats to measure memory allocation differences
// before and after execution.
// Note that ReadMemStats itself has overhead and may affect very small measurements.
func getMemInfo(fn func()) memInfo {
	var memBefore, memAfter runtime.MemStats
	runtime.ReadMemStats(&memBefore)
	start := time.Now()
	fn()
	dur := time.Since(start).Nanoseconds()
	runtime.ReadMemStats(&memAfter)
	return memInfo{
		timeNs: dur,
		bytes:  int64(memAfter.TotalAlloc - memBefore.TotalAlloc),
	}
}
