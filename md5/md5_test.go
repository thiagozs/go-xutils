package md5

import "testing"

func TestEncrypt(t *testing.T) {
	t.Log(New().MD5Hash("123456"))
}

func BenchmarkEncrypt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		New().MD5Hash("123456")
	}
}
