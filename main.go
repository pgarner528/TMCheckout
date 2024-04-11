package main

import (
	"TMCheckout/checkout"
	"fmt"
)

func main() {
	var testCart checkout.Totaliser = &checkout.SimpleCheckout{}
	testCart.Scan("A")
	testCart.Scan("B")
	i, _ := testCart.GetTotalPrice()
	fmt.Println("Total", i)
}
