package sorts

import (
	"github.com/ilovelili/Algorithms/algorithmutil"
)

// BubbleSort implements bubble sort algorithm
func (sort *Sorter) BubbleSort(arrayzor []algorithmutil.IComparable) []algorithmutil.IComparable {
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < len(arrayzor); i++ {
			if arrayzor[i-1].CompareTo(arrayzor[i]) > 0 {
				sort.Swap(arrayzor, i, i-1)
				swapped = true
			}
		}
	}

	return arrayzor
}
