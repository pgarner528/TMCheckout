package checkout

import (
	"TMCheckout/pricing"
	"errors"
	"strings"
)

type Totaliser interface {
	Scan(item string) error
	GetTotalPrice() (int, error)
	Clear()
}

type SimpleCheckout struct {
	cart    map[string]int
	Pricing pricing.Pricer
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
	if c.Pricing == nil {
		return 0, errors.New("no pricing available")
	}
	if c.cart == nil {
		return 0, nil
	}
	iTotal := 0
	for item, val := range c.cart {
		iItemValue, err := c.Pricing.ItemPrice(item, val)
		if err != nil {
			return 0, err
		}
		iTotal += iItemValue
	}
	return iTotal, nil
}

func (c *SimpleCheckout) Clear() {
	if c.cart != nil {
		clear(c.cart)
	}
}
