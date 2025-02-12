// run code using `go run basics.go`
package main

import (
	"fmt"
	"sort"
)

/*
simpleBinarySearch performs a binary search on a sorted slice of integers to determine whether a given value exists.

### Function Signature:
- `func simpleBinarySearch(arrSlice []int, searchValue *int) bool`
  - `arrSlice []int`: A slice of integers passed by value which means modifications inside the function will affect the original slice
  - `searchValue *int`: A pointer to an integer. The `*int` indicates that searchValue is a pointer, meaning it stores the memory address of an integer and not the actual value

- The function **returns a `bool`** indicating whether `searchValue` exists in `arrSlice`.

### Key Go Syntax Features:
- Dereferencing (`*searchValue`): Since `searchValue` is a pointer, we use `*searchValue` to access the actual integer it points to.
- For Loop (`for start <= end`): Implements a binary search loop where `start` and `end` define the search range.
- Midpoint Calculation (`mid := (start + end) / 2`):
  - The `/` operator in Go automatically floors the result of integer division.

- Conditional Checks (`if, else if, else`):
  - Compares `midValue` (the value at `arrSlice[mid]`) with `*searchValue`.
  - Adjusts `start` or `end` to refine the search range.

- Returns `true` or `false` depending on whether the value is found.
*/
func simpleBinarySearch(arrSlice []int, searchValue *int) bool {
	end := len(arrSlice) - 1
	start := 0
	for start <= end {
		mid := (start + end) / 2
		midValue := arrSlice[mid]
		if midValue == *searchValue { // since searchValue is being passed as pointer we have to dereference with * to get the actual value and not the address
			return true
		} else if midValue > *searchValue {
			end = mid - 1
		} else if midValue < *searchValue {
			start = mid + 1
		}
	}
	return false
}

func main() {
	numbers := []int{5, 4, 2, 1, 7, 3, 2} // a slice of numberrs
	sort.Ints(numbers)                    // sorts slice in place
	searchFor := 4
	isValueInArrSlice := simpleBinarySearch(numbers, &searchFor) // Passing a pointer to searchFor
	fmt.Println("Is value in given array: ", isValueInArrSlice)
}
