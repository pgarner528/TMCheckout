package checkout

import (
	"strings"
)

type Totaliser interface {
	Scan(item string) error
	GetTotalPrice() (int, error)
	Clear()
}

type SimpleCheckout struct {
	cart map[string]int
}

func (c *SimpleCheckout) Scan(item string) error {
	if c.cart == nil {
		c.cart = map[string]int{}
	}
	saneItem := strings.ToUpper(item)
	curVal, ok := c.cart[saneItem]
	if ok {
		c.cart[saneItem] = curVal + 1
	} else {
		c.cart[saneItem] = 1
	}
	return nil
}

func (c *SimpleCheckout) GetTotalPrice() (int, error) {
	// Everything has a nominal value of 1 for the moment
	if c.cart == nil {
		return 0, nil
	}
	iTotal := 0
	for _, val := range c.cart {
		iTotal += val
	}
	return iTotal, nil
}

func (c *SimpleCheckout) Clear() {
	if c.cart != nil {
		clear(c.cart)
	}
}
