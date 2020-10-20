package main

import "sort"

//func sortedSquares(A []int) []int {
//	result := []int{}
//	for _, value := range A {
//		result = append(result, value*value)
//	}
//	sort.Ints(result)
//	return result
//}
func sortedSquares(A []int) []int {

	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}
