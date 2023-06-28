// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	num       int
	frequency int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	var pairs []Pair
	for num, frequency := range m {
		pairs = append(pairs, Pair{num, frequency})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].frequency > pairs[j].frequency
	})

	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i].num
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k)) // [1, 2]

	nums = []int{1}
	k = 1
	fmt.Println(topKFrequent(nums, k)) // [1]
}
