package main

import (
	"os"
	"stock-aggregator/pkg/db"
	r "stock-aggregator/pkg/routes"
)

func main() {
	// Instantiate DB object using env variables
	db.Instantiate(os.Getenv("SYMBOL"), os.Getenv("API_KEY"), os.Getenv("NDAYS"))

	//Expose Endpoint for web-query
	r.Routes()
}
