package misc

/*
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"

Input: str1 = "LEET", str2 = "CODE"
Output: ""
*/

func GcdOfStrings(str1 string, str2 string) string {
	// If str1 + str2 != str2 + str1, no GCD exists
	if str1+str2 != str2+str1 {
		return ""
	}

	// Find GCD of lengths using Euclidean algorithm
	gcdLength := gcd(len(str1), len(str2))

	// Return substring of str1 up to gcdLength
	return str1[:gcdLength]
}

// Helper function to compute GCD of two integers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
