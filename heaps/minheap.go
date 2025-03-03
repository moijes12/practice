// A basic program to create a min heap from an array
package main

import "fmt"

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}

func buildHeap(A []int, index int) {
	l := left(index)
	r := right(index)
	var smallest int
	fmt.Printf("index = %d, left = %d, right = %d\n", index, l, r)
	for i := range A {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
	if l < len(A) && A[l] < A[index] {
		smallest = l
	} else {
		smallest = index
	}
	if r < len(A) && A[r] < A[smallest] {
		smallest = r
	}
	if smallest != index {
		temp := A[index]
		A[index] = A[smallest]
		A[smallest] = temp
		buildHeap(A, smallest)
	}
}

func main() {
	A := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}
	for i := len(A)/2 - 1; i >= 0; i-- {
		buildHeap(A, i)
	}
}
