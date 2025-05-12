package utils

import (
	"strings"
)

func NormalizeFiat(fiat string) string {
	normalizedFiat := strings.Trim(fiat, " ")
	normalizedFiat = strings.ToUpper(normalizedFiat)
	return normalizedFiat
}
