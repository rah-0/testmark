package testmark

import (
	"fmt"
	"regexp"
)

const (
	GiB = int64(1 << 30)
	MiB = int64(1 << 20)
	KiB = int64(1 << 10)
)

var (
	RegexBenchLine = regexp.MustCompile(`^(Benchmark\S+)\s+(\d+)\s+([\d.]+) ns/op\s+([\d.]+) B/op\s+(\d+) allocs/op$`)
)

func HumanNs(ns int64) string {
	out := ""
	if ns >= 3600000000000 {
		out += fmt.Sprintf("%dh ", ns/3600000000000)
		ns %= 3600000000000
	}
	if ns >= 60000000000 {
		out += fmt.Sprintf("%dm ", ns/60000000000)
		ns %= 60000000000
	}
	if ns >= 1000000000 {
		out += fmt.Sprintf("%ds ", ns/1000000000)
		ns %= 1000000000
	}
	if ns >= 1000000 {
		out += fmt.Sprintf("%dms ", ns/1000000)
		ns %= 1000000
	}
	if ns >= 1000 {
		out += fmt.Sprintf("%dµs ", ns/1000)
		ns %= 1000
	}
	if ns > 0 || out == "" {
		out += fmt.Sprintf("%dns", ns)
	}
	return trimTrailingSpace(out)
}

func HumanBytes(b int64) string {
	out := ""
	if b >= GiB {
		out += fmt.Sprintf("%dGiB ", b/GiB)
		b %= GiB
	}
	if b >= MiB {
		out += fmt.Sprintf("%dMiB ", b/MiB)
		b %= MiB
	}
	if b >= KiB {
		out += fmt.Sprintf("%dKiB ", b/KiB)
		b %= KiB
	}
	if b > 0 || out == "" {
		out += fmt.Sprintf("%dB", b)
	}
	return trimTrailingSpace(out)
}

func trimTrailingSpace(s string) string {
	n := len(s)
	if n > 0 && s[n-1] == ' ' {
		return s[:n-1]
	}
	return s
}
