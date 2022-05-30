package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"stock-aggregator/pkg/db"
	"testing"
	"time"
)

func TestGetRecords(t *testing.T) {
	var date1 = time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	var date2 = time.Now().AddDate(0, 0, -2).Format("2006-01-02")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{
   "Meta Data":{
      "1. Information":"Daily Prices (open, high, low, close) and Volumes",
      "2. Symbol":"MSFT",
      "3. Last Refreshed":"2022-05-27",
      "4. Output Size":"Compact",
      "5. Time Zone":"US/Eastern"
   },
   "Time Series (Daily)":{
      "%s":{
         "1. open":"268.4800",
         "2. high":"273.3400",
         "3. low":"267.5600",
         "4. close":"273.2400",
         "5. volume":"26910806"
      },
      "%s":{
         "1. open":"262.2700",
         "2. high":"267.1100",
         "3. low":"261.4294",
         "4. close":"265.9000",
         "5. volume":"25002105"
      }
   }
}`, date1, date2)))
	}))
	record := db.BuildAverageStockRecord("fakeStock", "fakeAPIKey", "3", server.URL+"/query?apikey=")
	if record.Metadata.Days != 2 {
		t.Errorf("Expected number of days 2 found %d", record.Metadata.Days)
	}
}
