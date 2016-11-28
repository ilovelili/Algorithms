package sorts

import (
	"github.com/ilovelili/Algorithms/algorithmutil"
)

// SelectionSort implements selection sort algorithm
func (sort *Sorter) SelectionSort(arrayzor []algorithmutil.IComparable) []algorithmutil.IComparable {
	count := len(arrayzor)

	for i := 0; i < count; i++ {
		min := i

		for j := i + 1; j < count; j++ {
			if sort.Less(arrayzor[j], arrayzor[min]) {
				min = j
			}
		}

		sort.Swap(arrayzor, i, min)
	}

	return arrayzor
}
