package strings

import (
	"testing"
)

func BenchmarkUniqueSlug(b *testing.B) {
	input := "Exemplo de Título com Acentuação e ## símbolos!"
	for i := 0; i < b.N; i++ {
		_ = UniqueSlug(input)
	}
}

func BenchmarkRemoveStopWords(b *testing.B) {
	input := "Este é um texto de exemplo para testar a remoção de palavras de parada como e, ou, mas, por que, quando"
	for i := 0; i < b.N; i++ {
		_ = RemoveStopWords(input)
	}
}

func BenchmarkRandomAlphanumeric(b *testing.B) {
	generator := NewGenerator()
	for i := 0; i < b.N; i++ {
		_ = generator.RandomAlphanumeric(32)
	}
}
