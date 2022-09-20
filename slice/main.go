package main

import "fmt"

func main() {
	sliceInts := []int{1, 2, 3}
	sliceInts1 := make([]int, 10, 10)
	changeSlice(sliceInts)
	changeSlice(sliceInts1)

	fmt.Printf("sliceInts: %v\n", sliceInts)
	fmt.Printf("sliceInts1: %v\n", sliceInts1)
	changeAppend(sliceInts1)
}

func changeAppend(sliceInts1 []int) {
	sliceInts1 = append(sliceInts1, 3)
	sliceInts1 = append(sliceInts1, 23)
	fmt.Printf("changeAppend.sliceInts1: %v\n", sliceInts1)
}
func changeSlice(s []int) {
	s[0] = 3
	s[1] = 3
	s[2] = 3
}
