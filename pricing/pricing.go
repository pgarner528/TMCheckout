package pricing

import (
	"fmt"
	"strings"
)

type Pricer interface {
	ItemPrice(item string, numberOff int) (int, error)
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
