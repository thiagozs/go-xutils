package strings

import (
	"testing"
)

func BenchmarkGenerateUniqueSlug(b *testing.B) {
	input := "Exemplo de Título com Acentuação e ## símbolos!"
	for i := 0; i < b.N; i++ {
		_ = New().GenerateUniqueSlug(input)
	}
}

func BenchmarkRemoveStopWords(b *testing.B) {
	s := New()
	input := "Este é um texto de exemplo para testar a remoção de palavras de parada como e, ou, mas, por que, quando"
	for i := 0; i < b.N; i++ {
		_ = s.RemoveStopWords(input)
	}
}

func BenchmarkRandomStr(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		_ = s.RandomStr(32)
	}
}
