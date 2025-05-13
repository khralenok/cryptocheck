package commands

import (
	"fmt"

	"example.com/cryptocheck/api"
	"example.com/cryptocheck/utils"
)

func Top(amount *int, fiat *string) error {
	normalizedFiat, err := utils.NormalizeFiat(*fiat)

	if err != nil {
		return err
	}

	topData, err := api.FetchTopAssets(*amount, normalizedFiat)

	if err != nil {
		return err
	}

	for key, value := range topData {
		fmt.Printf("%s: %v\n", key, value)
	}

	return nil
}
