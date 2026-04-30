package sort

import (
	"fmt"
	"sort"
)

func ExampleDesc() {
	arr := []int{5, 2, 8, 1, 3}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j] // descending
	})

	fmt.Println(arr) // Output: [8 5 3 2 1]
}

func ExampleAsc() {
	arr := []int{5, 2, 8, 1, 3}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j] // ascending
	})

	fmt.Println(arr) // Output: [1 2 3 5 8]
}


// slice.Sort method in go directly pass any slice and sort it in ascending order
func ExampleSort() {
	arr := []int{5, 2, 8, 1, 3}

	sort.Ints(arr) // sort in ascending order

	fmt.Println(arr) // Output: [1 2 3 5 8]

	arr2 := []string{"banana", "apple", "cherry"}

	sort.Strings(arr2) // sort in ascending order

	fmt.Println(arr2) // Output: [apple banana cherry]
}

