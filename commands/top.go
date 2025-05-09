package commands

import (
	"fmt"

	"example.com/cryptocheck/utils"
)

func Top(amount *int,fiat *string) error{
	normalizedFiat, err := utils.NormalizeFiat(*fiat)

	if err != nil {
		return err
	}

	fmt.Println("Top command")
	fmt.Println("Amount: ", *amount)
	fmt.Println("Fiat: ", normalizedFiat)

	return nil
}