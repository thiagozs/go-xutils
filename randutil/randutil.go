package randutil

import (
	"math/rand"
	"sync"
	"time"
)

// Source wraps math/rand.Rand so it can safely be shared by packages and
// goroutines. It must only be used for non-cryptographic randomness.
type Source struct {
	mu sync.Mutex
	r  *rand.Rand
}

// New returns a concurrency-safe pseudo-random source.
func New(seed int64) *Source {
	return &Source{r: rand.New(rand.NewSource(seed))}
}

func (s *Source) Intn(n int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.r.Intn(n)
}

func (s *Source) Int31n(n int32) int32 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.r.Int31n(n)
}

func (s *Source) Int63n(n int64) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.r.Int63n(n)
}

func (s *Source) Int63() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.r.Int63()
}

func (s *Source) Seed(seed int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.r.Seed(seed)
}

var defaultSource = New(time.Now().UnixNano())

// Default returns the concurrency-safe process-wide non-cryptographic source.
// Prefer New and dependency injection when deterministic behavior is needed.
func Default() *Source { return defaultSource }
