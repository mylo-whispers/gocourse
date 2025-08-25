package basis

import "fmt"

type Employee struct {
	first_name string
	last_name  string
	age        string
}

func main() {
	// Pascal case
	// eg. calculatearea , userinfo , newHTTPRequest
	// Structs , interfaces , enums

	// snake_case
	// eg. user_id, first_name, http_request

	// UPPERCASE
	// Use case is constants

	// mixedcase
	// eg. javaSxript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employee_id = 1001
	fmt.Println("Employee: ", employee_id)
}
