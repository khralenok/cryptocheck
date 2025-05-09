package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func FetchRates(baseCur string) (map[string]float64, error) {
	url := fmt.Sprintf("https://api.frankfurter.dev/v1/latest?base=%s", baseCur)

	data, err := http.Get(url)

	if err != nil {
		return nil, errors.New("can't get the response from Frankfurter")
	}

	defer data.Body.Close()

	body, err :=  io.ReadAll(data.Body)

	if err != nil {
		return nil, errors.New("can't read the data from Frankfurter response")
	}

	responseJson := make(map[string]interface{})
	rates := make(map[string]float64)

	err = json.Unmarshal(body, &responseJson)

	if err != nil {
		return nil, errors.New("can't convert the data")
	}

	switch valuesOfType := responseJson["rates"].(type) {
		case map[string]interface{}:
			for key, value := range valuesOfType{
				cleanValue, ok := value.(float64)
				if ok {
					rates[key] = cleanValue
				}			
			}
		default:
			return nil, errors.New("unknown type")
	}

	return rates, nil
}

func FetchAvailableCurrencies() (map[string]interface{}, error){
	data, err := http.Get("https://api.frankfurter.dev/v1/currencies")

	if err != nil {
		return nil, errors.New("can't get the response from Frankfurter")
	}

	defer data.Body.Close()

	body, err :=  io.ReadAll(data.Body)

	if err != nil {
		return nil, errors.New("can't read the data from Frankfurter response")
	}

	responseJson := make(map[string]interface{})
	err = json.Unmarshal(body, &responseJson)

	if err != nil {
		return nil, errors.New("can't convert the data")
	}

	return responseJson, nil
}