package exchange

// Exch provides a structure to whom we can add a method.
// A structure was added for convenience (maybe elaborate with a constructor and refining config),
// we could maybe delete the receiver and use the function directly.
type Exch struct{}

const url = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

func (*Exch) ConvertedAmount(amount float64, currencyISO, forDate string) (float64, error) {
	if errValidCurr := validateCurrency(currencyISO); errValidCurr != nil {
		return 0.00, errValidCurr
	}

	if errValidDate := validateDate(forDate); errValidDate != nil {
		return 0.00, errValidDate
	}

	rate, errGet := getRate(url, currencyISO, forDate)
	if errGet != nil {
		return 0.00, errGet
	}

	return amount * rate, nil
}
