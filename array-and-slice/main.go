package main

import (
	"golang.org/x/example/hello/reverse"

	"fmt"
)

func main() {
	fmt.Println(reverse.String("Hello"))

	SliceLengthAndCapacity()

}

func ArrayLength() {
	var arr = [...]int{1, 2}
	arr = append(arr, 3)

}

func SliceLengthAndCapacity() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	s = append(s, 7)
	s = append(s, 8)
	printSlice(s)

	s = append(s, 11) // Slice capacity doubled when its size has reached the max capacity.
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
