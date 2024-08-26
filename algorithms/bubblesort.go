package algorithms

import "cmp"

func BubbleSort[T cmp.Ordered](values []T) []T {
	for i, _ := range values {
		intermediate := values[1:]
		for j, _ := range intermediate {
			if values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}

	return values

}
