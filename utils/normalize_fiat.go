package utils

import (
	"strings"
)

func NormalizeFiat(fiat string)(string, error){
	normalizedFiat := strings.Trim(fiat, " ")
	normalizedFiat = strings.ToUpper(normalizedFiat)
	return normalizedFiat, nil
}