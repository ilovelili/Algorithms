package goalgorithms

// BubbleSort Bubble Sort
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
