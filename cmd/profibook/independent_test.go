package profibook

import (
	"testing"
	"time"
)

func Add(a, b int) int {
	time.Sleep(500 * time.Millisecond)
	return a + b
}

func TestAddWithSubTest(t *testing.T) {
	testCases := []struct {
		name   string
		a      int
		b      int
		result int
	}{
		{"OneDigit", 1, 2, 3},
		{"TwoDigits", 12, 30, 42},
		{"ThreeDigits", 100, -1, 99},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := Add(testCase.a, testCase.b)
			if result != testCase.result {
				t.Errorf("Adding %d and %d doesn't produce %d, instead it produces %d",
					testCase.a, testCase.b, testCase.result, result)
			}
		})
	}
}
