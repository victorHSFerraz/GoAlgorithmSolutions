/* Given a string s, find the length of the longest
* substring
* without repeating characters.
 */

package main

func main() {
	lengthOfLongestSubstring("aab")
}

func lengthOfLongestSubstring(s string) int {
	var max int
	var characters []rune

	for _, char := range s {
		for i, existingChar := range characters {
			if existingChar == char {
				characters = characters[i+1:]
				break
			}
		}

		characters = append(characters, char)
		if len(characters) > max {
			max = len(characters)
		}
	}

	return max
}

//another way with map

// func lengthOfLongestSubstring(s string) int {
//     lastIndex := make(map[rune]int)
//     start, maxLength := 0, 0

//     for end, char := range s {
//         if last, exists := lastIndex[char]; exists && last >= start {
//             start = last + 1
//         }

//         if length := end - start + 1; length > maxLength {
//             maxLength = length
//         }

//         lastIndex[char] = end
//     }

//     return maxLength
// }
