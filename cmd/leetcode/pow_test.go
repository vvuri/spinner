package leetcode

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "math"
  "testing"
)

func myPow(x float64, n int) float64 {
  if x < 0 {
    x = x * (-1)
    if n%2 != 0 {
      return -math.Exp(float64(n) * math.Log(x))
    }
  }
  return math.Exp(float64(n) * math.Log(x))
}

func TestPow1(t *testing.T) {
  assert.Equal(t, myPow(2.0, 2), 4.0, "Should be the same")
}

func TestPow2(t *testing.T) {
  assert.Equal(t, myPow(-2.0, 2), 4.0, "Should be the same")
}

func TestPow99(t *testing.T) {
  fmt.Println(math.Exp(-2.0))
  fmt.Println(math.Log(2.0))
}
