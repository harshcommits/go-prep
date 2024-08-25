package algorithms

import (
	"cmp"
	"fmt"
)

// using type Ordered since comparison operators don't work on slice, maps and functions; any and comparable restricts those types from use in generics
func BinarySearch[T cmp.Ordered](values []T, search T) bool {
	large := len(values) - 1
	low := 0

	for low <= large {
		mid := (large + low) / 2

		if values[mid] == search {
			return true
		}

		if values[mid] > search {
			large = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}

func Iterate[T any](values []T) {

	for _, value := range values {
		fmt.Println(value)
	}

	fmt.Println(values[2])

}
