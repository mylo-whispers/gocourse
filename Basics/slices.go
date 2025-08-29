package Basics

import (
	"fmt"
	"sort"
)

func main() {
	// SLICE FUNCTIONS AND TECHNIQUES
	// Common operations and patterns

	numbers := []int{5, 2, 8, 1, 9, 3}
	fmt.Println("Original:", numbers)

	// Sorting slices
	sort.Ints(numbers) // Sorts in place
	fmt.Println("Sorted:", numbers)

	// Checking if slice is sorted
	fmt.Println("Is sorted:", sort.IntsAreSorted(numbers))

	// Searching in sorted slice
	index := sort.SearchInts(numbers, 8)
	fmt.Println("Index of 8:", index)

	// Reversing a slice
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i] // Swap elements
	}
	fmt.Println("Reversed:", numbers)

	// Filtering elements (creating new slice)
	var evenNumbers []int
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Println("Even numbers:", evenNumbers)

	// Mapping (transforming elements)
	var squaredNumbers []int
	for _, num := range numbers {
		squaredNumbers = append(squaredNumbers, num*num)
	}
	fmt.Println("Squared numbers:", squaredNumbers)

	// Reducing (aggregating elements)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// Finding element
	found := false
	for _, num := range numbers {
		if num == 5 {
			found = true
			break
		}
	}
	fmt.Println("Contains 5:", found)

	// Removing elements by index
	// Create new slice excluding element at index 2
	indexToRemove := 2
	numbers = append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)
	fmt.Println("After removing index 2:", numbers)

	// Inserting elements
	// Insert 100 at index 1
	indexToInsert := 1
	numbers = append(numbers[:indexToInsert], append([]int{100}, numbers[indexToInsert:]...)...)
	fmt.Println("After inserting 100 at index 1:", numbers)
}
