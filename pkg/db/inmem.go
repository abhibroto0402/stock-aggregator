package db

import (
	"github.com/google/uuid"
	"stock-aggregator/pkg/domain"
	"time"
)

type db struct {
	symbol string
	apiKey string
	nDays  string
	apiURl string
}

var dbInstance *db
var FORMAT = "2006-01-02"

func Instantiate(symbol string, apiKey string, nDays string) {
	dbInstance = &db{
		symbol: symbol,
		apiKey: apiKey,
		nDays:  nDays,
		apiURl: "https://www.alphavantage.co/query?apikey=",
	}
}

var instance *domain.Record

// GetRecord Cache the record in memory; refresh if stale
func GetRecord() *domain.Record {
	if instance == nil || instance.Metadata.CreatedDate != time.Now().Format(FORMAT) {
		instance = BuildAverageStockRecord(dbInstance.symbol, dbInstance.apiKey, dbInstance.nDays, dbInstance.apiURl)
	}
	instance.Metadata.RequestId = uuid.New()
	return instance
}
