package utils

import "strings"

func NormalizedSymbols(input string) []string {
	listOfSymbols := strings.Split(input, ",")

	for i := 0; i < len(listOfSymbols); i++ {
		listOfSymbols[i] = strings.Trim(listOfSymbols[i], " ")
		listOfSymbols[i] = strings.ToUpper(listOfSymbols[i])
	}

	return listOfSymbols
}
