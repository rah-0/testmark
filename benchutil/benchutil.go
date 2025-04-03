package benchutil

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rah-0/testmark/model"
)

// AppendConvertedLine checks for a benchmark output line and appends human-friendly conversions.
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
		if nsHuman != nsOp+"ns" {
			parts = append(parts, nsHuman)
		}
	}
	if hasBOp {
		bHuman := HumanBytes(int64(parseFloat(bOp)))
		if bHuman != bOp+"B" {
			parts = append(parts, bHuman)
		}
	}

	return strings.Join(parts, "\t")
}

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

// parseFloat parses float-like numbers from Go benchmark lines
func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func trimTrailingSpace(s string) string {
	n := len(s)
	if n > 0 && s[n-1] == ' ' {
		return s[:n-1]
	}
	return s
}
