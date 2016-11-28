package sorts

import (
	"github.com/ilovelili/Algorithms/algorithmutil"
)

// ISorter defines customized sorter interface
type ISorter interface {
	Less(v, w algorithmutil.IComparable) bool
	Swap(a []algorithmutil.IComparable, i, j int)
	IsSorted(a []algorithmutil.IComparable) bool
	Sort(a []algorithmutil.IComparable)
}

// Sorter ISorter instance
type Sorter struct {
}

// Less get the less one of two algorithmutil.IComparable elements
func (sorter *Sorter) Less(v, w algorithmutil.IComparable) bool {
	return v.CompareTo(w) < 0
}

// Swap swap two algorithmutil.IComparable elements
func (sorter *Sorter) Swap(a []algorithmutil.IComparable, i, j int) {
	a[j], a[i] = a[i], a[j]
}

// IsSorted check if an algorithmutil.IComparable array is sorted or not
func (sorter *Sorter) IsSorted(a []algorithmutil.IComparable) bool {
	for i := 1; i < len(a); i++ {
		if sorter.Less(a[i], a[i-1]) {
			return false
		}
	}

	return true
}
