package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type assetData map[string]map[string]interface{}

func FetchAssetMetadata(symbols []string, quoteAsset string) (assetData, []string, error) {
	url := fmt.Sprintf("https://data-api.coindesk.com/asset/v2/metadata?assets=%s&asset_lookup_priority=SYMBOL&quote_asset=%s&groups=BASIC,PRICE", strings.Join(symbols, ","), quoteAsset)

	data, err := http.Get(url)

	if err != nil {
		return nil, nil, errors.New("can't get response from Coindesk API")
	}

	defer data.Body.Close()

	body, err := io.ReadAll(data.Body)

	if err != nil {
		return nil, nil, errors.New("can't read the data from Coindesk API response")
	}

	respJson := make(map[string]interface{})

	err = json.Unmarshal(body, &respJson)

	if err != nil {
		return nil, nil, errors.New("query was builded wrong")
	}

	err = handleError(respJson, quoteAsset)

	if err != nil {
		return nil, nil, err
	}

	normalizedData, err := normalizeData(respJson)

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

func normalizeData(data map[string]interface{}) (assetData, error) {
	normalizedData := make(assetData)

	switch v := data["Data"].(type) {
	case map[string]interface{}:
		for key, value := range v {
			interfacedValue, ok := value.(map[string]interface{})
			if ok {
				normalizedData[key] = interfacedValue
			}
		}
	default:
		return nil, errors.New("unexpected json stucture")
	}

	return normalizedData, nil
}

func handleError(data map[string]interface{}, quoteAsset string) error {
	switch v := data["Err"].(type) {
	case map[string]interface{}:
		switch innerV := v["type"].(type) {
		case interface{}:
			errorType := fmt.Sprint(innerV)
			errorTypeClean, err := strconv.ParseInt(errorType, 0, 0)

			if err != nil {
				return err
			}

			if errorTypeClean == 1 {
				if quoteAsset != "USD" {
					errorText := fmt.Sprintf("there is no such fiat currency as \"%s\"", quoteAsset)
					return errors.New(errorText)
				}
				return errors.New("there are no such symbols available")
			}
			if errorTypeClean == 2 {
				return errors.New("-symbol flag argument(s) are required")
			}
		}
	}
	return nil
}
