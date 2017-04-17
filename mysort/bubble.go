package mysort

// BubbleSort Best Ω(n) Avg Θ(n^2) Worst O(n^2) Space O(n)
func BubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}
