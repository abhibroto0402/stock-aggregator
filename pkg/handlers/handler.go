package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"stock-aggregator/pkg/db"
	resp "stock-aggregator/pkg/domain"
	"time"
)

func HandleGetStockAverage(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	jsonBytes, err := json.Marshal(newResponse())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	duration := time.Since(start)
	log.Printf("[INFO] Time taken by request:%fs\n", duration.Seconds())
}

func newResponse() *resp.Record {
	record := db.GetRecord()
	return record
}
