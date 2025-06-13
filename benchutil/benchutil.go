// Package benchutil provides utilities for formatting Go benchmark output and measuring performance.
// It includes functions to convert raw benchmark values to human-readable formats
// and utilities for running custom benchmarks outside of the standard testing framework.
package benchutil

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rah-0/testmark/model"
)

// AppendConvertedLine takes a standard Go benchmark output line and appends human-friendly conversions.
// It detects ns/op, B/op, and allocs/op measurements, and appends more readable
// representations like "CPU[3ms 651µs 349ns]" and "MEM[1KiB 104B]".
// If the input line is not a benchmark line, it returns the original line unmodified.
func AppendConvertedLine(line string) string {
	hasNsOp := strings.Contains(line, "ns/op")
	hasBOp := strings.Contains(line, "B/op")
	hasAllocsOp := strings.Contains(line, "allocs/op")
	if !hasNsOp && !hasBOp && !hasAllocsOp {
		return line
	}

	fields := strings.Fields(line)

	label := fields[0]
	count := fields[1]

	nsOp := ""
	if hasNsOp {
		nsOp = fields[slices.Index(fields, "ns/op")-1]
	}

	bOp := ""
	if hasBOp {
		bOp = fields[slices.Index(fields, "B/op")-1]
	}

	allocsOp := ""
	if hasAllocsOp {
		allocsOp = fields[slices.Index(fields, "allocs/op")-1]
	}

	parts := []string{label, count}
	if hasNsOp {
		parts = append(parts, nsOp+" ns/op")
	}
	if hasBOp {
		parts = append(parts, bOp+" B/op")
	}
	if hasAllocsOp {
		parts = append(parts, allocsOp+" allocs/op")
	}

	if hasNsOp {
		nsHuman := HumanNs(int64(parseFloat(nsOp)))
		if nsHuman != nsOp+"ns" && nsHuman != "0ns" {
			parts = append(parts, "CPU["+nsHuman+"]")
		}
	}
	if hasBOp {
		bHuman := HumanBytes(int64(parseFloat(bOp)))
		if bHuman != bOp+"B" {
			parts = append(parts, "MEM["+bHuman+"]")
		}
	}

	return strings.Join(parts, "\t")
}

// HumanNs converts nanoseconds to a human-readable string with appropriate units.
// It breaks down the time into hours, minutes, seconds, milliseconds, microseconds, and nanoseconds
// as appropriate for the magnitude of the input.
func HumanNs(ns int64) string {
	out := ""

	if ns >= model.Hour {
		out += fmt.Sprintf("%dh ", ns/model.Hour)
		ns %= model.Hour
	}
	if ns >= model.Min {
		out += fmt.Sprintf("%dm ", ns/model.Min)
		ns %= model.Min
	}
	if ns >= model.Sec {
		out += fmt.Sprintf("%ds ", ns/model.Sec)
		ns %= model.Sec
	}
	if ns >= model.Milli {
		out += fmt.Sprintf("%dms ", ns/model.Milli)
		ns %= model.Milli
	}
	if ns >= model.Micro {
		out += fmt.Sprintf("%dµs ", ns/model.Micro)
		ns %= model.Micro
	}
	if ns > 0 || out == "" {
		out += fmt.Sprintf("%dns", ns)
	}

	return trimTrailingSpace(out)
}

// HumanBytes converts a byte count to a human-readable string with appropriate units.
// It scales the value using binary prefixes (KiB, MiB, GiB) based on powers of 1024.
func HumanBytes(b int64) string {
	out := ""
	if b >= model.GiB {
		out += fmt.Sprintf("%dGiB ", b/model.GiB)
		b %= model.GiB
	}
	if b >= model.MiB {
		out += fmt.Sprintf("%dMiB ", b/model.MiB)
		b %= model.MiB
	}
	if b >= model.KiB {
		out += fmt.Sprintf("%dKiB ", b/model.KiB)
		b %= model.KiB
	}
	if b > 0 || out == "" {
		out += fmt.Sprintf("%dB", b)
	}
	return trimTrailingSpace(out)
}

// parseFloat parses float-like numbers from Go benchmark lines.
func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// trimTrailingSpace removes a single trailing space from a string if one exists.
// This is used to clean up the formatted output strings.
func trimTrailingSpace(s string) string {
	n := len(s)
	if n > 0 && s[n-1] == ' ' {
		return s[:n-1]
	}
	return s
}
