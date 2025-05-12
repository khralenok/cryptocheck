package commands

import (
	"fmt"
	"strconv"

	"example.com/cryptocheck/api"
	"example.com/cryptocheck/utils"
)

func Check(symbols *string, fiat *string) error {
	NormalizedSymbols := utils.NormalizedSymbols(*symbols)
	normalizedFiat, err := utils.NormalizeFiat(*fiat)

	if err != nil {
		return err
	}

	assetMetaData, err := api.FetchAssetMetadata(NormalizedSymbols, normalizedFiat)

	if err != nil {
		return err
	}

	for key, value := range assetMetaData {
		priceAsString := fmt.Sprint(value["PRICE_CONVERSION_VALUE"])
		price, err := strconv.ParseFloat(priceAsString, 64)

		if err != nil {
			return err
		}

		fmt.Printf("%s(%s): %.4f %s\n", value["NAME"], key, price, normalizedFiat)
	}

	return nil
}
