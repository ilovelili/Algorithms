package sort

import (
	"algorithmutil"
)

// MergeSort implements merge sort algorithm
func (sort *Sorter) MergeSort(arrayzor []algorithmutil.IComparable) []algorithmutil.IComparable {
	if len(arrayzor) <= 1 {
		return arrayzor
	}

	mid := len(arrayzor) / 2
	left := arrayzor[:mid]
	right := arrayzor[mid:]

	left = sort.MergeSort(left)
	right = sort.MergeSort(right)

	return sort.merge(left, right)
}

func (sort *Sorter) merge(left, right []algorithmutil.IComparable) []algorithmutil.IComparable {
	var result []algorithmutil.IComparable
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if sort.Less(left[0], right[0]) {
				result = append(result, left[0])
				left = left[1:]
			} else {
				result = append(result, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}
