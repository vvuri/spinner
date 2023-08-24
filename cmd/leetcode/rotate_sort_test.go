package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func search(nums []int, target int) int {
	var min int = 0
	var max = len(nums) - 1
	var mid int

	for {
		mid = min + (max-min)/2

		fmt.Println(min, mid, max, " - ", nums[min], nums[mid], nums[max])

		if nums[mid] == target {
			return mid
		}
		if nums[max] == target {
			return max
		}
		if nums[min] == target {
			return min
		}
		if max == mid || min == mid {
			return -1
		}

		if nums[max] > nums[min] {
			if nums[min] > target || nums[max] < target {
				return -1
			}
			if nums[mid] > target {
				max = mid
				continue
			}
			if nums[mid] < target {
				min = mid
				continue
			}
		}

		if nums[mid] > nums[min] {
			if (nums[min] > target && nums[mid] > target) || (nums[min] < target && nums[mid] < target) {
				min = mid
				continue
			}
			if nums[min] < target && nums[mid] > target {
				max = mid
				continue
			}
		} else {
			if nums[mid] > target || nums[max] < target {
				max = mid
			} else {
				min = mid
			}
		}
	}
}

func TestPrintArr0(t *testing.T) {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	assert.Equal(t, search(arr, 0), 4, "Should be the same")
}

func TestPrintArr3(t *testing.T) {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	assert.Equal(t, search(arr, 3), -1, "Should be the same")
}

func TestPrintArr(t *testing.T) {
	arr := []int{1}
	assert.Equal(t, search(arr, 0), -1, "Should be the same")
}

func TestPrintArr10(t *testing.T) {
	arr := []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4}
	assert.Equal(t, search(arr, 3), 8, "Should be the same")
}

func TestPrintArr2(t *testing.T) {
	arr := []int{1, 3}
	assert.Equal(t, search(arr, 2), -1, "Should be the same")
}

func TestPrintArr4(t *testing.T) {
	arr := []int{3, 5, 1}
	assert.Equal(t, search(arr, 4), -1, "Should be the same")
}

func TestPrintArr8(t *testing.T) {
	arr := []int{4, 5, 6, 7, 8, 1, 2, 3}
	assert.Equal(t, search(arr, 8), 4, "Should be the same")
}
