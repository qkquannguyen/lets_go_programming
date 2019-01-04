// --- Specifies this file belongs in the "main" package
package main

// --- Importing an exisiting package
import "fmt"

func main() {
	// --- Using Println() function in the "fmt" package
	fmt.Println("Lets Go Packages!")
}

// --- IMPORTANT NOTE: Your custom packages go to your Go Workspace. Otherwise, you will receive a
// --- a error regarding the package not existing in the Go Path.

// --- OTHER NOTES:
// --- Used to organize source code for better reusability and readability.
// --- Every Go file should have some "package packagename" in the first line.
// --- In this example, compile by typing "go install goPackageExample".
// --- This will generate a binary inside the bin folder of the workspace.

