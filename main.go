package main

import (
	"fmt"
	"mysort"
	"search"
)

func main() {
	// classicAlgorithms.PhilosophersDinner()
	// s1, s2 := "GCTAC", "CTCA"
	// fmt.Printf("string1 %s, string2 is %s\n", s1, s2)
	// distance, ops := dynamicProgramming.MinEditDistance(s1, s2)
	// fmt.Printf("distance is %d\n", distance)
	// fmt.Println(ops)

	fmt.Println(mysort.InsertionSort([]int{3, 2, 5, 1, 9, 8, 7, 4, 6}))
	fmt.Println(mysort.SelectionSort([]int{3, 2, 5, 1, 9, 8, 7, 4, 6}))
	fmt.Println(mysort.HeapSort([]int{3, 2, 5, 1, 9, 8, 7, 4, 6}))
	fmt.Println(mysort.QuickSort([]int{3, 2, 5, 1, 9, 8, 7, 4, 6}))
	fmt.Println(mysort.BubbleSort([]int{3, 2, 5, 1, 9, 8, 7, 4, 6}))

	fmt.Println(search.BinarySearch([]int{1, 3, 5, 6, 7}, 6))
}
