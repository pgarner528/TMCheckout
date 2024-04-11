package pricing

type Pricer interface {
	ItemPrice(item string, numberOff int) (int, error)
}

type NominalPricer struct {
}

func (N NominalPricer) ItemPrice(item string, numberOff int) (int, error) {
	return 1 * numberOff, nil
}
