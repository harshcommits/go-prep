package algorithms

import (
	"fmt"
)

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

func Iterate[T any](values []T) {

	for _, value := range values {
		fmt.Println(value)
	}

	fmt.Println(values[2])

}
