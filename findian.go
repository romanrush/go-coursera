/* Write a program which prompts the user to enter a string.
The program searches through the entered string for the
characters ‘i’, ‘a’, and ‘n’. The program should print “Found!”
 if the entered string starts with the character ‘i’, ends
 with the character ‘n’, and contains the character ‘a’.
 The program should print “Not Found!” otherwise. The program
  should not be case-sensitive, so it does not matter if the
  characters are upper-case or lower-case.

Examples: The program should print “Found!” for the
following example entered strings, “ian”, “Ian”, “iuiygaygn”,
 “I d skd a efju N”. The program should print “Not Found!”
 for the following strings, “ihhhhhn”, “ina”, “xian”.

Submit your source code for the program,
“findian.go”. */

/* Write a program which prompts the user to enter a
floating point number and prints the integer which is a
truncated version of the floating point number that was
entered. Truncation is the process of removing the digits
to the right of the decimal place.
Submit your source code for the program, “trunc.go”. */

package main

import (
	"fmt"

	"strings"
)

func main() {

	fmt.Println("Enter the string to check: ")
	// var then variable name then variable type
	var inputString string
	// Taking input from user
	fmt.Scanln(&inputString)
	lowerCaseInput := strings.ToLower(inputString)
	if strings.HasPrefix(lowerCaseInput, "i") && strings.Contains(lowerCaseInput, "a") && strings.HasSuffix(lowerCaseInput, "n") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not found!")
	}
}
