package commands

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"example.com/cryptocheck/api"
	"example.com/cryptocheck/utils"
)

func Check(symbols *string, fiat *string) error {
	normalizedSymbols := utils.NormalizedSymbols(*symbols)
	normalizedFiat, err := utils.NormalizeFiat(*fiat)

	if err != nil {
		return err
	}

	assetMetaData, wrongSymbols, err := api.FetchAssetMetadata(normalizedSymbols, normalizedFiat)

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

	if len(wrongSymbols) > 0 && len(assetMetaData) > 0 {
		fmt.Printf("next symbols doesn't exists: %s\n", strings.Join(wrongSymbols, ", "))
	} else if len(assetMetaData) == 0 {
		return errors.New("at least one -symbol flag argument is required")
	}

	return nil
}
