package aes

import "testing"

const (
	key = "IgkibX71IEf382PT"
	iv  = "IgkibX71IEf382PT"
)

func TestEncrypt(t *testing.T) {
	aes := New()
	a := aes.RegisterKeys(key, iv)
	t.Log(a.Encrypt("123456"))
}

func TestDecrypt(t *testing.T) {
	aes := New()
	a := aes.RegisterKeys(key, iv)
	t.Log(a.Decrypt("GO-ri84zevE-z1biJwfQPw=="))
}

func BenchmarkEncryptAndDecrypt(b *testing.B) {
	b.ResetTimer()
	aes := New()
	a := aes.RegisterKeys(key, iv)
	for i := 0; i < b.N; i++ {
		encryptString, _ := a.Encrypt("123456")
		_, _ = a.Decrypt(encryptString)
	}
}
