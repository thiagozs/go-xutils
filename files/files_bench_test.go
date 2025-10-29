package files

import (
	"os"
	"path/filepath"
	"testing"
)

func prepareTempFile(t *testing.B, size int) string {
	// create temp dir
	dir := t.TempDir()
	fpath := filepath.Join(dir, "bench.txt")
	data := make([]byte, size)
	for i := range data {
		data[i] = 'a' + byte(i%26)
	}
	_ = os.WriteFile(fpath, data, 0644)
	return fpath
}

func BenchmarkReadFileLines(b *testing.B) {
	f := New()
	filePath := prepareTempFile(b, 1024*100) // 100KB
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := f.ReadFileLines(filePath)
		if err != nil {
			b.Fatalf("read error: %v", err)
		}
	}
}

func BenchmarkReadFileLinesBytes(b *testing.B) {
	f := New()
	filePath := prepareTempFile(b, 1024*100) // 100KB
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := f.ReadFileLinesBytes(filePath)
		if err != nil {
			b.Fatalf("read error: %v", err)
		}
	}
}
