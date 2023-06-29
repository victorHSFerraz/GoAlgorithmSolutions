package main

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	twoSum(nums, target)
}

//most efficient solution for memory

func twoSum(nums []int, target int) []int {
	for index := range nums {
		for i := index + 1; i < len(nums); i++ {
			if nums[index]+nums[i] == target {
				return []int{index, i}
			}
		}
	}
	return []int{0, 0}
}

//most efficient solution for time

// func twoSumFast(nums []int, target int) []int {
// 	m := make(map[int]int)
// 	for i, num := range nums {
// 		if j, ok := m[target-num]; ok {
// 			return []int{j, i}
// 		}
// 		m[num] = i
// 	}
// 	return []int{}
// }
