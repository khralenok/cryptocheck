package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type assetData map[string]map[string]interface{}

func FetchAssetMetadata(symbols []string, quoteAsset string) (assetData, []string, error) {
	url := fmt.Sprintf("https://data-api.coindesk.com/asset/v2/metadata?assets=%s&asset_lookup_priority=SYMBOL&quote_asset=%s&groups=BASIC,PRICE", strings.Join(symbols, ","), quoteAsset)

	response, err := getResponse(url)

	if err != nil {
		return nil, nil, err
	}

	err = handleError(response, quoteAsset)

	if err != nil {
		return nil, nil, err
	}

	normalizedData, err := handleData(response)

	if err != nil {
		return nil, nil, err
	}

	wrongSymbols := make([]string, 0)

	for _, key := range symbols {
		if normalizedData[key] == nil {
			wrongSymbols = append(wrongSymbols, key)
		}
	}

	return normalizedData, wrongSymbols, nil
}

func FetchTopAssets(amount int, fiat string) ([]any, error) {
	verifiedAmount := amount

	if amount < 10 {
		verifiedAmount = 10
	}

	if amount > 100 {
		verifiedAmount = 100
		fmt.Println("-amount value is too big. Here is top 100 assets:")
	}

	url := fmt.Sprintf("https://data-api.coindesk.com/asset/v1/top/list?page=1&page_size=%v&sort_by=CIRCULATING_MKT_CAP_USD&sort_direction=DESC&groups=ID,%%20PRICE,MKT_CAP&toplist_quote_asset=%s", verifiedAmount, fiat)

	response, err := getResponse(url)

	if err != nil {
		return nil, err
	}

	err = handleError(response, fiat)

	if err != nil {
		return nil, err
	}

	respData, ok := response["Data"].(map[string]any)

	if ok {
		respList, ok := respData["LIST"].([]any)
		if ok {
			return respList, nil
		}
	}

	return nil, errors.New("can't parse Coindesk API response")
}

func getResponse(url string) (map[string]interface{}, error) {
	data, err := http.Get(url)

	if err != nil {
		return nil, errors.New("can't get response from Coindesk API")
	}

	defer data.Body.Close()

	body, err := io.ReadAll(data.Body)

	if err != nil {
		return nil, errors.New("can't read the data from Coindesk API response")
	}

	var response map[string]any

	err = json.Unmarshal(body, &response)

	if err != nil {
		return nil, errors.New("can't convert the data from Coindesk API response")
	}

	return response, nil
}

func handleError(data map[string]any, quoteAsset string) error {
	dataError, ok := data["Err"].(map[string]any)

	if ok {
		errorTypeRaw, ok := dataError["type"].(float64)
		if ok {
			errorType := int(errorTypeRaw)

			switch errorType {
			case 1:
				if quoteAsset != "USD" {
					errorText := fmt.Sprintf("there is no such fiat currency as \"%s\"", quoteAsset)
					return errors.New(errorText)
				}
				return errors.New("there are no such symbols available")
			case 2:
				return errors.New("-symbol flag argument(s) are required")
			default:
				return errors.New("unexpected error")
			}
		}
	}

	return nil
}

func handleData(data map[string]any) (assetData, error) {
	normalizedData := make(assetData)

	switch v := data["Data"].(type) {
	case map[string]interface{}:
		for key, value := range v {
			interfacedValue, ok := value.(map[string]any)
			if ok {
				normalizedData[key] = interfacedValue
			}
		}
	default:
		return nil, errors.New("unexpected json stucture")
	}

	return normalizedData, nil
}
