package exchange

import (
	"fmt"

	"github.com/antchfx/xmlquery"
)

func fetchData(url string) (*xmlquery.Node, error) {
	return xmlquery.LoadURL(url)
}

func findDayData(data *xmlquery.Node, day string) *xmlquery.Node {
	daysData := xmlquery.Find(data, "//Cube")

	for i, v := range daysData {
		if v.SelectAttr("time") == day {
			fmt.Println(i, v.SelectAttr("time"))

			return v
		}
	}

	return nil
}

func findDayRate(dayData *xmlquery.Node, currencyISO string) (float64, error) {
	dayRates := xmlquery.Find(dayData, "//Cube")

	for i, v := range dayRates {
		if v.SelectAttr("currency") == currencyISO {
			fmt.Println(i, v.SelectAttr("rate"))

			return 0.00, nil
		}
	}

	return 0.00, nil
}

func getRate(url, currencyISO string) error {
	data, err := fetchData(url)
	if err != nil {
		return err
	}

	d := "2021-04-13"
	day := findDayData(data, d)
	findDayRate(day, currencyISO)

	return nil
}
