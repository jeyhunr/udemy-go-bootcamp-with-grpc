package main

import "fmt"

func Slice() {
	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d to slice at index [%d][%d]\n", i+j, i, j)
		}
	}
	fmt.Println(twoD)
}
