package commands

import (
	"errors"
	"fmt"
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
			price, ok := value["PRICE_CONVERSION_VALUE"].(float64)

			if ok {
				fmt.Printf("%s(%s): %.4f %s\n", value["NAME"], key, price, normalizedFiat)
				continue
			}
		}
	}

	if len(wrongSymbols) > 0 && len(assetMetaData) > 0 {
		fmt.Printf("next symbols doesn't exists: %s\n", strings.Join(wrongSymbols, ", "))
	} else if len(assetMetaData) == 0 {
		return errors.New("at least one -symbol flag argument is required")
	}

	return nil
}
