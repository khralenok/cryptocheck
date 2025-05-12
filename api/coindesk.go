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

func FetchAssetMetadata(symbols []string, quoteAsset string) (assetData, error) {
	url := fmt.Sprintf("https://data-api.coindesk.com/asset/v2/metadata?assets=%s&asset_lookup_priority=SYMBOL&quote_asset=%s&groups=BASIC,PRICE", strings.Join(symbols, ","), quoteAsset)

	data, err := http.Get(url)

	if err != nil {
		return nil, errors.New("can't get response from Coindesk API")
	}

	defer data.Body.Close()

	body, err := io.ReadAll(data.Body)

	if err != nil {
		return nil, errors.New("can't read the data from Coindesk API response")
	}

	respJson := make(map[string]interface{})

	err = json.Unmarshal(body, &respJson)

	if err != nil {
		return nil, errors.New("unexpected response format")
	}

	normalizedData, err := normalizeData(respJson)

	if err != nil {
		return nil, err
	}

	return normalizedData, nil
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
