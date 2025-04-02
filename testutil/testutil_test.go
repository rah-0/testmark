package testutil

import (
	"errors"
	"testing"
)

func TestMain(m *testing.M) {
	TestMainWrapper(TestConfig{
		M: m,
		LoadResources: func() error {
			internalTestingMode = true
			return nil
		},
		UnloadResources: func() error {
			if !internalTestingMode {
				return errors.New("internalTestingMode should be true")
			}
			return nil
		},
	})
}

func TestRecoverTestHandlerWorks(t *testing.T) {
	called := false
	RunTestWithRecover(t, func(t *testing.T) {
		defer func() { called = true }()
		panic("simulated test panic")
	})
	if !called {
		t.Errorf("Expected recovery block to execute")
	}
}

func TestRecoverTestHandler_NoPanic(t *testing.T) {
	RunTestWithRecover(t, func(t *testing.T) {
		// no panic
	})
}

func BenchmarkRecoverBenchHandlerWorks(b *testing.B) {
	called := false
	RunBenchWithRecover(b, func(b *testing.B) {
		defer func() { called = true }()
		panic("simulated bench panic")
	})
	if !called {
		b.Errorf("Expected recovery block to execute")
	}
}

func BenchmarkRecoverBenchHandler_NoPanic(b *testing.B) {
	RunBenchWithRecover(b, func(b *testing.B) {
		// no panic
	})
}
