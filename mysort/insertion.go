package mysort

// InsertionSort best O(n) for already sorted array, worst O(n^2) for reverse sorted array. Good for nearly sorted array
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}

	return arr
}
