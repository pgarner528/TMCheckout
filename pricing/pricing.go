package pricing

import (
	"fmt"
	"strings"
)

type Pricer interface {
	ItemPrice(item string, numberOff int) (int, error)
}

func NewMultibuyPricer() MultibuyPricer {
	newPricer := MultibuyPricer{}
	newPricer.priceMap = map[string]itemPrice{}
	newPricer.loadPriceFile()
	return newPricer
}

type itemPrice struct {
	unitPrice      int
	multiBreakQty  int
	multipackPrice int
}
type MultibuyPricer struct {
	priceMap map[string]itemPrice
}

func (N MultibuyPricer) ItemPrice(item string, numberOff int) (int, error) {
	anItemPrice, ok := N.priceMap[strings.ToUpper(item)]
	if !ok {
		return 0, fmt.Errorf("no price for item SKU %s", item)
	} else {
		units, multipacks := multiBreak(numberOff, anItemPrice.multiBreakQty)
		return (anItemPrice.unitPrice * units) + (anItemPrice.multipackPrice * multipacks), nil
	}
}

func (N *MultibuyPricer) loadPriceFile() error {
	// Ideally this would be returned from a DB or web service but this is outside the scope of this demo
	N.priceMap["A"] = itemPrice{unitPrice: 50, multiBreakQty: 3, multipackPrice: 130}
	N.priceMap["B"] = itemPrice{unitPrice: 30, multiBreakQty: 2, multipackPrice: 45}
	N.priceMap["C"] = itemPrice{unitPrice: 20}
	N.priceMap["D"] = itemPrice{unitPrice: 15}

	return nil
}

func multiBreak(numberOff, multipackBreak int) (int, int) {
	if multipackBreak == 0 {
		return numberOff, 0
	}
	return numberOff % multipackBreak, numberOff / multipackBreak
}

type SimplePricer struct {
}

func (N SimplePricer) ItemPrice(item string, numberOff int) (int, error) {
	switch {
	case strings.EqualFold("A", item):
		return 50 * numberOff, nil
	case strings.EqualFold("B", item):
		return 30 * numberOff, nil
	case strings.EqualFold("C", item):
		return 20 * numberOff, nil
	case strings.EqualFold("D", item):
		return 15 * numberOff, nil
	default:
		return 0, fmt.Errorf("no price for item SKU %s", item)
	}
}

type NominalPricer struct {
}

func (N NominalPricer) ItemPrice(item string, numberOff int) (int, error) {
	return 1 * numberOff, nil
}
