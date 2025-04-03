package benchutil

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
			"BenchmarkFastOp-8\t1000000\t31.2 ns/op\t31ns",
		},

		// ns/op + B/op
		{
			"BenchmarkAllocOp-8         500000     512.0 ns/op    128 B/op",
			"BenchmarkAllocOp-8\t500000\t512.0 ns/op\t128 B/op\t512ns\t128B",
		},

		// full line with all metrics
		{
			"BenchmarkFullOp-8          200000     1024.0 ns/op    2048 B/op    5 allocs/op",
			"BenchmarkFullOp-8\t200000\t1024.0 ns/op\t2048 B/op\t5 allocs/op\t1µs 24ns\t2KiB",
		},

		// malformed line
		{
			"Some unrelated log output",
			"Some unrelated log output",
		},
	}

	for _, tt := range tests {
		got := AppendConvertedLine(tt.input)
		if got != tt.want {
			t.Errorf("AppendConvertedLine(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestHumanNs(t *testing.T) {
	tests := []struct {
		input int64
		want  string
	}{
		{0, "0ns"},
		{999, "999ns"},
		{1000, "1µs"},
		{1500, "1µs 500ns"},
		{1000000, "1ms"},
		{1234567, "1ms 234µs 567ns"},
		{1000000000, "1s"},
		{3661000000000, "1h 1m 1s"},
	}

	for _, tt := range tests {
		got := HumanNs(tt.input)
		if got != tt.want {
			t.Errorf("HumanNs(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestHumanBytes(t *testing.T) {
	tests := []struct {
		input int64
		want  string
	}{
		// Only bytes
		{0, "0B"},
		{512, "512B"},

		// KiB range
		{1024, "1KiB"},
		{1536, "1KiB 512B"},

		// MiB range
		{1048576, "1MiB"},
		{1572864, "1MiB 512KiB"},

		// GiB range
		{1073741824, "1GiB"},
		{1610612736, "1GiB 512MiB"},
		{1611137024, "1GiB 512MiB 512KiB"},
		{1611137536, "1GiB 512MiB 512KiB 512B"},
	}

	for _, tt := range tests {
		got := HumanBytes(tt.input)
		if got != tt.want {
			t.Errorf("HumanBytes(%d) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
