package main

import (
	"TMCheckout/checkout"
	"TMCheckout/pricing"
	"fmt"
)

func main() {
	var testCart checkout.Totaliser = &checkout.SimpleCheckout{
		Pricing: pricing.SimplePricer{},
	}
	testCart.Scan("A")
	testCart.Scan("B")
	testCart.Scan("A")
	testCart.Scan("Q")
	i, err := testCart.GetTotalPrice()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Total", i)
}
