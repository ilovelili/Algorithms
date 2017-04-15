package main

import (
	"fmt"
	"mysort"
)

func main() {
	// classicAlgorithms.PhilosophersDinner()
	// s1, s2 := "GCTAC", "CTCA"
	// fmt.Printf("string1 %s, string2 is %s\n", s1, s2)
	// distance, ops := dynamicProgramming.MinEditDistance(s1, s2)
	// fmt.Printf("distance is %d\n", distance)
	// fmt.Println(ops)

	arr := []int{3, 2, 5, 1, 9, 8, 7, 4, 6}
	fmt.Println(mysort.InsertionSort(arr))
	fmt.Println(mysort.SelectionSort(arr))
}
