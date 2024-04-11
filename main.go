package main

import (
	"TMCheckout/checkout"
	"fmt"
)

func main() {
	var testCart checkout.Totaliser = &checkout.SimpleCheckout{}
	TestCart(testCart)
	testCart.Scan("A")
	testCart.Scan("B")
	i, _ := testCart.GetTotalPrice()
	fmt.Println("Total", i)
}

func TestCart(a checkout.Totaliser) {
	a.Clear()
	a.Scan("A")
	a.Scan("A")
	a.Scan("A")
	i, _ := a.GetTotalPrice()
	fmt.Println("Total", i)
	a.Clear()
}
