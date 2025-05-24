package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
)

type RateResponse struct {
	Base string `json:"base"`
	Date string `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	url := "https://api.exchangerate-api.com/v4/latest/USD"

	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("HTTP hatası: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("API error: ", resp.Status)
		return
	}

	var data RateResponse

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("JSON error: ", err)
		return
	}

	fmt.Printf("Kur: %s\nTarih: %s\n\n", data.Base, data.Date)
	fmt.Printf("1 USD = %.2f TRY\n", data.Rates["TRY"])
}