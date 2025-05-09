package commands

import (
	"fmt"

	"example.com/cryptocheck/utils"
)

func Check(symbols *string, fiat *string)error{
	listOfSymbols := utils.NormalizedSymbols(*symbols)
	normalizedFiat, err := utils.NormalizeFiat(*fiat)

	if err != nil {
		return err
	}

	for _, value := range listOfSymbols{
		fmt.Printf("%s: 1000.00 %s\n", value, normalizedFiat)
	}

	return nil
}