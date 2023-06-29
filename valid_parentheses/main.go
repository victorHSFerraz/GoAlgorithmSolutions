// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()[]{}"))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	bracketMap := map[rune]rune{')': '(', ']': '[', '}': '{'}
	stack := make([]rune, 0)

	for _, ch := range s {
		if val, exists := bracketMap[ch]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != val {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return len(stack) == 0
}
