package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRsndomString(t *testing.T) {
	first := NewRandomString(6)
	second := NewRandomString(6)
	//fmt.Println(first, second)
	assert.NotEqual(t, first, second, "Should not be equal")
}
