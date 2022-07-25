/*
	Example: using arrays
*/
package main

import "fmt"

// En los array se necesita especificar la capacidad
func main() {
	var names [3]string
	names[0] = "Emely"
	names[1] = "Tom"
	names[2] = "Alexis"

	fmt.Println(names)
	lastnames := [3]string{"Johnson", "Taylor", "Smith"}
	fmt.Println(lastnames)
	fmt.Println("name1: ", names[1], "lastname1: ", lastnames[1])

	size := len(names)
	fmt.Println("Size array: ", size)

	fmt.Println(names[0:2])
}
