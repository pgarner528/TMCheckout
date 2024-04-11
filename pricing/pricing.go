package pricing

import (
	"fmt"
	"strings"
)

type Pricer interface {
	ItemPrice(item string, numberOff int) (int, error)
}
type MultibuyPricer struct {
}

func (N MultibuyPricer) ItemPrice(item string, numberOff int) (int, error) {
	switch {
	case strings.EqualFold("A", item):
		units, multipacks := multiBreak(numberOff, 3)
		return (50 * units) + (130 * multipacks), nil
	case strings.EqualFold("B", item):
		units, multipacks := multiBreak(numberOff, 2)
		return (30 * units) + (45 * multipacks), nil
	case strings.EqualFold("C", item):
		return 20 * numberOff, nil
	case strings.EqualFold("D", item):
		return 15 * numberOff, nil
	default:
		return 0, fmt.Errorf("no price for item SKU %s", item)
	}
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
