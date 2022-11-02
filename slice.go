/* Write a program which prompts the user to enter integers and stores the integers
in a sorted slice. The program should be written as a loop. Before entering the loop,
the program should create an empty integer slice of size (length) 3. During each pass through the
loop, the program prompts the user to enter an integer to be added to the slice. The program adds
the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter.
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of
an integer. */
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	intSlice := make([]int, 0, 3) // initializing integer slice with length and capcity of 3

	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Please type in an integer or press 'X' to quit:")
	for input.Scan() { //triggers a new loop cycle whenever there’s new input received from the program’s standard input stream.
		userInput := input.Text()
		fmt.Println("> you typed in:", userInput)
		userInput = strings.ToLower(userInput)
		//check if input is x: in which case you have to break the loop
		if strings.Compare(userInput, "x") == 0 {
			fmt.Println("the final slice is:", intSlice)
			fmt.Println("the program will quit")
			break
		} else if intVar, err := strconv.Atoi(userInput); err == nil { // check if the number is an integer, in which case push to the slice:
			// push to slice:
			intSlice = append(intSlice, intVar)
		} else {
			// if the number is not an integer, warn the user:
			fmt.Println("the input is not an integer!")
		}
		// sort the slice before next iteration:
		sort.Ints(intSlice)
		fmt.Println("the current slice is:", intSlice)
		//finally:
		fmt.Println("Please type in an other integer or press 'X' to quit:")
	}
}
