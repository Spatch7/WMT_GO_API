package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Alert struct {
	AlertID     string `json:"alert_id"`
	ServiceID   string `json:"service_id"`
	ServiceName string `json:"service_name"`
	Model       string `json:"model"`
	AlertType   string `json:"alert_type"`
	AlertTS     string `json:"alert_ts"`
	Severity    string `json:"severity"`
	TeamSlack   string `json:"team_slack"`
}

func main() {
	http.HandleFunc("/alerts", func(w http.ResponseWriter, r *http.Request) {
		// Post Method
		if r.Method == http.MethodPost {
			var alert Alert
			err := json.NewDecoder(r.Body).Decode(&alert)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Create file for alert
			file, err := os.Create("alerts/" + alert.AlertID + ".json")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer file.Close()

			// Convert go-struct to JSON
			encodedAlert, _ := json.Marshal(alert)
			_, err = file.Write(encodedAlert)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Alert data saved successfully")

			// Get Method
		} else if r.Method == http.MethodGet {

			serviceID := r.URL.Query().Get("service_id")
			startTS := r.URL.Query().Get("start_ts")
			endTS := r.URL.Query().Get("end_ts")

			if serviceID == "" || startTS == "" || endTS == "" {
				http.Error(w, " Query paramater missing", http.StatusBadRequest)
			}
			startTime, err := strconv.ParseInt(startTS, 10, 64)
			if err != nil {
				log.Printf("Error parsing start_ts: %v\n", err)
				http.Error(w, "Invalid start_ts", http.StatusBadRequest)
				return
			}

			endTime, err := strconv.ParseInt(endTS, 10, 64)
			if err != nil {
				log.Printf("Error parsing end_ts: %v\n", err)
				http.Error(w, "Invalid end_ts", http.StatusBadRequest)
				return
			}
			alerts := make([]Alert, 0)

			files, err := os.ReadDir("alerts")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			for _, file := range files {
				if strings.HasSuffix(file.Name(), ".json") {
					content, err := os.ReadFile("alerts/" + file.Name())
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					var alert Alert
					err = json.Unmarshal(content, &alert)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					alertTime, err := strconv.ParseInt(alert.AlertTS, 10, 64)
					if err != nil {
						http.Error(w, "Invalid alert timestamp", http.StatusInternalServerError)
						return
					}

					// Check for valid time-range
					if (alert.ServiceID == serviceID) && (alertTime >= startTime) && (alertTime <= endTime) {
						alerts = append(alerts, alert)
					}
				}
			}

			// Return all alerts with appropriate TS and serviceID
			encodedAlerts, err := json.Marshal(alerts)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(encodedAlerts)

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
