package search

// BinarySearch binary search
func BinarySearch(arr []int, target int) (index int) {
	startIndex := 0
	endIndex := len(arr) - 1

	for startIndex <= endIndex {
		medianIndex := (endIndex + startIndex) / 2
		median := arr[medianIndex]

		if target < median {
			endIndex = medianIndex - 1
		} else if target > median {
			startIndex = medianIndex + 1
		} else {
			return medianIndex
		}
	}

	if startIndex == len(arr) || arr[startIndex] != target {
		return -1
	}

	return startIndex
}
