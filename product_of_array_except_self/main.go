package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums)) // [24, 12, 8, 6]

	nums = []int{-1, 1, 0, -3, 3}
	fmt.Println(productExceptSelf(nums)) // [0, 0, 9, 0, 0]
}

func productExceptSelf(nums []int) []int {
	size := len(nums)
	output := make([]int, size)

	for i := range output {
		output[i] = 1
	}

	prefixProduct := 1
	for i := range nums {
		output[i] *= prefixProduct
		prefixProduct *= nums[i]
	}

	suffixProduct := 1
	for i := size - 1; i >= 0; i-- {
		output[i] *= suffixProduct
		suffixProduct *= nums[i]
	}

	return output
}
