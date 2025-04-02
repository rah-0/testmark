[![Go Report Card](https://goreportcard.com/badge/github.com/rah-0/testmark)](https://goreportcard.com/report/github.com/rah-0/testmark)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

<a href="https://www.buymeacoffee.com/rah.0" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/arial-orange.png" alt="Buy Me A Coffee" height="50"></a>


# testmark

## Installation for CLI usage
To install the latest version of testmark, run the following command:
```
go install github.com/rah-0/testmark@latest
```

## CLI Usage
Once installed, you can use `testmark` to process Go **benchmark** output.
### Example
To format your **benchmark** results using `testmark`, you can pipe the output from go test directly into the tool:
```
go test -run=^$ -bench="^BenchmarkFunction$" -benchmem | testmark
```
This will take the **benchmark** results and output them in a more readable format, where time units like `ns/op` and `B/op` will be converted to human-friendly forms separated by `\t`.

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
While `testmark` enhances Go benchmark output by converting raw values into human-readable formats, be cautious when using it in automated toolchains or with other CLI tools.
- **Formatting changes**: The tool adds human-readable units (e.g., `s`, `KiB`) to the benchmark output, which **may break downstream tools** expecting a specific format (e.g., exact tab or space separation).
- **Overwriting output files**: If redirecting output, ensure files are not unintentionally overwritten by `testmark`.
- **Unmatched input**: `testmark` will leave **non-benchmark lines** or lines not matching the expected Go benchmark format unchanged, ensuring no unexpected alterations occur.

Always test the full integration if you plan to use `testmark` as part of a larger automation pipeline.

# 💚 Support
If this saved you time or brought value to your project, feel free to show some support. Every bit is appreciated 🙂

[![Buy Me A Coffee](https://cdn.buymeacoffee.com/buttons/default-orange.png)](https://www.buymeacoffee.com/rah.0)
