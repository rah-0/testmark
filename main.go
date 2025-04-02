package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rah-0/testmark/benchutil"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(benchutil.AppendConvertedLine(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
