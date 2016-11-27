package sort

import (
	"github.com/ilovelili/Algorithms/algorithmutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("selection sort", func() {
	var sorter *Sorter

	arrayzor := algorithmutil.MakeIntegerSlice([]int{3, 2, 8, 1, 4, 6, 5, 9, 7, 0})

	It("should sort correctly", func() {
		actual := sorter.SelectionSort(arrayzor)
		expected := algorithmutil.MakeIntegerSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

		Expect(actual).To(Equal(expected))
	})
})
