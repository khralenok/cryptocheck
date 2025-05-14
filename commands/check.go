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

	err = outputCheckList(assetMetaData, wrongSymbols, normalizedFiat)

	if err != nil {
		return err
	}

	return nil
}

func outputCheckList(assetMetaData map[string]map[string]any, wrongSymbols []string, fiat string) error {
	for key, value := range assetMetaData {
		if key == "" {
			return errors.New("wrong symbol name. try valid symbol please")
		}
		if value != nil {
			price, ok := value["PRICE_CONVERSION_VALUE"].(float64)

			if ok {
				fmt.Printf("%s(%s): %.4f %s\n", value["NAME"], key, price, fiat)
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
