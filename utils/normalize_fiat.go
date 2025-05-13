package utils

import (
	"errors"
	"strings"
)

func NormalizeFiat(input string) (string, error) {
	arguments := strings.Split(input, ",")

	if len(arguments) > 1 {
		return "", errors.New("only one curency accepted as -fiat flag argument")
	}

	normalizedFiat := strings.Trim(input, " ")
	normalizedFiat = strings.ToUpper(normalizedFiat)
	return normalizedFiat, nil
}
