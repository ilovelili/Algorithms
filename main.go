package main

import "dynamicProgramming"
import "fmt"

func main() {
	// classicAlgorithms.PhilosophersDinner()
	s1, s2 := "GCTAC", "CTCA"
	fmt.Printf("string1 %s, string2 is %s\n", s1, s2)
	distance, ops := dynamicProgramming.MinEditDistance(s1, s2)
	fmt.Printf("distance is %d\n", distance)
	fmt.Println(ops)
}
