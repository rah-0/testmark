package testutil

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"testing"
)

var internalTestingMode = false

// TestConfig defines arguments for the TestMain wrapper.
type TestConfig struct {
	M               *testing.M   // The testing.M instance from TestMain.
	LoadResources   func() error // Function to load necessary resources.
	UnloadResources func() error // Function to unload resources.
}

// TestMainWrapper is a wrapper for TestMain to handle resource loading/unloading.
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
func RunTestWithRecover(t *testing.T, testFunc func(*testing.T)) {
	defer RecoverTestHandler(t)
	testFunc(t)
}

// RecoverTestHandler recovers from a panic and marks the test as failed, printing the stack trace.
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

// RunBenchWithRecover executes a bench function and recovers from panics, failing the test if a panic occurs.
func RunBenchWithRecover(b *testing.B, testFunc func(*testing.B)) {
	defer RecoverBenchHandler(b)
	testFunc(b)
}

// RecoverBenchHandler recovers from a panic and marks the bench as failed, printing the stack trace.
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
