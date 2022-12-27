package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/exchange-rates", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		rates := generateExchangeRates()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rates)
	})

	http.ListenAndServe(":8001", nil)
}

func generateExchangeRates() map[string]float64 {
	exchangeRates := make(map[string]float64)
	exchangeRates["GBP"] = roundToFourDecimals(0.85 + rand.Float64()*(0.89-0.85))
	exchangeRates["CHF"] = roundToFourDecimals(1.07 + rand.Float64()*(1.11-1.07))
	exchangeRates["SEK"] = roundToFourDecimals(10.50 + rand.Float64()*(10.60-10.50))
	exchangeRates["NOK"] = roundToFourDecimals(10.10 + rand.Float64()*(10.20-10.10))
	exchangeRates["DKK"] = roundToFourDecimals(7.40 + rand.Float64()*(7.50-7.40))
	exchangeRates["USD"] = roundToFourDecimals(1.15 + rand.Float64()*(1.19-1.15))
	exchangeRates["CAD"] = roundToFourDecimals(1.52 + rand.Float64()*(1.56-1.52))
	exchangeRates["AUD"] = roundToFourDecimals(1.55 + rand.Float64()*(1.59-1.55))
	exchangeRates["JPY"] = roundToFourDecimals(130.00 + rand.Float64()*(136.00-130.00))
	exchangeRates["INR"] = roundToFourDecimals(80.00 + rand.Float64()*(90.00-80.00))
	exchangeRates["SGD"] = roundToFourDecimals(1.60 + rand.Float64()*(1.70-1.60))
	exchangeRates["HKD"] = roundToFourDecimals(9.40 + rand.Float64()*(9.50-9.40))
	exchangeRates["CNY"] = roundToFourDecimals(7.80 + rand.Float64()*(7.90-7.80))
	return exchangeRates
}

func roundToFourDecimals(num float64) float64 {
	return float64(int(num*10000+0.5)) / 10000
}
