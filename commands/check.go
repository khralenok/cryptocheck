package commands

import (
	"errors"
	"fmt"
	"strconv"

	"example.com/cryptocheck/api"
	"example.com/cryptocheck/utils"
)

//To do:
//1. Handle the case if all symbols are wrong
//2. Handle the case if -fiat value is wrong

func Check(symbols *string, fiat *string) error {
	normalizedSymbols := utils.NormalizedSymbols(*symbols)
	normalizedFiat := utils.NormalizeFiat(*fiat)

	assetMetaData, err := api.FetchAssetMetadata(normalizedSymbols, normalizedFiat)

	if err != nil {
		return err
	}

	for key, value := range assetMetaData {
		if key == "" {
			return errors.New("wrong symbol name. try valid symbol please")
		}
		if value != nil {
			priceAsString := fmt.Sprint(value["PRICE_CONVERSION_VALUE"])
			price, err := strconv.ParseFloat(priceAsString, 64)

			if err != nil {
				return errors.New("can't convert the price")
			}

			fmt.Printf("%s(%s): %.4f %s\n", value["NAME"], key, price, normalizedFiat)
			continue
		}
	}

	return nil
}
