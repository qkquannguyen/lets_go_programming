package main

import (
	"fmt"
)

// --- NOTE: One way to write functions
func calculateBill(price int, tax int) int {
	var totalPrice = price * tax
	return totalPrice
}

// --- NOTE: If parameters are of same type, you can do this shorthand
func calculateMyBills(price, tax int) int {
	var totalPrice = price * tax
	return totalPrice
}

// --- NOTE: You can have multiple return values in functions
func rectangleProperties(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// --- NOTE: You can do named return value similar to below
func squareProperties(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {
	// --- Testing Functions 1 & 2
	price, tax := 100, 8
	totalPrice := calculateBill(price, tax)
	totalPrice2 := calculateMyBills(price, tax)
	fmt.Println("Total in Function 1:", totalPrice, "\nTotal in Function 2:", totalPrice2)

	// --- Testing Function 3
	area, perimeter := rectangleProperties(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	// --- Testing Function 4
	area2, perimeter2 := squareProperties(11.1, 11.1)
	fmt.Printf("\nArea %f Perimeter %f", area2, perimeter2)
}
