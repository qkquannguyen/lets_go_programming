package main

import (
	"fmt"
)

// --- This shows how to write a if-Else statement in Go
func checkNum() {
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

// --- This is another way of writing a if-Else statement in Go
func checkNumAgain() {
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even!")
	} else {
		fmt.Println(num, "is odd")
	}
}

// --- This will show how to write an "else if" in a if-Else
func main() {
	num := 99
	if num <= 50 {
		fmt.Println("number is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println("number is bwetween 51 amd 100")
	} else {
		fmt.Println("number is greater than 100")
	}
}

// --- NOTES
// --- The "else" statement needs to be on the SAME line as the } bracket of the "if" or "else if"
// --- Compiler will complain if the "else" is on another line.
// --- The reason is because of Go-Lang's semicolon insertion rule
