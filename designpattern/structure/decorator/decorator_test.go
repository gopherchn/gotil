package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorSquare(t *testing.T) {
	square := Square{}
	colorSquare := NewColorSquare(&square, "yellow")
	res := colorSquare.Draw()
	assert.Equal(t, "square:yellow", res)
}