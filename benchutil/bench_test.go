package benchutil

import (
	"testing"

	"github.com/rah-0/testmark/model"
)

func allocLarge() []byte {
	return make([]byte, 2*model.MiB)
}

func fastFunc() {
	s := 0
	for i := 0; i < 1000; i++ {
		s += i
	}
}

func allocSmall() []byte {
	return make([]byte, model.MiB)
}

func slowFunc() {
	s := 0
	for i := 0; i < 100000; i++ {
		s += i
	}
}

func TestMeasure_Basic(t *testing.T) {
	tuner := NewBench().SetRuns(10000)
	fast := tuner.Measure(fastFunc)
	slow := tuner.Measure(slowFunc)

	if fast.TimeNs <= 0 || slow.TimeNs <= 0 {
		t.Fatalf("Expected non-zero time results: fast=%d, slow=%d", fast.TimeNs, slow.TimeNs)
	}

	if slow.TimeNs <= fast.TimeNs {
		t.Errorf("Expected slowFunc to take longer than fastFunc")
	}
}

func TestMeasure_Delta(t *testing.T) {
	ref := BenchResult{NsPerOp: 1000, BytesPerOp: 100}
	target := BenchResult{NsPerOp: 2000, BytesPerOp: 300}

	nsDelta := ref.DeltaNsPct(target)
	bDelta := ref.DeltaBPct(target)

	if nsDelta != 100.0 {
		t.Errorf("Expected 100%% ns delta, got %.2f%%", nsDelta)
	}
	if bDelta != 200.0 {
		t.Errorf("Expected 200%% B delta, got %.2f%%", bDelta)
	}
}

func TestMeasure_Allocations(t *testing.T) {
	tuner := NewBench().SetRuns(10000)
	small := tuner.Measure(func() {
		_ = allocSmall()
	})
	large := tuner.Measure(func() {
		_ = allocLarge()
	})

	if small.BytesPerOp <= 0 {
		t.Errorf("Expected small allocation to report >0 B/op, got %d", small.BytesPerOp)
	}
	if large.BytesPerOp <= small.BytesPerOp {
		t.Errorf("Expected large allocation to use more B/op than small")
	}
}

func BenchmarkFastFunc(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		fastFunc()
	}
}

func BenchmarkSlowFunc(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		slowFunc()
	}
}

func BenchmarkAllocSmall(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		allocSmall()
	}
}

func BenchmarkAllocLarge(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		allocLarge()
	}
}
