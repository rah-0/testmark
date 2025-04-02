package model

import (
	"regexp"
)

const (
	GiB = int64(1 << 30)
	MiB = int64(1 << 20)
	KiB = int64(1 << 10)
)

const (
	Hour  = int64(3600000000000)
	Min   = int64(60000000000)
	Sec   = int64(1000000000)
	Milli = int64(1000000)
	Micro = int64(1000)
)

var (
	RegexBenchLine = regexp.MustCompile(`^(Benchmark\S+)\s+\d+\s+([\d.e+-]+) ns/op(?:\s+([\d.e+-]+) B/op)?(?:\s+(\d+) allocs/op)?`)
)
