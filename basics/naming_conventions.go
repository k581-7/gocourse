package basics

import "fmt"

type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

type EmployeeRTX struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// Ex. CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case
	// Ex. user_id, first_name, http_request

	//UPPERCASE

	// mixedCase
	// Ex. javaScript, htmlDocument, isValid

	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)

}
