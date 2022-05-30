package dto

import (
	"encoding/json"
	"log"
)

type Data struct {
	Metadata   map[string]string            `json:"Meta Data"`
	TimeSeries map[string]map[string]string `json:"Time Series (Daily)"`
}

func Parse(body []byte) (*Data, error) {
	dataObj := Data{}
	jsonErr := json.Unmarshal(body, &dataObj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return nil, jsonErr
	}
	return &dataObj, nil
}
