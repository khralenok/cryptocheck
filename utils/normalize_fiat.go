package utils

import (
	"errors"
	"fmt"
	"strings"

	"example.com/cryptocheck/api"
)

func NormalizeFiat(fiat string)(string, error){
	normalizedFiat := strings.Trim(fiat, " ")
	normalizedFiat = strings.ToUpper(normalizedFiat)
	err := validateFiat(normalizedFiat)

	if err != nil{
		return "", err
	}

	return normalizedFiat, nil
}

func validateFiat(fiatToValidate string)error{
	availableCur, err := api.FetchAvailableCurrencies()

	if err != nil {
		return errors.New("can't fetch the data")
	}

	if availableCur[fiatToValidate] == nil {
		errorText := fmt.Sprintf("there is no such fiat currency as \"%s\" available.\nuse help -fiat to see list of available currencies", fiatToValidate)
		return errors.New(errorText)
	}

	return nil
}