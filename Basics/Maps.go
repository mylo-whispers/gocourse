package Basics

import "fmt"

type Product struct {
	ID       string
	Name     string
	Price    float64
	Category string
}

func main() {
	// SHOPPING CART SYSTEM
	// Using maps to manage a shopping cart

	// Available products
	products := map[string]Product{
		"P001": {ID: "P001", Name: "Laptop", Price: 999.99, Category: "Electronics"},
		"P002": {ID: "P002", Name: "Mouse", Price: 25.50, Category: "Electronics"},
		"P003": {ID: "P003", Name: "Book", Price: 15.99, Category: "Education"},
		"P004": {ID: "P004", Name: "Coffee Mug", Price: 12.75, Category: "Kitchen"},
	}

	// Shopping cart: map of product ID to quantity
	cart := make(map[string]int)

	// Add items to cart
	cart["P001"] = 1 // Add laptop
	cart["P002"] = 2 // Add 2 mice
	cart["P004"] = 3 // Add 3 coffee mugs

	// Display cart contents
	fmt.Println("Shopping Cart:")
	total := 0.0
	for productID, quantity := range cart {
		if product, exists := products[productID]; exists {
			itemTotal := product.Price * float64(quantity)
			total += itemTotal
			fmt.Printf("%s: %s x%d = $%.2f\n", productID, product.Name, quantity, itemTotal)
		}
	}
	fmt.Printf("Total: $%.2f\n", total)

	// Update cart quantity
	cart["P002"] = 1 // Change mouse quantity from 2 to 1
	fmt.Printf("\nUpdated mouse quantity to: %d\n", cart["P002"])

	// Remove item from cart
	delete(cart, "P004") // Remove coffee mugs
	fmt.Println("Removed coffee mugs from cart")

	// Calculate new total
	total = 0.0
	for productID, quantity := range cart {
		if product, exists := products[productID]; exists {
			total += product.Price * float64(quantity)
		}
	}
	fmt.Printf("New total: $%.2f\n", total)

	// Check if cart is empty
	if len(cart) == 0 {
		fmt.Println("Your cart is empty!")
	} else {
		fmt.Printf("You have %d items in your cart\n", len(cart))
	}
}
