package mysort

import "math/rand"

// QuickSort quick sort
func QuickSort(arr []int) []int {
	length := len(arr)

	if length <= 1 {
		return arr
	}

	pivot := arr[rand.Intn(length)]

	less := make([]int, 0, length)
	midian := make([]int, 0, length)
	more := make([]int, 0, length)

	for _, item := range arr {
		switch {
		case item < pivot:
			less = append(less, item)
		case item == pivot:
			midian = append(midian, item)
		case item > pivot:
			more = append(more, item)
		}
	}

	less, more = QuickSort(less), QuickSort(more)
	less = append(less, midian...)
	less = append(less, more...)

	return less
}
