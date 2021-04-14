package exchange

import (
	"time"

	_ "github.com/antchfx/xmlquery"
	"github.com/pkg/errors"
)

type Exch struct{}

const url = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

func (*Exch) Rate(currencyISO, atDate string) (float64, error) {
	if errValidCurr := validateCurrency(currencyISO); errValidCurr != nil {
		return 0.00, errValidCurr
	}

	if errValidDate := validateDate(atDate); errValidDate != nil {
		return 0.00, errValidDate
	}

	return 0.00, nil
}

func validateCurrency(iso string) error {
	// TODO: add validation that passed is actual ISO symbol.

	if len(iso) != 3 {
		return errors.New("passed currency does has invalid length")
	}

	return nil
}

func validateDate(date string) error {
	layout := "2021-04-13"

	if _, errDate := time.Parse(layout, date); errDate != nil {
		return errors.WithMessagef(errDate, "passed \"%s\" is not a valid date", date)
	}

	return nil
}
