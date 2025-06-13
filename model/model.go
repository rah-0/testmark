// Package model defines constants used by testmark for formatting benchmark results.
// It includes time unit constants and memory size constants used for human-readable formatting.
package model

// Memory size constants in bytes using binary prefixes (powers of 1024)
const (
	// GiB represents one gibibyte (2^30 bytes)
	GiB = int64(1 << 30)
	// MiB represents one mebibyte (2^20 bytes)
	MiB = int64(1 << 20)
	// KiB represents one kibibyte (2^10 bytes)
	KiB = int64(1 << 10)
)

// Time unit constants in nanoseconds
const (
	// Hour represents one hour in nanoseconds (3.6e12)
	Hour  = int64(3600000000000)
	// Min represents one minute in nanoseconds (6e10)
	Min   = int64(60000000000)
	// Sec represents one second in nanoseconds (1e9)
	Sec   = int64(1000000000)
	// Milli represents one millisecond in nanoseconds (1e6)
	Milli = int64(1000000)
	// Micro represents one microsecond in nanoseconds (1e3)
	Micro = int64(1000)
)
