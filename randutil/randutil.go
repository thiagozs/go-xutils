package randutil

import (
	"math/rand"
	"time"
)

// Global is a seeded, concurrency-safe rand.Rand instance for non-cryptographic usage.
var Global *rand.Rand

func init() {
	Global = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// SeedForTest replaces the global RNG with a deterministic one for tests.
// Call defer randutil.SeedForTest(0) to restore behavior if needed.
func SeedForTest(seed int64) {
	Global = rand.New(rand.NewSource(seed))
}
