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
