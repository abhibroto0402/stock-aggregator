package db

import (
	"errors"
	"fmt"
	"log"
	"stock-aggregator/pkg/client"
	"stock-aggregator/pkg/domain"
	"stock-aggregator/pkg/dto"
	"strconv"
	"time"
)

func BuildAverageStockRecord(stockName string, apiKey string, nDays string, url string) *domain.Record {
	numDays, err := strconv.Atoi(nDays)
	if err != nil {
		log.Fatal("[ERROR] Number of Days has to be a number")
		return nil
	}
	body, err := client.GetStockStats(stockName, apiKey, url)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	data, err := dto.Parse(body)
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	nDaysData, err := getNdaysData(data, numDays)
	nDaysAverageClosingPrice := getNDaysAverageClosingPrice(nDaysData)
	log.Printf("[INFO] Average Closing Price for %d days: %f\n", numDays, nDaysAverageClosingPrice)
	log.Println(nDaysData)
	return recordBuilderHelper(nDaysData, nDaysAverageClosingPrice, stockName)
}

func recordBuilderHelper(data []map[string]string, avgPrice float64, stockName string) *domain.Record {
	return &domain.Record{
		Metadata: domain.Metadata{
			Days:                len(data),
			AverageClosingPrice: avgPrice,
			Stock:               stockName,
			CreatedDate:         time.Now().Format(FORMAT),
		},
		Daily: data,
	}
}

func getNDaysAverageClosingPrice(data []map[string]string) float64 {
	sum := 0.0
	count := 0
	for _, items := range data {
		if cp, err := strconv.ParseFloat(items["4. close"], 64); err == nil {
			sum += cp
		}
		count++
	}
	return sum / float64(count)

}

func getNdaysData(data *dto.Data, days int) ([]map[string]string, error) {
	date := time.Now()
	var tradingDate, err = getPreviousTradingDate(date.AddDate(0, 0, 1), data)
	if err != nil {
		return nil, err
	}
	var nDaysData []map[string]string

	for i := 1; len(nDaysData) != days && i <= days; i++ {
		value, ok := data.TimeSeries[tradingDate.Format(FORMAT)]
		if ok {
			log.Printf("[INFO] Data obtained for %s \n", tradingDate.Format(FORMAT))
			value["date"] = tradingDate.Format(FORMAT)
			nDaysData = append(nDaysData, value)
		}
		tradingDate, err = getPreviousTradingDate(tradingDate.AddDate(0, 0, -i), data)
		if err != nil {
			log.Printf("[WARN] No trading date found in data from %s", tradingDate.Format(FORMAT))
			break
		}
	}
	if len(nDaysData) == 0 {
		return nil, fmt.Errorf("no data found for last %d days", days)
	}
	log.Printf("[INFO] Processed data for %d days\n", len(nDaysData))
	return nDaysData, nil
}

/**
Due to public holidays and/or weekends the last trading date could be more than one day ago the current date. Hence, we
should look for valid trading date going back one day. Upto 7 days. This is assuming there cannot be two consecutive weeks
during which no trading has happened. This can be changed as per business rules.
*/

func getPreviousTradingDate(date time.Time, data *dto.Data) (time.Time, error) {
	for count := 0; count < 7; count++ {
		var currDate = date.AddDate(0, 0, -count)
		_, ok := data.TimeSeries[currDate.Format(FORMAT)]
		if ok {
			return currDate, nil
		}
	}
	return date, errors.New("no matching data found in last 7 days")

}
