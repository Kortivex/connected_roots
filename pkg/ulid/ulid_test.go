package ulid

import (
	"testing"
)

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Generate()
	}
}
