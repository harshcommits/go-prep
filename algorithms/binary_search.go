package algorithms

import "fmt"

func BinarySearch(num int, nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1

	for low <= high {
		mid := low + (high-low)/2

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

func Iterate[T any](values []T) {

	for _, value := range values {
		fmt.Println(value)
	}

	if len(values) > 2 {
		fmt.Println(values[2])
	}
}
