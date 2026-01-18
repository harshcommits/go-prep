package misc

/*
Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
*/

func MergeStrings(s1 string, s2 string) string {

	merged := ""

	n := max(len(s1), len(s2))

	for i := 0; i < n; i++ {
		if i < len(s1) && i < len(s2) {
			merged += string(s1[i]) + string(s2[i])
		}
	}

	return merged

}
