package commands

import (
	"errors"
	"fmt"
	"strings"

	"example.com/cryptocheck/api"
	"example.com/cryptocheck/utils"
)

type topListItem struct {
	name                    string
	symbol                  string
	circulatingMarketCapUSD float64
	conversionRate          float64
	fiatName                string
}

func Top(amount *int, fiat *string) error {
	normalizedFiat, err := utils.NormalizeFiat(*fiat)

	if err != nil {
		return err
	}

	rawTopList, err := api.FetchTopAssets(*amount, normalizedFiat)

	if err != nil {
		return err
	}

	topList, err := extractTopValues(rawTopList, *fiat)

	if err != nil {
		return err
	}

	outputTopList(topList, *amount)

	return nil
}

func extractTopValues(rawList []any, fiat string) ([]topListItem, error) {
	var cleanList []topListItem

	for _, value := range rawList {
		assetData, ok := value.(map[string]any)

		if ok {
			name, ok := assetData["URI"].(string)

			if !ok {
				return nil, errors.New("can't parse asset name")
			}

			symbol, ok := assetData["SYMBOL"].(string)
			if !ok {
				return nil, errors.New("can't parse asset symbol")
			}
			circulatingMarketCapUSD, ok := assetData["CIRCULATING_MKT_CAP_USD"].(float64)
			if !ok {
				return nil, errors.New("can't parse marketcap")
			}
			conversionRate, ok := assetData["PRICE_CONVERSION_RATE"].(float64)
			if !ok {
				return nil, errors.New("can't parse conversion rate")
			}

			newListItem := topListItem{
				name:                    name,
				symbol:                  symbol,
				circulatingMarketCapUSD: circulatingMarketCapUSD,
				conversionRate:          conversionRate,
				fiatName:                fiat,
			}

			cleanList = append(cleanList, newListItem)
		}
	}

	return cleanList, nil
}

func outputTopList(topList []topListItem, amount int) {
	for index, value := range topList {
		if index >= amount {
			break
		}
		circulatingMarketCap := value.circulatingMarketCapUSD * value.conversionRate
		name := strings.ToUpper(value.name[:1]) + strings.ToLower(value.name[1:])
		fmt.Printf("%v. %s (%s): %.2f %s\n", index+1, name, value.symbol, circulatingMarketCap, value.fiatName)
	}
}
