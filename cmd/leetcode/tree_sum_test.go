package leetcode

import (
  "fmt"
  "github.com/stretchr/testify/assert"
  "sort"
  "testing"
)

func threeSum(nums []int) [][]int {
	var rez [][]int
	if len(nums) < 3 {
		return rez
	}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j]
				if sum+nums[k] == 0 {
					// sort := make([]int, 3)
					arr := []int{nums[i], nums[j], nums[k]}
					sort.Slice(arr[:], func(i, j int) bool {
						return arr[:][i] > arr[:][j]
					})

					flag := true
					fmt.Println(nums[i], nums[j], nums[k])
					for m := 0; m < len(rez); m++ {
						if arr == rez[m] {
							flag = false
							break
						}
					}
					if flag {
						rez = append(rez, sort)
					}
				}
			}
		}
	}
	return rez
}

func TestTSum1(t *testing.T) {
	inp := []int{-1, 0, 1, 2, -1, -4}
	out := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	assert.Equal(t, threeSum(inp), out, "Should be the same")
}

func TestTSum2(t *testing.T) {
	assert.Equal(t, threeSum([]int{0, 1, 1}), [][]int{}, "Should be the same")
}
