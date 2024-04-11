package checkout

import (
	"TMCheckout/pricing"
	"testing"
)

func TestAddItem(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{Pricing: pricing.NominalPricer{}}
	testCart.Scan("A")
	i, _ := testCart.GetTotalPrice()
	if i != 1 {
		t.Errorf("Incorrect value in cart -should be 1")
	}
}

func TestAddItems(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{Pricing: pricing.NominalPricer{}}
	testCart.Scan("A")
	testCart.Scan("B")
	i, _ := testCart.GetTotalPrice()
	if i != 2 {
		t.Errorf("Incorrect value in cart -should be 2")
	}
}

func TestAddItemWithClear(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{Pricing: pricing.NominalPricer{}}
	testCart.Scan("A")
	testCart.Clear()
	testCart.Scan("B")
	i, _ := testCart.GetTotalPrice()
	if i != 1 {
		t.Errorf("Incorrect value in cart -should be 1")
	}
}

func TestNoPricing(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{}
	testCart.Scan("A")
	i, err := testCart.GetTotalPrice()
	if err == nil {
		t.Errorf("No pricer created but not reported")
	}
	if i != 0 {
		t.Errorf("Incorrect value in cart -should be 0")
	}
}

func TestSimplePricing(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{Pricing: pricing.SimplePricer{}}
	err := testCart.Scan("A")
	if err != nil {
		t.Errorf("Error returned on scan %s", err.Error())
	}
	i, err := testCart.GetTotalPrice()
	if err != nil {
		t.Errorf("Error returned on GetTotalPrice %s", err.Error())
	}
	if i != 50 {
		t.Errorf("Incorrect value in cart -should be 50, actual %d", i)
	}
	err = testCart.Scan("B")
	if err != nil {
		t.Errorf("Error returned on scan %s", err.Error())
	}
	i, err = testCart.GetTotalPrice()
	if err != nil {
		t.Errorf("Error returned %s", err.Error())
	}
	if i != 80 {
		t.Errorf("Incorrect value in cart -should be 80, actual %d", i)
	}
	err = testCart.Scan("Q")
	if err != nil {
		t.Errorf("Error returned on scan %s", err.Error())
	}
	i, err = testCart.GetTotalPrice()
	if err == nil {
		t.Errorf("Error NOT returned with invalid item Q in basket")
	}
	if i != 0 {
		t.Errorf("Incorrect value in cart -should be 0, actual %d", i)
	}
}
