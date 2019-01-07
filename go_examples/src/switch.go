package main

import (
	"fmt"
)

func main() {
	// --- Basic Switch Example
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("Incorrect finger number")
	}

	// --- Multiple Expression in Case
	letter := "i"
	switch letter {
	// --- Below is an example of a multiple expression in Switch Cases
	case "a", "e", "i", "o", "u":
		fmt.Println("Vowel")
	default:
		fmt.Println("Not a vowel")
	}
}
