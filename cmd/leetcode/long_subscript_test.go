package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	max_len := 0
	max := len(s)
	for i := 0; i < max; i++ {
		for j := i + 1; j < max; j++ {
			sub := s[i:j]
			fmt.Println("= ", sub)
			for k := j; k < max-len(sub); k++ {
				//fmt.Println(s[k : k+len(sub)])
				if sub == s[k:k+len(sub)] {
					fmt.Println("==>", sub)
					if max_len < len(sub) {
						max_len = len(sub)
						break
					}
				}
				if k > max_len+1 {
					break
				}
			}
		}
	}
	return max_len
}

func TestLongSubscriptionOne(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
}

func TestLongSubscriptionTwo(t *testing.T) {
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
}

func TestLongSubscriptionThree(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
}
