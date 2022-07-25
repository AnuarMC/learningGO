/*
	Example: try conditionals
*/
package main

import "fmt"

/*
	or: ||
	and: &&
	not: !
*/

func main() {
	value := true
	number1, number2 := 5, 8

	if value {
		fmt.Println("Value is true ")
	} else {
		fmt.Println("Value is false ")
	}

	if number1 > 5 {
		fmt.Println("number1 is greate 5")
	} else if number1 < 5 {
		fmt.Println("number1 is less 5")
	} else if number1 == 5 {
		fmt.Println("number1 is equal 5")
	}
	fmt.Println(number2)
}
