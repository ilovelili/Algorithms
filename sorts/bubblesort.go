// Package goalgorithms provides basic algorithm and data structure implementations
package goalgorithms

// BubbleSort implements bubble sort algorithm
func BubbleSort(arrayzor []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for i := 1; i < len(arrayzor); i++ {
			if arrayzor[i-1] > arrayzor[i] {
				arrayzor[i], arrayzor[i-1] = arrayzor[i-1], arrayzor[i]
				swapped = true
			}
		}
	}

	return arrayzor
}
