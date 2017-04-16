package mysort

// HeapSort Heap construction uses <= 2N compares and exchanges
// heapSort uses <= 2 N lg N compares and exchanges
// heapSort is optimal for both time and space, but :
// - Inner loop longer than quickSort's
// - Makes poor use of cache memory
// - Not stable
// worst: 2 N lg N , average: 2 N lg N , best: N lg N
// N log N guarantee , in place

// HeapSort heap sort
func HeapSort(arr []int) []int {
	heap := BuildHeap(arr)
	for length := len(arr); length > 1; length-- {
		heap.RemoveTop(arr, length)
	}

	return arr
}

// Heap max heap
type Heap struct {
}

// RemoveTop remove top
func (heap *Heap) RemoveTop(array []int, length int) {
	lastIndex := length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	heap.Heapify(array, 0, lastIndex)
}

// BuildHeap build max heap
func BuildHeap(array []int) *Heap {
	heap := &Heap{}
	for i := len(array) / 2; i >= 0; i-- {
		heap.Heapify(array, i, len(array))
	}
	return heap
}

// Heapify Max heap
func (heap *Heap) Heapify(array []int, root, length int) {
	max := root
	l, r := heap.Left(array, root), heap.Right(array, root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		heap.Heapify(array, max, length)
	}
}

// Left node
func (*Heap) Left(array []int, root int) int {
	return (root * 2) + 1
}

// Right right node
func (*Heap) Right(array []int, root int) int {
	return (root * 2) + 2
}
