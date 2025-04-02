package main

import (
	"fmt"

	"github.com/rah-0/testmark/model"
)

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

func trimTrailingSpace(s string) string {
	n := len(s)
	if n > 0 && s[n-1] == ' ' {
		return s[:n-1]
	}
	return s
}
