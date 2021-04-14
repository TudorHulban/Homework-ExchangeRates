package exchange

type Exch struct{}

const url = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

func (*Exch) Rate(currencyISO, atDate string) (float64, error) {
	if errValidCurr := validateCurrency(currencyISO); errValidCurr != nil {
		return 0.00, errValidCurr
	}

	if errValidDate := validateDate(atDate); errValidDate != nil {
		return 0.00, errValidDate
	}

	getRate(url, currencyISO)

	return 0.00, nil
}
