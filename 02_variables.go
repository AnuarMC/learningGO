/*
	Example: declaring variables
*/
package main

import "fmt"

func main() {

	// first definition of variables
	var name, lastname string
	name = "Michel"
	lastname = "Fech"

	// second definition of variables
	address := "456 White Finch St."

	//Third definition of variables
	user1, user2 := "emely", "alexis"

	fmt.Println("Data: ", name, lastname)
	fmt.Println("Address: ", address)
	fmt.Println("First user: ", user1)
	fmt.Println("Second user: ", user2)

	//constant, variables that cannot be modified
	const name2 = "Dickson"
	fmt.Println(name2)

	// zero or null values variable
	var name3 string
	var number int
	var value bool
	fmt.Println("name3: ", name3)
	fmt.Println("number: ", number)
	fmt.Println("value: ", value)
}
