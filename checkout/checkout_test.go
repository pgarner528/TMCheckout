package checkout

import "testing"

func TestAddItem(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{}
	testCart.Scan("A")
	i, _ := testCart.GetTotalPrice()
	if i != 1 {
		t.Errorf("Incorrect value in cart -should be 1")
	}
}

func TestAddItems(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{}
	testCart.Scan("A")
	testCart.Scan("B")
	i, _ := testCart.GetTotalPrice()
	if i != 2 {
		t.Errorf("Incorrect value in cart -should be 2")
	}
}

func TestAddItemWithClear(t *testing.T) {
	var testCart Totaliser = &SimpleCheckout{}
	testCart.Scan("A")
	testCart.Clear()
	testCart.Scan("B")
	i, _ := testCart.GetTotalPrice()
	if i != 1 {
		t.Errorf("Incorrect value in cart -should be 1")
	}
}
