[![Go Report Card](https://goreportcard.com/badge/github.com/rah-0/testmark)](https://goreportcard.com/report/github.com/rah-0/testmark)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<a href="https://www.buymeacoffee.com/rah.0" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/arial-orange.png" alt="Buy Me A Coffee" height="50"></a>


# testmark
`testmark` is a Go tool and library for enhancing and automating benchmark workflows.
It provides a CLI for formatting `go test -bench` output and a library with utilities for custom benchmarking.

---
## CLI
To install the latest version of testmark, run the following command:
```
go install github.com/rah-0/testmark@latest
```

## Usage
Once installed, you can use `testmark` to process Go **benchmark** output.
### Example
To format your **benchmark** results using `testmark`, you can pipe the output from go test directly into the tool:
```
go test -run=^$ -bench="^BenchmarkFunction$" -benchmem | testmark
```
This will take the **benchmark** results and output them in a more readable format, where time units like `ns/op` and `B/op` will be converted to friendly forms separated by `\t`.

### Output Example
Before:
```
Sample/Count10-8       114430 ns/op    808 B/op        18 allocs/op
Sample/Count100-8      126958 ns/op    840 B/op        18 allocs/op
Sample/Count1000-8     248419 ns/op    848 B/op        18 allocs/op
Sample/Count10000-8    797351 ns/op    848 B/op        18 allocs/op
Sample/Count100000-8   3651349 ns/op   856 B/op        18 allocs/op
Sample/Count1000000-8  37593152 ns/op  863 B/op        18 allocs/op
```
After:
```
Sample/Count10-8       114430 ns/op    808 B/op        18 allocs/op    114µs 430ns     808B
Sample/Count100-8      126958 ns/op    840 B/op        18 allocs/op    126µs 958ns     840B
Sample/Count1000-8     248419 ns/op    848 B/op        18 allocs/op    248µs 419ns     848B
Sample/Count10000-8    797351 ns/op    848 B/op        18 allocs/op    797µs 351ns     848B
Sample/Count100000-8   3651349 ns/op   856 B/op        18 allocs/op    3ms 651µs 349ns 856B
Sample/Count1000000-8  37593152 ns/op  863 B/op        18 allocs/op    37ms 593µs 152ns        863B
```

## ⚠️ Warning: Potential Issues with Tool Integration
While `testmark` enhances Go benchmark output by converting raw values into readable formats, be cautious when using it in automated toolchains or with other CLI tools.
- **Formatting changes**: The tool adds readable units (e.g., `s`, `KiB`) to the benchmark output, which **may break downstream tools** expecting a specific format (e.g., exact tab or space separation).
- **Overwriting output files**: If redirecting output, ensure files are not unintentionally overwritten by `testmark`.
- **Unmatched input**: `testmark` will leave **non-benchmark lines** or lines not matching the expected Go benchmark format unchanged, ensuring no unexpected alterations occur.

Always test the full integration if you plan to use `testmark` as part of a larger automation pipeline.

---

# 📦 Library Utilities
In addition to CLI formatting, `testmark` also provides **benchmarking utilities** as a Go package under `benchutil` and testing recovery helpers under `testutil`.

## Benchmarking Tools
The `benchutil` package provides a lightweight, self-contained benchmarking utility to measure performance and memory usage without relying on `go test`.

### Example Usage

```go
package main

import (
	"fmt"

	"github.com/rah-0/testmark/benchutil"
)

func allocLarge() []byte {
	return make([]byte, 2*int64(1 << 20)) // 2MiB
}

func allocSmall() []byte {
	return make([]byte, int64(1 << 20)) // 1MiB
}

func main() {
	b := benchutil.NewBench().SetRuns(1000)
	slow := b.Measure(func(){_ = allocSmall()})
	fast := b.Measure(func(){_ = allocLarge()})

	nsDiff := fast.DeltaNsPct(slow)
	bDiff := fast.DeltaBPct(slow)

	fmt.Printf("Time diff: %.2f%%\n", nsDiff)
	fmt.Printf("Memory diff: %.2f%%\n", bDiff)
}
```
### ✎ Note
Naturally, there are precision issues when benching outside of `*testing.B`.
```go
func BenchmarkAllocSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		allocSmall()
	}
}
//Outputs:
//BenchmarkAllocSmall-8   	   15234	     81454 ns/op	 1048579 B/op	       1 allocs/op

func BenchmarkAllocLarge(b *testing.B) {
    b.ResetTimer()
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        allocLarge()
    }
}
//Outputs:
//BenchmarkAllocLarge-8   	    3946	    286477 ns/op	 2097160 B/op	       1 allocs/op

func TestMeasure_Allocations(t *testing.T) {
    tuner := NewBench().SetRuns(10000)
    small := tuner.Measure(func() {
        _ = allocSmall()
    })
    large := tuner.Measure(func() {
        _ = allocLarge()
    })

    fmt.Println(small.NsPerOp) //84120
    fmt.Println(small.BytesPerOp) //1048596
    
    fmt.Println(large.NsPerOp) //216284
    fmt.Println(large.BytesPerOp) //2097165
}
```
So, what's the drift?

| Source      | Type  | Ns/Op   | B/Op     | Drift Ns | Drift B |
|-------------|-------|---------|----------|----------|---------|
| go test     | Small | 81454   | 1048579  | —        | —       |
| benchutil   | Small | 84120   | 1048596  | +3.27%   | +0.00%  |
| go test     | Large | 286477  | 2097160  | —        | —       |
| benchutil   | Large | 216284  | 2097165  | −24.51%  | +0.00%  |

### 🧠 Notes:
- Memory (`B/op`) is nearly identical in both tools, confirming accuracy of the GC-based measurement.
- Time (`ns/op`) is less consistent, `benchutil` reports higher time for small allocs and lower for large allocs. This is somewhat expected due to GC noise, scheduling, and lack of internal runtime calibration (`b.ResetTimer()`, etc.).

---

## Test Wrappers
The `testutil` package offers wrappers to:
- Catch and report panics in tests or benchmarks
- Wrap TestMain() with setup/teardown hooks

### Example: Wrapping TestMain
```go
func TestMain(m *testing.M) {
	testutil.TestMainWrapper(testutil.TestConfig{
		M: m,
		LoadResources: func() error {
			// Setup before tests
			return nil
		},
		UnloadResources: func() error {
			// Cleanup after tests
			return nil
		},
	})
}
```
### Example: Recovering From Panics
```go
func TestSafe(t *testing.T) {
	testutil.RunTestWithRecover(t, func(t *testing.T) {
		panic("unexpected")
	})
}
```
## ✎ Notes:
- `UnloadResources` runs even if tests fail or panic
- Keeps tests alive even if panics occur
- Prints full stack traces for debug

# 💚 Support
If this saved you time or brought value to your project, feel free to show some support. Every bit is appreciated 🙂

[![Buy Me A Coffee](https://cdn.buymeacoffee.com/buttons/default-orange.png)](https://www.buymeacoffee.com/rah.0)
