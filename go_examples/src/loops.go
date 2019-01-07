package main

import (
	"fmt"
)

func main() {
	// --- Example of for-Loop with break statement
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf(" %d", i)
	}
	fmt.Printf("\nline after for loop")

	// --- Example of for-Loop with continue statement
	// --- This should skip the even numbers and print only odd numbers
	fmt.Printf("\n")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	// --- Example of a for-Loop when initialization and post are omitted
	// --- Semicolons can also be omitted
	fmt.Printf("\n")
	j := 0
	for j <= 10 {
		fmt.Printf("%d ", j)
		j += 2
	}

	// --- Example of multiple initialization and increment
	for no, k := 10, 1; k <= 10 && no <= 19; k, no = k+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, k, no*k)
	}
}

// --- NOTES:
// --- In Go, the only loop available is the "for" loop
// --- for loop syntax:
// --- for initialization; condition; post {}
// --- initialization --> Executed only once
// --- After initialization, the condition is checked. If true, body in the {} runs
// --- and then post statement. After the post statement is executed, the condition is rechecked.
// --- If true, the loop continues, otherwise the loop terminates
