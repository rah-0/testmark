package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rah-0/testmark/model"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(appendConvertedLine(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

// appendConvertedLine checks for a benchmark output line and appends human-friendly conversions.
func appendConvertedLine(line string) string {
	m := model.RegexBenchLine.FindStringSubmatch(line)
	if m == nil {
		return line
	}

	label := m[1]
	nsStr := m[2]
	nsVal := parseFloat(nsStr)
	nsHuman := HumanNs(int64(nsVal))

	var bHuman string
	var parts []string

	// Start with benchmark label + ns/op
	parts = append(parts, fmt.Sprintf("%s\t%s ns/op", label, nsStr))

	if m[3] != "" {
		// B/op present
		bVal := parseFloat(m[3])
		bHuman = HumanBytes(int64(bVal))
		parts = append(parts, fmt.Sprintf("%s B/op", m[3]))
	}

	// Append allocs/op, if it exists
	if m[4] != "" {
		parts = append(parts, fmt.Sprintf("%s allocs/op", m[4]))
	}

	// Append human conversions: ns first, then bytes
	human := nsHuman
	if bHuman != "" {
		human += "\t" + bHuman
	}
	parts = append(parts, human)

	// Join and return the final formatted string
	return strings.Join(parts, "\t")
}

// parseFloat parses float-like numbers from Go benchmark lines
func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
