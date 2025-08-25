package basic

import "fmt"

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b

	fmt.Println("Addition: ", result)

	result = a - b

	fmt.Println("Addition: ", result)

	result = a / b

	fmt.Println("division: ", result)

	result = a % b

	fmt.Println("reminder: ", result)

	result = a * b

	fmt.Println("multiplication: ", result)

	const p float64 = 22 / 7.0

	fmt.Println("Float change: ", p)

}
