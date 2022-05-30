package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"stock-aggregator/pkg/client"
	"testing"
)

func TestGetStockStats(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"some":"someValue"}`))
	}))
	defer server.Close()

	body, err := client.GetStockStats("fakeStock", "fakeAPI", server.URL+"/query?apikey=")
	if err != nil {
		t.Errorf("No err expected %s", err)
	}
	m := make(map[string]string)
	parseErr := json.Unmarshal(body, &m)
	if parseErr != nil {
		t.Errorf("No err expected")
	}
	if m["some"] == "data" {
		t.Errorf("Expected someValue Got: %s", m["some"])
	}
}

func TestRetryWhenNotSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err := client.GetStockStats("fakeStock", "fakeAPI", server.URL+"/query?apikey=")
	if err == nil {
		t.Errorf("err expected %s", err)
	}
}
