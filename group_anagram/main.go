//Given an array of strings strs, group the anagrams together. You can return the answer in any order.

//An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

package main

import (
	"sort"
	"strings"
)

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}

func groupAnagrams(strs []string) [][]string {
	ans := make([][]string, 0)
	m := make(map[string][]string)

	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		sortedStr := strings.Join(s, "")

		m[sortedStr] = append(m[sortedStr], str)
	}

	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}
