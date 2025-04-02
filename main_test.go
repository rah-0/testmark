package main

import (
	"testing"
)

func TestAppendConvertedLine(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// ns/op only
		{
			"BenchmarkFastOp-8          1000000    31.2 ns/op",
			"BenchmarkFastOp-8\t31.2 ns/op\t31ns",
		},

		// ns/op + B/op
		{
			"BenchmarkAllocOp-8         500000     512.0 ns/op    128 B/op",
			"BenchmarkAllocOp-8\t512.0 ns/op\t128 B/op\t512ns\t128B",
		},

		// full line with all metrics
		{
			"BenchmarkFullOp-8          200000     1024.0 ns/op    2048 B/op    5 allocs/op",
			"BenchmarkFullOp-8\t1024.0 ns/op\t2048 B/op\t5 allocs/op\t1µs 24ns\t2KiB",
		},

		// malformed line
		{
			"Some unrelated log output",
			"Some unrelated log output",
		},
	}

	for _, tt := range tests {
		got := appendConvertedLine(tt.input)
		if got != tt.want {
			t.Errorf("appendConvertedLine(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
