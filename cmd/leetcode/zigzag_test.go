package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func convert(s string, numRows int) string {
	//if numRows == 1 {
	//	return s
	//}

	var arr []string
	for i := 0; i < numRows; i++ {
		arr = append(arr, "")
	}
	for j := 0; j < len(s); {
		for i := 0; i < numRows; i++ {
			if j >= len(s) {
				break
			}
			arr[i] = arr[i] + string(s[j])
			j++
		}

		for i := numRows - 2; i > 0; i-- {
			if j >= len(s) {
				break
			}
			arr[i] = arr[i] + string(s[j])
			j++
		}
	}
	//for i := 0; i < numRows; i++ {
	//	fmt.Printf(arr[i])
	//	fmt.Print("\n")
	//}
	return strings.Join(arr, "")
}

func TestPrintOne(t *testing.T) {
	assert.Equal(t, convert("PAY", 1), "PAY", "Should be the same")
}

func TestPrintDigits(t *testing.T) {
	assert.Equal(t, convert("12345678901", 3), "15924680371", "Should be the same")
}

func TestPrintTree(t *testing.T) {
	assert.Equal(t, convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR", "Should be the same")
}

func TestPrintFore(t *testing.T) {
	assert.Equal(t, convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI", "Should be the same")
}
