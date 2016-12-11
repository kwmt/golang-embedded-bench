package embeded

import (
	"testing"
)

func BenchmarkEmbeded(b *testing.B) {
	rg := RouterGroup{root: true}
	engine := &Engine{rg}
	for i := 0; i < b.N; i++ {
		Embedded(engine)
	}
}

func BenchmarkEmbededNo(b *testing.B) {
	rg := RouterGroup{root: true}
	engine := &Engine{rg}
	for i := 0; i < b.N; i++ {
		EmbeddedNo(engine)
	}
}
