/* Write a program which prompts the user to first enter
a name, and then enter an address.
Your program should create a map and add the name and
address to the map using the keys “name” and “address”,
respectively. Your program should use Marshal() to
create a JSON object from the map, and then your program should
print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var userName, userAddress string
	fmt.Println("Enter your name:")
	fmt.Scanln(&userName)
	fmt.Println("> your name is:", userName)
	fmt.Println("\n Enter your address:")
	fmt.Scanln(&userAddress)
	fmt.Println("> your address is:", userAddress)

	// create a map
	m := make(map[string]string)
	m["name"] = userName
	m["address"] = userAddress
	fmt.Println("\n created map:", m)

	// create a json object
	a, _ := json.Marshal(m)

	//print the json object
	fmt.Println("\n JSON object:")
	fmt.Println(string(a))

}
