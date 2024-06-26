package main

import (
	"TMCheckout/checkout"
	"TMCheckout/pricing"
	"fmt"
)

func main() {
	var testCart checkout.Totaliser = &checkout.SimpleCheckout{
		Pricing: pricing.NewMultibuyPricer(),
	}
	testCart.Scan("A")
	testCart.Scan("B")
	testCart.Scan("A")
	testCart.Scan("A")
	testCart.Scan("B")
	i, err := testCart.GetTotalPrice()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Total", i)
}
