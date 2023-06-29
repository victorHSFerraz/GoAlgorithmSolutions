// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // 9
}

func longestConsecutive(nums []int) int {
	results := make([]int, 0)
	longest := make([]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for _, num := range nums {
		if len(longest) == 0 {
			longest = append(longest, num)
		} else {
			if (num - longest[len(longest)-1]) == 1 {
				longest = append(longest, num)
			} else if (num - longest[len(longest)-1]) == 0 {
				continue
			} else {
				longest = append(longest[:0])
				longest = append(longest, num)
			}
		}
		results = append(results, len(longest))
	}
	max := 0
	for _, result := range results {
		if result > max {
			max = result
		}
	}

	return max
}
