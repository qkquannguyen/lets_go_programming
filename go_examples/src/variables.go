package main

import "fmt"

func main() {
	// --- Basic Variable Declaration
	var age int
	fmt.Println("The member age is:", age)

	// --- Method to Declare Variables
	var width, height int = 100, 50
	fmt.Println("The width is:", width, "\nThe height is:", height)

	// --- Multiple Variable Declaration
	var (
		name     = "B3LL"
		guild    = "Revolution"
		position = "member"
	)
	fmt.Println("User Id:", name, "\nGuild:", guild, "\nPosition:", position)

	// --- Short Hand Declaration
	member, status := "Monji", "retired"
	fmt.Println("Member Name:", member, "\nStatus:", status)

	// --- NOTE:
	// --- Go-Lang is capable of type inference
}
