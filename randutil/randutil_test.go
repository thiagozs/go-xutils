package randutil

import "testing"

func TestGlobalNotNil(t *testing.T) {
	if Global == nil {
		t.Fatal("randutil.Global is nil")
	}
}

func TestSeedForTestDeterministic(t *testing.T) {
	SeedForTest(42)
	v1 := Global.Int63()
	SeedForTest(42)
	v2 := Global.Int63()
	if v1 != v2 {
		t.Fatalf("expected deterministic values for same seed, got %d and %d", v1, v2)
	}
}
