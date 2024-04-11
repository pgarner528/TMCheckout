package checkout

import (
	"TMCheckout/pricing"
	"testing"
)

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

type AddItem struct {
	item          string
	runningTotal  int
	errorExpected bool
}

func TestMultibuyPricing(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{Pricing: pricing.NewMultibuyPricer()}
	testSlice := []AddItem{
		{item: "A", runningTotal: 50},
		{item: "B", runningTotal: 80},
		{item: "B", runningTotal: 95},
		{item: "B", runningTotal: 125},
		{item: "A", runningTotal: 175},
		{item: "A", runningTotal: 205},
		{item: "C", runningTotal: 225},
		{item: "D", runningTotal: 240},
		{item: "Z", runningTotal: 0, errorExpected: true},
	}
	for _, itemPrice := range testSlice {
		err := testCart.Scan(itemPrice.item)
		if err != nil {
			t.Errorf("Error returned on scan %s", err.Error())
		}
		i, err := testCart.GetTotalPrice()
		if (err != nil) && (!itemPrice.errorExpected) {
			t.Errorf("Error returned on GetTotalPrice %s", err.Error())
		} else {
			if (err == nil) && (itemPrice.errorExpected) {
				t.Errorf("Error NOT returned on GetTotalPrice")
			}
		}
		if i != itemPrice.runningTotal {
			t.Errorf("Incorrect value in cart -should be %d, actual %d", itemPrice.runningTotal, i)
		}
	}
}
