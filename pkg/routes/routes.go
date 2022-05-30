package routes

import (
	"net/http"
	h "stock-aggregator/pkg/handlers"
)

func Routes() {
	http.HandleFunc("/", h.HandleGetStockAverage)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
