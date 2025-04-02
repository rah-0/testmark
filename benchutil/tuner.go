package benchutil

import (
	"runtime"
	"time"
)

type BenchResult struct {
	TimeNs     int64
	Bytes      int64
	NsPerOp    int64
	BytesPerOp int64
}

type Bench struct {
	runsPerCase int
}

func NewBench() *Bench {
	return &Bench{
		runsPerCase: 1000,
	}
}

func (t *Bench) SetRuns(n int) *Bench {
	t.runsPerCase = n
	return t
}

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

func (a BenchResult) DeltaNsPct(target BenchResult) float64 {
	if a.NsPerOp == 0 {
		return 0
	}
	return float64(target.NsPerOp-a.NsPerOp) / float64(a.NsPerOp) * 100
}

func (a BenchResult) DeltaBPct(target BenchResult) float64 {
	if a.BytesPerOp == 0 {
		return 0
	}
	return float64(target.BytesPerOp-a.BytesPerOp) / float64(a.BytesPerOp) * 100
}

type memInfo struct {
	timeNs int64
	bytes  int64
}

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
