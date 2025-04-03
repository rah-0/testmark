package benchutil

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rah-0/testmark/model"
)

// AppendConvertedLine checks for a benchmark output line and appends human-friendly conversions.
func AppendConvertedLine(line string) string {
	matches := model.RegexBenchLine.FindStringSubmatch(line)
	if matches == nil {
		return line
	}

	label := matches[1]
	count := matches[2]
	nsStr := matches[3]
	bStr := matches[4]
	aStr := matches[5]

	nsVal := parseFloat(nsStr)
	nsHuman := HumanNs(int64(nsVal))

	var bHuman string
	var parts []string

	// Add label, count, and ns/op
	parts = append(parts, fmt.Sprintf("%s\t%s\t%s ns/op", label, count, nsStr))

	// Add B/op if present
	if bStr != "" {
		bVal := parseFloat(bStr)
		bHuman = HumanBytes(int64(bVal))
		parts = append(parts, fmt.Sprintf("%s B/op", bStr))
	}

	// Add allocs/op if present
	if aStr != "" {
		parts = append(parts, fmt.Sprintf("%s allocs/op", aStr))
	}

	// Add human-readable summary
	human := nsHuman
	if bHuman != "" {
		human += "\t" + bHuman
	}
	parts = append(parts, human)

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
