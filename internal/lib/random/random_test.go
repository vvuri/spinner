package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRandomStringLength(t *testing.T) {
	assert.Equal(t, len(NewRandomString(1)), 1, "Length should be 1")
	assert.Equal(t, len(NewRandomString(3)), 3, "Length should be 3")
	assert.Equal(t, len(NewRandomString(6)), 6, "Length should be 6")
}

func TestRandomStringDouble(t *testing.T) {
	first := NewRandomString(6)
	time.Sleep(1)
	second := NewRandomString(6)
	assert.NotEqual(t, first, second, "Should not be equal")
}
