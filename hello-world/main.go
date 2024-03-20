package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

type WeatherData struct {
	Country       string  `json:"country"`
	City          string  `json:"city"`
	ConditionText string  `json:"conditionText"`
	TempC         float64 `json:"temp_c"`
}

func weatherDataHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var weather WeatherData
	if err := json.NewDecoder(r.Body).Decode(&weather); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received weather data: %+v\n", weather)

	// here we can work with weather data
	// for example add them to db

	// back response
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func main() {
	component := Hello("Sem")

	http.Handle("/", templ.Handler(component))
	http.HandleFunc("/weather-data", weatherDataHandler)

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
