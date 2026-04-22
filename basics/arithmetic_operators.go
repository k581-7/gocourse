package basics

import (
	"fmt"
	"math"
)

func main() {
	// variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Addition:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Remainder:", result)

	const p float64 = 10 / 3.0
	fmt.Println(p)

	// Overflow with signed integers - both positive and negative numbers
	// signed integers result to negative number because it already exceeded what it can hold
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt = maxInt + 1
	fmt.Println(maxInt)

	// Overflow with unsigned integers - only zero to positive numbers
	// unsigned integers result to zero because it cannot produce negative numbers
	var uMaxInt uint64 = 18446744073709551615
	fmt.Println(uMaxInt)

	uMaxInt = uMaxInt + 1
	fmt.Println(uMaxInt)

	// Underflow with floating point numbers
	var smallFloat float64 = 1.0e-323 // This is a compact way of writing a decimal point followed by 322 zeros and a 1 like 0.000...01

	fmt.Println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	fmt.Println(smallFloat)

}
