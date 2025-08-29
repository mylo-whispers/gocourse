package basics

import "fmt"

func main() {
	// ARRAYS VS. SLICES
	// Important differences and when to use each

	// Array: Fixed size, value type
	array := [3]int{1, 2, 3}

	// Slice: Dynamic size, reference type (more common in Go)
	slice := []int{1, 2, 3}

	fmt.Println("Array:", array, "Length:", len(array))
	fmt.Println("Slice:", slice, "Length:", len(slice))

	// Modifying a copy
	arrayCopy := array // Creates a complete copy
	sliceCopy := slice // Creates a reference to the same underlying array

	arrayCopy[0] = 100 // Doesn't affect original array
	sliceCopy[0] = 100 // Affects original slice!

	fmt.Println("After modification:")
	fmt.Println("Original array:", array) // Unchanged
	fmt.Println("Array copy:", arrayCopy) // Changed
	fmt.Println("Original slice:", slice) // Changed!
	fmt.Println("Slice copy:", sliceCopy) // Changed

	// Converting between arrays and slices
	// Array to slice (very common)
	sliceFromArray := array[:] // Creates slice that references the array
	fmt.Println("Slice from array:", sliceFromArray)

	// Modifying the slice affects the original array
	sliceFromArray[0] = 999
	fmt.Println("After slice modification, array:", array) // Changed!

	// Slice to array (Go 1.17+)
	// This creates a copy, not a reference
	arrayFromSlice := [3]int(slice) // Note: size must match
	fmt.Println("Array from slice:", arrayFromSlice)

	// When to use arrays vs slices:
	// - Use arrays when you need fixed size and value semantics
	// - Use slices for most cases (more flexible, dynamic size)
	// - Use arrays for mathematical operations where fixed size is important
	// - Use arrays when you need compile-time size checking
}
