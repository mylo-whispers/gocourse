package Basics

// func main() {
// 	// BASIC RANGE OPERATIONS
// 	// The range keyword is used to iterate over elements in various data structures

// 	// 1. Range over slices/arrays
// 	fruits := []string{"apple", "banana", "cherry", "date"}

// 	fmt.Println("Iterating over a slice:")
// 	for index, fruit := range fruits {
// 		// index: position in the slice (0, 1, 2, ...)
// 		// fruit: value at that position
// 		fmt.Printf("Index %d: %s\n", index, fruit)
// 	}

// 	// 2. Range over maps
// 	inventory := map[string]int{
// 		"apples":  10,
// 		"oranges": 5,
// 		"bananas": 8,
// 	}

// 	fmt.Println("\nIterating over a map:")
// 	for key, value := range inventory {
// 		// key: map key (string)
// 		// value: map value (int)
// 		fmt.Printf("%s: %d\n", key, value)
// 	}

// 	// 3. Range over strings
// 	message := "Hello, 世界" // String with ASCII and Unicode characters

// 	fmt.Println("\nIterating over a string:")
// 	for index, char := range message {
// 		// index: byte position
// 		// char: rune (Unicode code point)
// 		fmt.Printf("Position %d: %c (Unicode: %U)\n", index, char, char)
// 	}

// 	// 4. Range over channels
// 	ch := make(chan string, 3) // Buffered channel with capacity 3
// 	ch <- "first"
// 	ch <- "second"
// 	ch <- "third"
// 	close(ch) // Close channel to indicate no more values

// 	fmt.Println("\nIterating over a channel:")
// 	for value := range ch {
// 		// value: received from channel
// 		fmt.Printf("Received: %s\n", value)
// 	}
// }

// func main() {
// 	// RANGE WITH DIFFERENT DATA STRUCTURES
// 	// Demonstrating range with various collection types

// 	// 1. Range with arrays (fixed size)
// 	coordinates := [3]float64{40.7128, -74.0060, 100.5} // New York coordinates with altitude
// 	fmt.Println("Array iteration:")
// 	for i, coord := range coordinates {
// 		fmt.Printf("Coordinate %d: %.4f\n", i, coord)
// 	}

// 	// 2. Range with slices (dynamic size)
// 	temperatures := []float32{22.5, 23.1, 24.8, 25.3, 26.0}
// 	fmt.Println("\nSlice iteration:")
// 	for i, temp := range temperatures {
// 		fmt.Printf("Day %d: %.1f°C\n", i+1, temp)
// 	}

// 	// 3. Range with maps (key-value pairs)
// 	studentAges := map[string]int{
// 		"Alice":   25,
// 		"Bob":     22,
// 		"Charlie": 23,
// 		"Diana":   24,
// 	}
// 	fmt.Println("\nMap iteration:")
// 	for name, age := range studentAges {
// 		fmt.Printf("%s is %d years old\n", name, age)
// 	}

// 	// 4. Range with strings (rune by rune)
// 	greeting := "Hello, 世界!" // Mix of ASCII and Unicode
// 	fmt.Println("\nString iteration (showing runes):")
// 	for position, char := range greeting {
// 		fmt.Printf("Position %d: '%c' (Unicode: U+%04X)\n", position, char, char)
// 	}

// 	// 5. Range with custom type (slice of structs)
// 	type Product struct {
// 		ID    string
// 		Name  string
// 		Price float64
// 	}

// 	products := []Product{
// 		{"P001", "Laptop", 999.99},
// 		{"P002", "Mouse", 25.50},
// 		{"P003", "Keyboard", 75.00},
// 	}

// 	fmt.Println("\nStruct slice iteration:")
// 	for _, product := range products {
// 		fmt.Printf("Product %s: %s - $%.2f\n", product.ID, product.Name, product.Price)
// 	}
// }

// import "fmt"

// func main() {
// 	// RANGE PATTERNS AND IDIOMS
// 	// Common patterns when using range in Go

// 	numbers := []int{10, 20, 30, 40, 50}

// 	// 1. Ignoring index using blank identifier
// 	fmt.Println("Values only (ignoring index):")
// 	for _, value := range numbers {
// 		fmt.Printf("%d ", value)
// 	}
// 	fmt.Println()

// 	// 2. Ignoring value using blank identifier
// 	fmt.Println("Indexes only (ignoring values):")
// 	for index := range numbers {
// 		fmt.Printf("%d ", index)
// 	}
// 	fmt.Println()

// 	// 3. Traditional for loop equivalent
// 	fmt.Println("Traditional for loop (equivalent to range):")
// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Printf("%d ", numbers[i])
// 	}
// 	fmt.Println()

// 	// 4. Range with pointer elements
// 	ptrNumbers := []*int{&numbers[0], &numbers[1], &numbers[2]}
// 	fmt.Println("\nRange with pointer elements:")
// 	for i, ptr := range ptrNumbers {
// 		fmt.Printf("Index %d: Pointer %p -> Value %d\n", i, ptr, *ptr)
// 	}

// 	// 5. Modifying slice elements during iteration
// 	fmt.Println("\nModifying elements during iteration:")
// 	for i := range numbers {
// 		numbers[i] *= 2 // Double each value
// 	}
// 	fmt.Println("Modified slice:", numbers)

// 	// 6. Range with nested slices (2D array)
// 	matrix := [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}

// 	fmt.Println("\nIterating over 2D slice:")
// 	for rowIndex, row := range matrix {
// 		for colIndex, value := range row {
// 			fmt.Printf("matrix[%d][%d] = %d\n", rowIndex, colIndex, value)
// 		}
// 	}

// 	// 7. Early exit from range loop
// 	fmt.Println("\nEarly exit from range loop:")
// 	for i, value := range numbers {
// 		if value > 80 {
// 			fmt.Printf("Stopping at index %d (value %d > 80)\n", i, value)
// 			break
// 		}
// 		fmt.Printf("Processing index %d: %d\n", i, value)
// 	}
// }

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// PRACTICAL RANGE EXAMPLES
// 	// Real-world use cases for range

// 	// 1. Processing lines from a file (simulated)
// 	document := `Line 1: Welcome to Go programming
// Line 2: Range makes iteration easy
// Line 3: Let's process each line
// Line 4: And extract useful information`

// 	lines := strings.Split(document, "\n")
// 	fmt.Println("Document line processing:")
// 	for lineNumber, line := range lines {
// 		fmt.Printf("Line %d: %s\n", lineNumber+1, line)
// 	}

// 	// 2. Word frequency counter
// 	text := "go python go java python go rust java go"
// 	words := strings.Fields(text)
// 	wordCount := make(map[string]int)

// 	for _, word := range words {
// 		wordCount[word]++
// 	}

// 	fmt.Println("\nWord frequency:")
// 	for word, count := range wordCount {
// 		fmt.Printf("%s: %d\n", word, count)
// 	}

// 	// 3. Finding maximum value
// 	temperatures := []float64{22.1, 23.5, 21.8, 25.3, 24.9}
// 	maxTemp := temperatures[0]
// 	maxDay := 0

// 	for day, temp := range temperatures {
// 		if temp > maxTemp {
// 			maxTemp = temp
// 			maxDay = day
// 		}
// 	}
// 	fmt.Printf("\nHottest day: Day %d (%.1f°C)\n", maxDay+1, maxTemp)

// 	// 4. Data transformation
// 	names := []string{"alice", "bob", "charlie", "diana"}
// 	capitalized := make([]string, len(names))

// 	for i, name := range names {
// 		capitalized[i] = strings.ToUpper(name)
// 	}
// 	fmt.Println("\nCapitalized names:", capitalized)

// 	// 5. Filtering elements
// 	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	var evenNumbers []int

// 	for _, num := range numbers {
// 		if num%2 == 0 {
// 			evenNumbers = append(evenNumbers, num)
// 		}
// 	}
// 	fmt.Println("Even numbers:", evenNumbers)

// 	// 6. Creating lookup maps
// 	products := []struct {
// 		ID   string
// 		Name string
// 	}{
// 		{"P001", "Laptop"},
// 		{"P002", "Mouse"},
// 		{"P003", "Keyboard"},
// 	}

// 	productMap := make(map[string]string)
// 	for _, product := range products {
// 		productMap[product.ID] = product.Name
// 	}

// 	fmt.Println("\nProduct lookup map:")
// 	for id, name := range productMap {
// 		fmt.Printf("%s -> %s\n", id, name)
// 	}
// }

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// RANGE WITH CHANNELS
// 	// Using range to receive values from channels

// 	// 1. Range with buffered channel
// 	messages := make(chan string, 3)
// 	messages <- "Hello"
// 	messages <- "World"
// 	messages <- "!"
// 	close(messages) // Must close channel for range to work

// 	fmt.Println("Messages from buffered channel:")
// 	for msg := range messages {
// 		fmt.Println(msg)
// 	}

// 	// 2. Range with channel and goroutine
// 	dataChannel := make(chan int)

// 	// Producer goroutine
// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			dataChannel <- i * 10
// 			time.Sleep(100 * time.Millisecond) // Simulate work
// 		}
// 		close(dataChannel) // Close when done sending
// 	}()

// 	// Consumer using range
// 	fmt.Println("\nData from goroutine:")
// 	for value := range dataChannel {
// 		fmt.Printf("Received: %d\n", value)
// 	}

// 	// 3. Multiple producers, single consumer
// 	workerCount := 3
// 	results := make(chan string, workerCount)

// 	// Start multiple worker goroutines
// 	for i := 1; i <= workerCount; i++ {
// 		go func(workerID int) {
// 			for task := 1; task <= 2; task++ {
// 				results <- fmt.Sprintf("Worker %d completed task %d", workerID, task)
// 				time.Sleep(time.Duration(workerID*100) * time.Millisecond)
// 			}
// 		}(i)
// 	}

// 	// Wait for all workers to complete
// 	go func() {
// 		time.Sleep(1000 * time.Millisecond) // Wait a bit
// 		close(results)                      // Close channel after all workers finish
// 	}()

// 	// Process results
// 	fmt.Println("\nWorker results:")
// 	for result := range results {
// 		fmt.Println(result)
// 	}
// }

// import "fmt"

// func main() {
// 	// ADVANCED RANGE TECHNIQUES
// 	// More complex patterns and use cases

// 	// 1. Range with custom collection type
// 	type StringCollection []string
// 	names := StringCollection{"Alice", "Bob", "Charlie", "Diana"}

// 	fmt.Println("Custom collection iteration:")
// 	for i, name := range names {
// 		fmt.Printf("%d: %s\n", i, name)
// 	}

// 	// 2. Range with array pointers
// 	numbers := [5]int{1, 2, 3, 4, 5}
// 	numbersPtr := &numbers // Pointer to array

// 	fmt.Println("\nArray pointer iteration:")
// 	for i, value := range numbersPtr { // Range automatically dereferences
// 		fmt.Printf("Index %d: %d\n", i, value)
// 	}

// 	// 3. Range with function returning slice
// 	fmt.Println("\nRange with function return:")
// 	for i, value := range generateFibonacci(10) { // Function returns slice
// 		fmt.Printf("Fibonacci %d: %d\n", i+1, value)
// 	}

// 	// 4. Parallel slice iteration
// 	names2 := []string{"Alice", "Bob", "Charlie"}
// 	ages := []int{25, 30, 35}

// 	fmt.Println("\nParallel slice iteration:")
// 	for i := range names2 {
// 		if i < len(ages) { // Ensure we don't go out of bounds
// 			fmt.Printf("%s is %d years old\n", names2[i], ages[i])
// 		}
// 	}

// 	// 5. Range with map and conditional processing
// 	inventory := map[string]int{
// 		"apples":  15,
// 		"oranges": 8,
// 		"bananas": 12,
// 		"grapes":  20,
// 	}

// 	fmt.Println("\nLow stock items:")
// 	for item, quantity := range inventory {
// 		if quantity < 10 {
// 			fmt.Printf("LOW STOCK: %s (%d left)\n", item, quantity)
// 		}
// 	}

// 	// 6. Range with index manipulation
// 	data := []int{10, 20, 30, 40, 50}
// 	fmt.Println("\nIterating with index manipulation:")
// 	for i := range data {
// 		if i > 0 {
// 			// Calculate difference with previous element
// 			diff := data[i] - data[i-1]
// 			fmt.Printf("Difference between index %d and %d: %d\n", i, i-1, diff)
// 		}
// 	}
// }

// func generateFibonacci(n int) []int {
// 	if n <= 0 {
// 		return []int{}
// 	}
// 	if n == 1 {
// 		return []int{0}
// 	}

// 	fib := make([]int, n)
// 	fib[0], fib[1] = 0, 1

// 	for i := 2; i < n; i++ {
// 		fib[i] = fib[i-1] + fib[i-2]
// 	}

// 	return fib
// }

import (
	"fmt"
	"strconv"
)

func main() {
	// ERROR HANDLING WITH RANGE
	// Dealing with potential errors during iteration

	// 1. Processing user input with validation
	userInputs := []string{"10", "20", "thirty", "40", "50"}
	var validNumbers []int
	var invalidInputs []string

	fmt.Println("Processing user inputs:")
	for i, input := range userInputs {
		if num, err := strconv.Atoi(input); err == nil {
			validNumbers = append(validNumbers, num)
			fmt.Printf("Input %d: Valid number %d\n", i+1, num)
		} else {
			invalidInputs = append(invalidInputs, input)
			fmt.Printf("Input %d: Invalid input '%s'\n", i+1, input)
		}
	}

	fmt.Println("\nValid numbers:", validNumbers)
	fmt.Println("Invalid inputs:", invalidInputs)

	// 2. Safe map access with error handling
	config := map[string]string{
		"host":     "localhost",
		"port":     "8080",
		"timeout":  "30",
		"database": "mydb",
	}

	requiredKeys := []string{"host", "port", "timeout", "database", "username"}

	fmt.Println("\nConfig validation:")
	for _, key := range requiredKeys {
		if value, exists := config[key]; exists {
			fmt.Printf("✓ %s: %s\n", key, value)
		} else {
			fmt.Printf("✗ Missing required config: %s\n", key)
		}
	}

	// 3. Processing with recovery from panics
	riskyOperations := []func() int{
		func() int { return 100 / 2 }, // Safe
		func() int { return 100 / 0 }, // Will panic
		func() int { return 200 / 4 }, // Safe
	}

	fmt.Println("\nSafe operation execution:")
	for i, op := range riskyOperations {
		// Use anonymous function with defer/recover for each operation
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Operation %d failed: %v\n", i+1, r)
				}
			}()

			result := op()
			fmt.Printf("Operation %d succeeded: %d\n", i+1, result)
		}()
	}

	// 4. Batch processing with error collection
	dataToProcess := []string{"1", "2", "3", "four", "5"}
	var processingErrors []error

	fmt.Println("\nBatch processing with error collection:")
	for i, item := range dataToProcess {
		if _, err := strconv.Atoi(item); err != nil {
			processingErrors = append(processingErrors,
				fmt.Errorf("item %d ('%s'): %v", i, item, err))
		}
	}

	if len(processingErrors) > 0 {
		fmt.Println("Processing errors:")
		for _, err := range processingErrors {
			fmt.Println("  -", err)
		}
	} else {
		fmt.Println("All items processed successfully!")
	}
}
