package sort

import (
	"../container"

	"algorithmutil"
)

type priorityqueue interface {
	comp(a []algorithmutil.IComparable, i, j int) bool
}

type heap struct {
	priorityqueue
	pq []algorithmutil.IComparable
	N  int
}

// Init initilize a heap
func (heap *heap) Init(capacity int) {
	heap.pq = make([]algorithmutil.IComparable, capacity+1)
	heap.N = 0
}

// IsEmpty heap is empty or not
func (heap *heap) IsEmpty() bool {
	return heap.N == 0
}

// Size size of the heap
func (heap *heap) Size() int {
	return heap.N
}

// Push
func (heap *heap) Push(element interface{}) {
	if e, ok := element.(algorithmutil.IComparable); ok {
		heap.pq[heap.N] = e
		heap.swim(heap.pq, heap.N)
	} else {
		panic("element is not Comparable")
	}
}

// Pop
func (heap *heap) Pop() *container.Item {
	item := &container.Item{Next: nil, Value: heap.pq[1]}
	heap.swap(heap.pq, 1, heap.N)
	heap.N--
	heap.pq[heap.N+1] = nil
	heap.sink(heap.pq, 1, heap.N)
	return item
}

// MaxHeap heap stores max value
type MaxHeap struct {
	heap
}

// NewMaxHeap create a new max heap
func NewMaxHeap(maxN int) *MaxHeap {
	this := &MaxHeap{}
	this.priorityqueue = this
	this.Init(maxN)
	return this
}

// MinHeap heap stores min value
type MinHeap struct {
	heap
}

// NewMinHeap create a new min heap
func NewMinHeap(maxN int) *MinHeap {
	this := &MinHeap{}
	this.priorityqueue = this
	this.Init(maxN)
	return this
}

func (heap *MaxHeap) comp(a []algorithmutil.IComparable, i, j int) bool {
	return a[i].CompareTo(a[j]) < 0
}

func (heap *heap) sink(a []algorithmutil.IComparable, k, N int) {
	for 2*k <= N {
		j := 2 * k
		if j < N && heap.priorityqueue.comp(a, j, j+1) {
			j++
		}
		if !heap.priorityqueue.comp(a, k, j) {
			break
		}
		heap.swap(a, k, j)
		k = j
	}
}

func (heap *heap) swim(pq []algorithmutil.IComparable, k int) {
	for k > 1 && heap.priorityqueue.comp(pq, k/2, k) {
		heap.swap(pq, k/2, k)
		k = k / 2
	}
}

func (heap *heap) swap(pq []algorithmutil.IComparable, i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
