package randutil

import "testing"

func TestDefaultNotNil(t *testing.T) {
	if Default() == nil {
		t.Fatal("randutil.Default() is nil")
	}
}

func TestSourceIsDeterministic(t *testing.T) {
	v1 := New(42).Int63()
	v2 := New(42).Int63()
	if v1 != v2 {
		t.Fatalf("expected deterministic values for same seed, got %d and %d", v1, v2)
	}
}
