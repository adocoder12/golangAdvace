package main

import (
	"fmt"

	"github.com/adocoder12/golangAdvace/model"
)

// purchasable is an INTERFACE: a list of methods a type must have.
// Any type with a CalculatePrice() int64 method automatically counts
// as "purchasable". No "implements" keyword needed in Go.
type purchasable interface {
	CalculatePrice() int64
}

// cart is a slice of the interface type, not of Shirt or Monitor.
// So it can hold ANY mix of types, as long as each one has CalculatePrice().
var cart []purchasable

// addToCart takes any number of purchasable items (the "..." means variadic).
// Inside the function, "products" is a []purchasable.
func addToCart(products ...purchasable) {
	// "products..." unpacks the slice into separate values for append.
	cart = append(cart, products...)
}

// getTotalCart adds up the price of everything in the cart.
func getTotalCart() int64 {
	var total int64

	for _, product := range cart {
		// THIS IS THE POLYMORPHISM:
		// the same call, product.CalculatePrice(), runs different code
		// depending on the real type inside the interface.
		//   Monitor -> price + 30% electronic tax
		//   Shirt   -> price - 20% discount
		//   Wine    -> price + 23% liquor tax
		// This loop doesn't know or care which type it is.
		total += product.CalculatePrice()
	}

	return total
}

func main() {

	myShirt := model.Shirt{
		Price: 50,
		Brand: "nike",
		Size:  "M",
		Color: "Blue",
	}

	myMonitor := model.Monitor{
		Price:     600,
		Brand:     "Sony",
		Size:      "22 inch",
		Resolutin: "1800 x 1600", // typo kept: it must match the field name in model/monitor.go
	}

	// Two different types go into the same cart. This works because both
	// have CalculatePrice(), so both satisfy the purchasable interface.
	addToCart(myMonitor, myShirt)

	// Monitor: 600 + 30% = 780
	// Shirt:    50 - 20% =  40
	// Total: 820 (if your model files use the fixed versions)
	fmt.Println("the total price is: ", getTotalCart())
}
