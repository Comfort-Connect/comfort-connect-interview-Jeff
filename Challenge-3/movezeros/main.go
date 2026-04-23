// Given a slice of integers arr, move all the zeros to the end of the slice
// while maintaining the relative order of all non-zero elements.
//
// Examples:
//
//	Input:  [1, 2, 0, 4, 3, 0, 5, 0]
//	Output: [1, 2, 4, 3, 5, 0, 0, 0]
//
//	Input:  [10, 20, 30]
//	Output: [10, 20, 30]
//
//	Input:  [0, 0]
//	Output: [0, 0]
package main

import "fmt"

// MoveZerosToEnd moves all zeros to the end of the slice in-place while
// maintaining the relative order of non-zero elements.
func MoveZerosToEnd(arr []int) {
	// TODO: Implement your solution here
}

func main() {
	// Test Case 1
	arr1 := []int{1, 2, 0, 4, 3, 0, 5, 0}
	fmt.Printf("Input:  %v\n", arr1)
	MoveZerosToEnd(arr1)
	fmt.Printf("Output: %v\n", arr1)
	fmt.Println("Expected: [1 2 4 3 5 0 0 0]")
	fmt.Println()

	// Test Case 2
	arr2 := []int{10, 20, 30}
	fmt.Printf("Input:  %v\n", arr2)
	MoveZerosToEnd(arr2)
	fmt.Printf("Output: %v\n", arr2)
	fmt.Println("Expected: [10 20 30]")
	fmt.Println()

	// Test Case 3
	arr3 := []int{0, 0}
	fmt.Printf("Input:  %v\n", arr3)
	MoveZerosToEnd(arr3)
	fmt.Printf("Output: %v\n", arr3)
	fmt.Println("Expected: [0 0]")
}
