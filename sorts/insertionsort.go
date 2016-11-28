package sorts

import (
	"github.com/ilovelili/Algorithms/algorithmutil"
)

// InsertionSort implements insertion sort algorithm
func (sorter *Sorter) InsertionSort(arrayzor []algorithmutil.IComparable) []algorithmutil.IComparable {
	count := len(arrayzor)

	for i := 1; i < count; i++ {
		for j := i; j > 0 && sorter.Less(arrayzor[j], arrayzor[j-1]); j-- {
			sorter.Swap(arrayzor, j, j-1)
		}
	}

	return arrayzor
}
