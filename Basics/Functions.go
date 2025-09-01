package main // Declare this file as part of the main package (executable program)

// Import necessary packages for the examples
import (
	"errors" // For creating error values
	"fmt"    // For input/output operations
	"math"   // For mathematical functions and constants

	// For string conversion functions
	"strings" // For string manipulation functions
	"time"    // For time-related operations
)

// 1. BASIC FUNCTION
// Simple function that takes two float64 parameters and returns a float64
func calculateArea(length, width float64) float64 {
	return length * width // Calculate and return the area
}

// 2. MULTIPLE RETURN VALUES
// Function that returns two values: area and circumference of a circle
func calculateCircle(radius float64) (area float64, circumference float64) {
	area = math.Pi * radius * radius     // Calculate area using πr² formula
	circumference = 2 * math.Pi * radius // Calculate circumference using 2πr formula
	return area, circumference           // Explicitly return both values
}

// 3. NAMED RETURN VALUES
// Function with named return values - min, max, and average are declared in the signature
func calculateStatistics(numbers []int) (min, max, average int) {
	if len(numbers) == 0 { // Check if the slice is empty
		return 0, 0, 0 // Return zeros if no numbers provided
	}

	min = numbers[0] // Initialize min with first element
	max = numbers[0] // Initialize max with first element
	sum := 0         // Initialize sum to zero

	// Iterate through all numbers in the slice
	for _, num := range numbers {
		if num < min { // Update min if current number is smaller
			min = num
		}
		if num > max { // Update max if current number is larger
			max = num
		}
		sum += num // Add current number to sum
	}

	average = sum / len(numbers) // Calculate average
	return                       // Return all named values (min, max, average)
}

// 4. VARIADIC FUNCTION
// Function that accepts variable number of arguments using ... syntax
func calculateSum(numbers ...int) int {
	total := 0                    // Initialize total to zero
	for _, num := range numbers { // Iterate through all provided numbers
		total += num // Add each number to total
	}
	return total // Return the sum
}

// 5. FUNCTION AS PARAMETER
// Function that accepts another function as a parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b) // Call the provided operation function with a and b
}

// Helper function for addition - used with applyOperation
func add(a, b int) int {
	return a + b // Return sum of a and b
}

// Helper function for multiplication - used with applyOperation
func multiply(a, b int) int {
	return a * b // Return product of a and b
}

// 6. ANONYMOUS FUNCTIONS
// Function that accepts a slice and a processor function
func processNumbers(numbers []int, processor func(int) int) []int {
	result := make([]int, len(numbers)) // Create a new slice with same length
	for i, num := range numbers {       // Iterate through input numbers
		result[i] = processor(num) // Apply processor function to each element
	}
	return result // Return the processed slice
}

// 7. CLOSURES
// Function that returns another function (closure)
func createCounter() func() int {
	count := 0          // Variable captured by the closure
	return func() int { // Return an anonymous function
		count++      // Increment the captured variable
		return count // Return the updated count
	}
}

// Another closure example - rate limiter pattern
func createRateLimiter(limit int, interval time.Duration) func() bool {
	lastCall := time.Now() // Track time of last call
	callCount := 0         // Track number of calls in current interval

	return func() bool { // Return the rate limiting function
		now := time.Now() // Get current time

		if now.Sub(lastCall) > interval { // Check if interval has passed
			callCount = 0  // Reset counter
			lastCall = now // Update last call time
		}

		if callCount >= limit { // Check if limit exceeded
			return false // Deny request
		}

		callCount++ // Increment call count
		return true // Allow request
	}
}

// 8. ERROR HANDLING
// Function that returns a value and an error
func divide(a, b float64) (float64, error) {
	if b == 0 { // Check for division by zero
		return 0, errors.New("division by zero") // Return error
	}
	return a / b, nil // Return result and nil error
}

// 9. METHOD - FUNCTION WITH RECEIVER
// Define a Rectangle struct type
type Rectangle struct {
	Width  float64 // Field for width
	Height float64 // Field for height
}

// Method with value receiver (operates on a copy)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height // Calculate and return area
}

// Method with pointer receiver (can modify the original struct)
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor  // Modify width
	r.Height *= factor // Modify height
}

// 10. INTERFACE IMPLEMENTATION
// Define Shape interface with two methods
type Shape interface {
	Area() float64      // Method to calculate area
	Perimeter() float64 // Method to calculate perimeter
}

// Define Circle struct type
type Circle struct {
	Radius float64 // Field for radius
}

// Implement Area method for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius // Calculate circle area
}

// Implement Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius // Calculate circle circumference
}

// Function that accepts any type implementing Shape interface
func printShapeInfo(s Shape) {
	// Print area and perimeter with formatting
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// 11. RECURSIVE FUNCTION
// Function that calls itself to calculate factorial
func factorial(n int) int {
	if n <= 1 { // Base case: factorial of 0 or 1 is 1
		return 1
	}
	return n * factorial(n-1) // Recursive case: n * factorial(n-1)
}

// 12. DEFER, PANIC, AND RECOVER
// Function demonstrating error handling with defer/recover
func readFile(filename string) (content string, err error) {
	fmt.Printf("Opening file: %s\n", filename) // Simulate file opening

	// Defer an anonymous function for cleanup (always executes)
	defer func() {
		fmt.Printf("Closing file: %s\n", filename) // Always close file

		// Recover from any panic that occurred
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r) // Convert panic to error
		}
	}()

	if filename == "invalid.txt" { // Simulate an error condition
		panic("invalid file format") // Cause a panic (will be caught by recover)
	}

	return "File content goes here", nil // Return success result
}

// 13. GENERIC FUNCTION (Go 1.18+)
// Generic function that works with any comparable type T
func findIndex[T comparable](slice []T, value T) int {
	for i, v := range slice { // Iterate through slice
		if v == value { // Compare values (works because T is comparable)
			return i // Return index if found
		}
	}
	return -1 // Return -1 if not found
}

// 14. FUNCTION RETURNING FUNCTION
// Function that returns a multiplier function
func createMultiplier(factor int) func(int) int {
	return func(x int) int { // Return an anonymous function
		return x * factor // Multiply input by the captured factor
	}
}

// 15. CALLBACK PATTERN
// Function that accepts a callback function
func processUserData(userID int, callback func(string)) {
	userData := fmt.Sprintf("Data for user %d", userID) // Simulate data processing

	callback(userData) // Execute the callback with the result
}

// 16. MIDDLEWARE PATTERN
// Define a type for HTTP handler functions
type HttpHandler func(string) string

// Middleware function that wraps another handler
func loggingMiddleware(handler HttpHandler) HttpHandler {
	return func(request string) string { // Return a new handler function
		fmt.Printf("Request received: %s\n", request) // Log the request

		response := handler(request) // Call the original handler

		fmt.Printf("Response sent: %s\n", response) // Log the response

		return response // Return the response
	}
}

// 17. CONCURRENT FUNCTION WITH CHANNELS
// Function that processes data concurrently using goroutines
func processConcurrently(data []string, processor func(string) string) []string {
	results := make([]string, len(data)) // Create result slice

	// Create a channel for receiving results with index information
	resultChan := make(chan struct {
		index int
		value string
	}, len(data)) // Buffered channel with capacity for all results

	// Launch a goroutine for each data item
	for i, item := range data {
		go func(idx int, item string) { // Start goroutine with current values
			// Send processed result through channel
			resultChan <- struct {
				index int
				value string
			}{idx, processor(item)} // Process item and send with its index
		}(i, item) // Pass current values to goroutine
	}

	// Collect all results from the channel
	for range data {
		result := <-resultChan               // Receive result from channel
		results[result.index] = result.value // Store result at correct index
	}

	close(resultChan) // Close the channel
	return results    // Return the processed results
}

// 18. FUNCTION WITH CONTEXT (for cancellation)
// Function that supports cancellation via a done channel
func longRunningOperation(done <-chan bool, operation func()) {
	select { // Select statement for non-blocking channel operations
	case <-done: // Check if cancellation signal received
		fmt.Println("Operation cancelled")
		return // Exit if cancelled
	default: // Default case if no cancellation
		operation() // Execute the operation
	}
}

// 19. CONFIGURABLE FUNCTION WITH OPTIONS PATTERN
// Define ServerConfig struct for configuration
type ServerConfig struct {
	Host    string        // Server hostname
	Port    int           // Server port
	Timeout time.Duration // Request timeout
	Logging bool          // Enable logging
}

// Define function type for configuration options
type ServerOption func(*ServerConfig)

// Option function to set host
func WithHost(host string) ServerOption {
	return func(c *ServerConfig) { // Return a function that modifies config
		c.Host = host // Set the host field
	}
}

// Option function to set port
func WithPort(port int) ServerOption {
	return func(c *ServerConfig) { // Return a function that modifies config
		c.Port = port // Set the port field
	}
}

// Function that creates config with optional parameters
func createServerConfig(options ...ServerOption) *ServerConfig {
	config := &ServerConfig{ // Create config with default values
		Host:    "localhost",      // Default host
		Port:    8080,             // Default port
		Timeout: 30 * time.Second, // Default timeout
		Logging: true,             // Default logging enabled
	}

	for _, option := range options { // Apply all provided options
		option(config) // Call each option function on the config
	}

	return config // Return the configured struct
}

// 20. VALIDATION FUNCTION
// Function that validates user input and returns errors
func validateUserInput(input map[string]string) (bool, map[string]string) {
	errors := make(map[string]string) // Create map to store validation errors

	// Validate username field
	if username, exists := input["username"]; exists { // Check if username exists
		if len(username) < 3 { // Validate minimum length
			errors["username"] = "must be at least 3 characters"
		}
	} else { // Username is missing
		errors["username"] = "is required"
	}

	// Validate email field
	if email, exists := input["email"]; exists { // Check if email exists
		if !strings.Contains(email, "@") { // Validate email format
			errors["email"] = "must be a valid email address"
		}
	} else { // Email is missing
		errors["email"] = "is required"
	}

	return len(errors) == 0, errors // Return validation status and errors
}

// MAIN FUNCTION - Entry point of the program
func main() {
	fmt.Println("=== FUNCTION EXAMPLES ===") // Print section header

	// 1. Basic function example
	area := calculateArea(5.0, 3.0)               // Call function with width and height
	fmt.Printf("1. Rectangle area: %.2f\n", area) // Print result

	// 2. Multiple return values example
	circleArea, circumference := calculateCircle(3.0) // Call function
	// Print both return values
	fmt.Printf("2. Circle area: %.2f, circumference: %.2f\n", circleArea, circumference)

	// 3. Named return values example
	numbers := []int{5, 2, 8, 1, 9, 3}            // Create slice of numbers
	min, max, avg := calculateStatistics(numbers) // Call function
	// Print all three return values
	fmt.Printf("3. Stats - Min: %d, Max: %d, Avg: %d\n", min, max, avg)

	// 4. Variadic function example
	sum := calculateSum(1, 2, 3, 4, 5)         // Call with variable arguments
	fmt.Printf("4. Sum of numbers: %d\n", sum) // Print result

	// 5. Function as parameter example
	result := applyOperation(5, 3, multiply) // Pass multiply function
	fmt.Printf("5. 5 * 3 = %d\n", result)    // Print result

	// 6. Anonymous functions example
	squared := processNumbers(numbers, func(x int) int { // Pass anonymous function
		return x * x // Square the input
	})
	fmt.Printf("6. Squared numbers: %v\n", squared) // Print result

	// 7. Closures example
	counter := createCounter() // Create a counter closure
	fmt.Printf("7. Counter: ") // Print label
	for i := 0; i < 3; i++ {   // Call counter three times
		fmt.Printf("%d ", counter()) // Print current count
	}
	fmt.Println() // New line

	// 8. Error handling example
	divisionResult, err := divide(10, 2) // Call function that might return error
	if err != nil {                      // Check if error occurred
		fmt.Printf("8. Error: %v\n", err) // Print error
	} else {
		fmt.Printf("8. 10 / 2 = %.2f\n", divisionResult) // Print result
	}

	// 9. Methods example
	rect := Rectangle{Width: 3, Height: 4}                      // Create Rectangle instance
	fmt.Printf("9. Rectangle area: %.2f\n", rect.Area())        // Call method
	rect.Scale(2)                                               // Call method with pointer receiver
	fmt.Printf("   Scaled rectangle area: %.2f\n", rect.Area()) // Call method again

	// 10. Interface example
	circle := Circle{Radius: 3} // Create Circle instance
	printShapeInfo(circle)      // Call function with Circle (implements Shape)

	// 11. Recursive function example
	fact := factorial(5)              // Calculate factorial of 5
	fmt.Printf("11. 5! = %d\n", fact) // Print result

	// 12. Defer, panic, recover example
	content, err := readFile("invalid.txt") // Call function that might panic
	if err != nil {                         // Check if error occurred
		fmt.Printf("12. Error reading file: %v\n", err) // Print error
	} else {
		fmt.Printf("12. File content: %s\n", content) // Print content
	}

	// 13. Generic function example
	stringSlice := []string{"apple", "banana", "cherry"} // Create string slice
	index := findIndex(stringSlice, "banana")            // Call generic function
	fmt.Printf("13. Index of 'banana': %d\n", index)     // Print result

	// 14. Function returning function example
	double := createMultiplier(2)                  // Get a function that doubles numbers
	fmt.Printf("14. Double of 5: %d\n", double(5)) // Use the returned function

	// 15. Callback pattern example
	processUserData(123, func(data string) { // Call function with callback
		fmt.Printf("15. Processed: %s\n", data) // Callback implementation
	})

	// 16. Middleware pattern example
	handler := func(request string) string { // Define simple handler
		return "Response to: " + request // Return response
	}
	wrappedHandler := loggingMiddleware(handler)     // Wrap handler with middleware
	response := wrappedHandler("test request")       // Call wrapped handler
	fmt.Printf("16. Final response: %s\n", response) // Print response

	// 17. Concurrent processing example
	data := []string{"item1", "item2", "item3"}             // Create data slice
	processed := processConcurrently(data, strings.ToUpper) // Process concurrently
	fmt.Printf("17. Concurrent results: %v\n", processed)   // Print results

	// 18. Context cancellation example
	done := make(chan bool) // Create cancellation channel
	go func() {             // Start goroutine
		time.Sleep(100 * time.Millisecond) // Wait a bit
		close(done)                        // Close channel to signal cancellation
	}()
	longRunningOperation(done, func() { // Call function with cancellation
		fmt.Println("18. This would be a long operation") // Operation to perform
	})

	// 19. Options pattern example
	config := createServerConfig(WithHost("example.com"), WithPort(9000)) // Create config with options
	fmt.Printf("19. Server config: %+v\n", config)                        // Print config

	// 20. Validation example
	userInput := map[string]string{ // Create test input
		"username": "ab",            // Invalid: too short
		"email":    "invalid-email", // Invalid: missing @
	}
	isValid, validationErrors := validateUserInput(userInput) // Validate input
	// Print validation results
	fmt.Printf("20. Validation passed: %t, Errors: %v\n", isValid, validationErrors)

	fmt.Println("\n=== ALL EXAMPLES COMPLETED ===") // Print completion message
}
