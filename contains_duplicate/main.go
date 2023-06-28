package main

func main() {
	containsDuplicate([]int{1, 2, 3, 1})
}

// func containsDuplicate(nums []int) bool {
// 	var list []int
// 	for _, num := range nums {
// 		for _, element := range list {
// 			if num == element {
// 				fmt.Println("contains")
// 				return true
// 			}
// 		}
// 		fmt.Println(num)
// 		list = append(list, num)
// 	}
// 	fmt.Println("not contains")
// 	return false
// }

//fast way
func containsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)
	for _, num := range nums {
		if _, exists := numMap[num]; exists {
			return true
		}
		numMap[num] = true
	}
	return false
}
