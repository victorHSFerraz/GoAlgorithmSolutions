package main

import "fmt"

func main() {
	isAnagram("a", "ab")
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		fmt.Println("false")
		return false
	}
	runeMapS := make(map[rune]int)
	runeMapT := make(map[rune]int)
	for _, s := range s {
		runeMapS[s]++
	}
	for _, t := range t {
		runeMapT[t]++
	}
	for _, r := range s {
		if runeMapS[r] != runeMapT[r] {
			fmt.Println("false")
			return false
		}
	}

	fmt.Println("true")
	return true
}
