package misc

import (
	"fmt"

	"github.com/harshcommits/go-prep/ds"
)

func TwoSum(nums []int, target int) []int {

	/*
		Input: nums = [2,7,11,15], target = 9
		Output: [0,1]
		Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	*/

	n := len(nums)
	sum := make([]int, 2)

	for i := 0; i < n; i++ {
		for j := n - 1; j > 0; j-- {
			if i < j {
				sumValue := nums[i] + nums[j]
				if sumValue == target {
					sum[i] = j
					return []int{i, sum[i]}
				}
			}
		}
	}

	return []int{-1}

}

func TwoSumOptimized(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[num] = i
	}

	return []int{}

}

func BinarySearch(num int, nums []int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2
		fmt.Printf("Searching for %d, mid index is %d, mid value is %d\n", num, mid, nums[mid])

		if num == nums[mid] {
			return mid
		} else if num > nums[mid] {
			low = mid + 1 // Search right half
		} else {
			high = mid - 1 // Search left half
		}
	}

	return -1 // Not found
}

func ValidPalindrome(value string) bool {

	stringAsRunes := []rune(value)
	for i, j := 0, len(stringAsRunes)-1; i < j; i, j = i+1, j-1 {
		stringAsRunes[i], stringAsRunes[j] = stringAsRunes[j], stringAsRunes[i]
	}

	if string(stringAsRunes) == value {
		return true
	} else {
		return false
	}

}

func DepthOfBinaryTree(t *ds.Tree) int {
	return depthOfNode(t.GetRoot())
}

func depthOfNode(node *ds.TreeNode) int {
	if node == nil {
		return 0
	}

	leftDepth := depthOfNode(node.GetLeft())
	rightDepth := depthOfNode(node.GetRight())

	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func LongestSubstringWithoutRepeatingChars(s string) int {

	runeOfString := []rune(s)
	visitedChars := make(map[rune]int)
	longestSubstring := 0

	for i, j := 0, 0; j < len(runeOfString); j++ {

		char := runeOfString[j]

		if lastIndex, exists := visitedChars[char]; exists && lastIndex >= i {
			i = lastIndex + 1
		}

		visitedChars[char] = j
		longestSubstring = max(longestSubstring, j-i+1)
	}

	return longestSubstring
}

func CharacterReplacement(s string, k int) int {

	charCount := make(map[rune]int)
	maxCount := 0 // Count of most frequent character in window
	maxLength := 0

	for i, j := 0, 0; j < len(s); j++ {
		char := rune(s[j])
		charCount[char]++
		maxCount = max(maxCount, charCount[char])

		// Window length - most frequent char count = replacements needed
		// If replacements needed > k, shrink window
		windowLength := j - i + 1
		if windowLength-maxCount > k {
			charCount[rune(s[i])]--
			i++
		}

		maxLength = max(maxLength, j-i+1)
	}

	return maxLength

}
