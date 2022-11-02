/* Write a program which reads information from a file and represents it in
a slice of structs. Assume that there is a text file which contains a series
 of names. Each line of the text file has a first name and a last name, in that
  order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first
 name, and lname for the last name. Each field will be a string of size 20
 (characters).

Your program should prompt the user for the name of the text file. Your program
 will successively read each line of the text file and create a struct which
 contains the first and last names found in the file. Each struct created will
 be added to a slice, and after all lines have been read from the file, your
 program will have a slice containing one struct for each line in the file. After
 reading all lines from the file, your program should iterate through your slice
 of structs and print the first and last names found in each struct. */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Name struct {
	fname string
	lname string
}

/* I tested the program with a file named nameFile.txt
 containing the following lines:
Nicholas Reily
Zoe Micaela
Marie Breonia
June Miranda
Ann Madelyn
*/
func main() {
	names := []Name{}

	//prompt user for the name of the text file:
	var fileName string
	fmt.Println("Enter the file name:")
	fmt.Scanln(&fileName)
	fmt.Println("> the file name is:", fileName)
	path, err := os.Getwd()
	// read file
	file, err := os.Open(filepath.Join(path, fileName)) //fileName OR nameFile.txt
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// read file line by line:
	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		// create struct with each line:
		//fmt.Printf("currently reading: %s \n", line)
		splittedLine := strings.Split(string(line), " ")
		firstName := splittedLine[0]
		lastName := splittedLine[1]
		var newLineStruct = Name{firstName, lastName}

		// add created struct to slice of structs:
		names = append(names, newLineStruct)
		//fmt.Println("names: \n", names)
	}

	//Once file read, iterate through the slice and print the first & last name of each struct:
	for _, name := range names {
		fmt.Printf("Firstname: %s || Lastname: %s\n", name.fname, name.lname)
		fmt.Println("")
	}
}
