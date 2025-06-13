// Package testutil provides testing utilities for Go applications.
// It includes functions for safe test execution with panic recovery,
// and a TestMain wrapper for proper resource management in tests.
package testutil

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"testing"
)

// internalTestingMode is used to control error reporting behavior in recovery handlers.
// When true, errors are logged instead of failing the test (used by the package's own tests).
var internalTestingMode = false

// TestConfig defines arguments for the TestMain wrapper.
// It encapsulates all configuration needed for the TestMainWrapper function.
type TestConfig struct {
	M               *testing.M   // The testing.M instance from TestMain.
	LoadResources   func() error // Function to load necessary resources.
	UnloadResources func() error // Function to unload resources.
}

// TestMainWrapper is a wrapper for TestMain to handle resource loading/unloading.
// It provides a consistent pattern for setting up test resources before tests run
// and cleaning them up afterward, even if tests panic or fail.
// This helps prevent resource leaks and ensures proper cleanup.
func TestMainWrapper(c TestConfig) {
	// Load resources if a loader function is provided.
	if c.LoadResources != nil {
		if err := c.LoadResources(); err != nil {
			log.Fatalf("Failed to load resources: %v", err)
		}
	}

	// Run the tests.
	exitCode := c.M.Run()

	// Unload resources if an unloader function is provided.
	if c.UnloadResources != nil {
		if err := c.UnloadResources(); err != nil {
			log.Printf("Error unloading resources: %v", err)
		}
	}

	// Exit with the test run's exit code.
	os.Exit(exitCode)
}

// RunTestWithRecover executes a test function and recovers from panics, failing the test if a panic occurs.
// This allows tests to continue running even after a panic, which is especially useful in test suites.
// The panic stack trace and message will be reported in the test output.
func RunTestWithRecover(t *testing.T, testFunc func(*testing.T)) {
	defer RecoverTestHandler(t)
	testFunc(t)
}

// RecoverTestHandler recovers from a panic and marks the test as failed, printing the stack trace.
// This is meant to be used with defer in tests that might panic.
// Example usage:
//
//     func TestSomething(t *testing.T) {
//         defer RecoverTestHandler(t)
//         // test code that might panic
//     }
func RecoverTestHandler(t *testing.T) {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("Test panicked: %v\nStack trace:\n%s", r, debug.Stack())
		if internalTestingMode {
			t.Logf("[Recovered in test mode] %s", msg)
		} else {
			t.Errorf("%s", msg)
		}
	}
}

// RunBenchWithRecover executes a benchmark function and recovers from panics, failing the benchmark if a panic occurs.
// Similar to RunTestWithRecover, but designed for benchmark functions that take *testing.B instead of *testing.T.
// This allows the benchmark suite to continue even if one benchmark panics.
func RunBenchWithRecover(b *testing.B, testFunc func(*testing.B)) {
	defer RecoverBenchHandler(b)
	testFunc(b)
}

// RecoverBenchHandler recovers from a panic and marks the benchmark as failed, printing the stack trace.
// This is meant to be used with defer in benchmarks that might panic.
// Example usage:
//
//     func BenchmarkSomething(b *testing.B) {
//         defer RecoverBenchHandler(b)
//         // benchmark code that might panic
//     }
func RecoverBenchHandler(b *testing.B) {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("Test panicked: %v\nStack trace:\n%s", r, debug.Stack())
		if internalTestingMode {
			b.Logf("[Recovered in test mode] %s", msg)
		} else {
			b.Errorf("%s", msg)
		}
	}
}
