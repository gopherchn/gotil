package singleton 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	assert.Equal(t, Instance(), Instance())
}

func BenchmarkInstanceParallel(b *testing.B) {
	b.RunParallel((func(p *testing.PB) {
		for p.Next() {
			if Instance() != Instance() {
				b.Error("failed")
			}
		}
	}))
}

func TestLazyInstance(t *testing.T) {
	assert.Equal(t, LazyInstance(), LazyInstance())
}

func BenchmarkLazyInstanceParallel(b *testing.B) {
	b.RunParallel((func(p *testing.PB) {
		for p.Next() {
			if LazyInstance() != LazyInstance() {
				b.Error("failed")
			}
		}
	}))
}
