/* Write a program which prompts the user to enter a
floating point number and prints the integer which is a
truncated version of the floating point number that was
entered. Truncation is the process of removing the digits
to the right of the decimal place.
Submit your source code for the program, “trunc.go”. */

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter A float of your choice: ")

	// var then variable name then variable type
	var inputFloat float32

	// Taking input from user
	fmt.Scanln(&inputFloat)

	//truncated the float:
	fmt.Printf("Your truncated float is: %.0f\n", inputFloat)

}
