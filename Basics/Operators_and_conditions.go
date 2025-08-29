package Basics

import "fmt"

func main() {
	// OPERATOR PRECEDENCE
	// Understanding which operations are performed first

	a, b, c := 5, 10, 15

	// Precedence: * and / have higher precedence than + and -
	result1 := a + b*c // Equivalent to a + (b * c)
	result2 := (a + b) * c
	fmt.Printf("%d + %d * %d = %d\n", a, b, c, result1)
	fmt.Printf("(%d + %d) * %d = %d\n", a, b, c, result2)

	// Precedence: Comparison operators have higher precedence than logical operators
	x, y, z := 10, 20, 30
	condition1 := x < y && y < z // Equivalent to (x < y) && (y < z)
	condition2 := x < y || y > z // Equivalent to (x < y) || (y > z)
	fmt.Printf("%d < %d && %d < %d = %t\n", x, y, y, z, condition1)
	fmt.Printf("%d < %d || %d > %d = %t\n", x, y, y, z, condition2)

	// Using parentheses to control evaluation order
	// complexResult1 := a + b*c/d - e%f       // Hard to read
	// complexResult2 := a + (b*c)/d - (e % f) // Clear precedence

	// Complex condition with proper parentheses
	age, income, hasJob := 25, 50000, true
	eligibleForLoan := (age >= 21) && (income > 30000 || hasJob) && (income < 100000 || age < 30)
	fmt.Printf("Eligible for loan: %t\n", eligibleForLoan)

	// De Morgan's Laws in practice
	// Not (A and B) = (Not A) or (Not B)
	// Not (A or B) = (Not A) and (Not B)

	isWeekend, isHoliday := false, true
	// Original condition
	shouldWork := !(isWeekend || isHoliday)
	// Applying De Morgan's Law
	shouldWork2 := !isWeekend && !isHoliday

	fmt.Printf("Should work (original): %t\n", shouldWork)
	fmt.Printf("Should work (De Morgan): %t\n", shouldWork2)
}

// Placeholder variables for demonstration
var d, e, f int = 2, 7, 3
