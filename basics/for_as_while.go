package basics

import "fmt"

func main() {

	// for as while with break
	// sum := 0
	// for {
	// 	sum += 10
	// 	fmt.Println("Sum:", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	num := 1 // declare ng number
	for num <= 10 { // maximum loop number
		if num%2 == 0 { // condition - if number is even, plus one sya
			num++ // eto yung plus one
			continue // this code means turn back to start if the number is even, it stops here and goes back to start
		}
		fmt.Println("Odd Number:", num) // this prints the odd number from above filter
		num++ // increments after getting odd number
	}
}