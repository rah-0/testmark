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
			"BenchmarkAllocOp-8\t500000\t512.0 ns/op\t128 B/op\t512ns",
		},

		// full line with all metrics
		{
			"BenchmarkFullOp-8          200000     1024.0 ns/op    2048 B/op    5 allocs/op",
			"BenchmarkFullOp-8\t200000\t1024.0 ns/op\t2048 B/op\t5 allocs/op\t1µs 24ns\t2KiB",
		},

		// full line with all metrics
		{
			"BenchmarkSample/Count100000-8              10000           9920268 ns/op             856 B/op         18 allocs/op",
			"BenchmarkSample/Count100000-8\t10000\t9920268 ns/op\t856 B/op\t18 allocs/op\t9ms 920µs 268ns",
		},

		// sub-nanosecond op
		{
			"BenchmarkNanoOp-8          10000000    0.9 ns/op",
			"BenchmarkNanoOp-8\t10000000\t0.9 ns/op\t0ns",
		},

		// no spacing at all (single spaces)
		{
			"BenchmarkTight-8  1  1000000.0 ns/op  4096 B/op  42 allocs/op",
			"BenchmarkTight-8\t1\t1000000.0 ns/op\t4096 B/op\t42 allocs/op\t1ms\t4KiB",
		},

		// only B/op (no ns/op)
		{
			"BenchmarkOnlyBOp-8    123456   256 B/op",
			"BenchmarkOnlyBOp-8\t123456\t256 B/op",
		},

		// only allocs/op (no ns/op or B/op)
		{
			"BenchmarkOnlyAllocs-8    98765   42 allocs/op",
			"BenchmarkOnlyAllocs-8\t98765\t42 allocs/op",
		},

		// values with decimal precision
		{
			"BenchmarkDecimal-8    1000   1234.56 ns/op   789.12 B/op   3 allocs/op",
			"BenchmarkDecimal-8\t1000\t1234.56 ns/op\t789.12 B/op\t3 allocs/op\t1µs 234ns\t789B",
		},

		// unusual label with special characters
		{
			"BenchmarkX/Y_Z-42    999    789.0 ns/op    123 B/op    1 allocs/op",
			"BenchmarkX/Y_Z-42\t999\t789.0 ns/op\t123 B/op\t1 allocs/op\t789ns",
		},

		// large values (GB range)
		{
			"BenchmarkBigMem-8    1   1000000000 ns/op   2147483648 B/op   100000 allocs/op",
			"BenchmarkBigMem-8\t1\t1000000000 ns/op\t2147483648 B/op\t100000 allocs/op\t1s\t2GiB",
		},

		// line with extra garbage at the end
		{
			"BenchmarkGarbage-8 100  500 ns/op  64 B/op  2 allocs/op  unexpected extra",
			"BenchmarkGarbage-8\t100\t500 ns/op\t64 B/op\t2 allocs/op",
		},

		// zero values
		{
			"BenchmarkZero-8    1    0 ns/op    0 B/op    0 allocs/op",
			"BenchmarkZero-8\t1\t0 ns/op\t0 B/op\t0 allocs/op",
		},

		// whitespace noise in the middle
		{
			"BenchmarkWeirdSpace-8    10     5000    ns/op     1024    B/op  4    allocs/op",
			"BenchmarkWeirdSpace-8\t10\t5000 ns/op\t1024 B/op\t4 allocs/op\t5µs\t1KiB",
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
