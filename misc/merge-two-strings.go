package misc

/*
Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
*/

func MergeStrings(word1 string, word2 string) string {

	merged := ""

	n := max(len(word1), len(word2))

	for i := 0; i < n; i++ {
		if i < len(word1) && i < len(word2) {
			merged += string(word1[i]) + string(word2[i])
		} else if i < len(word1) {
			merged += string(word1[i])
		} else if i < len(word2) {
			merged += string(word2[i])
		}
	}

	return merged

}

// Written by Co-pilot
func MergeStringsAlternate(word1 string, word2 string) string {

	merged := ""

	i, j := 0, 0

	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			merged += string(word1[i])
			i++
		}
		if j < len(word2) {
			merged += string(word2[j])
			j++
		}
	}

	return merged

}

// Optimized version from LeetCode
func MergeAlternately(word1 string, word2 string) string {
	ans := make([]byte, 0, len(word1)+len(word2))

	for w1, w2 := 0, 0; w1 < len(word1) || w2 < len(word2); w1, w2 = w1+1, w2+1 {
		var sw1, sw2 byte

		if w1 < len(word1) {
			sw1 = word1[w1]
		}

		if w2 < len(word2) {
			sw2 = word2[w2]
		}

		if sw1 != 0 {
			ans = append(ans, sw1)
		}
		if sw2 != 0 {
			ans = append(ans, sw2)
		}
	}
	return string(ans)
}
