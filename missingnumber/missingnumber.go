package missingnumber

import "sort"

func Solutions(A []int) int {
	sort.Ints(A)
	smallest := 1

	B := map[int]bool{}

	// Eliminate negative value
	for i := 0; i < len(A); i++ {
		if A[i] > 0 {
			B[A[i]] = true
		}
	}

	// Check the missing int in the sequences
	for i := 0; i < len(A); i++ {
		if B[i+1] == false {
			smallest = i + 1
			break
		} else {
			smallest = A[i] + 1
		}
	}

	return smallest

}
