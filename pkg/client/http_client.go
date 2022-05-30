package client

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func GetStockStats(stockName string, apiKey string, baseUrl string) ([]byte, error) {
	url := baseUrl + apiKey + "&function=TIME_SERIES_DAILY&symbol=" + stockName
	retry := 5
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	var statusCode = 0
	for retry > 0 {
		response, err := client.Get(url)
		if err == nil && response.StatusCode == 200 {
			body, readErr := ioutil.ReadAll(response.Body)
			if readErr != nil {
				log.Fatal(readErr)
				return nil, readErr
			}
			return body, nil
		}
		log.Printf("Error:%s", err)
		log.Printf("Response Code: %d", response.StatusCode)
		time.Sleep(10 * time.Second)
		log.Printf("Retry #%d", retry)
		statusCode = response.StatusCode
		retry--

	}
	return nil, fmt.Errorf("response Code:%d", statusCode)
}
