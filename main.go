// A program made by Noah Calhoun
package main

import (
	"encoding/json"
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

type AlertRes struct {
	ServiceID   string  `json:"service_id"`
	ServiceName string  `json:"service_name"`
	Alerts      []Alert `json:"alerts"`
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

			// Check for missing required fields
			if alert.AlertID == "" || alert.ServiceID == "" || alert.AlertTS == "" {
				http.Error(w, "Missing required fields in the alert struct", http.StatusInternalServerError)
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
				response := struct {
					AlertID string `json:"alert_id"`
					Error   string `json:"error"`
				}{
					AlertID: alert.AlertID,
					Error:   err.Error(),
				}

				encodedResponse, _ := json.Marshal(response)
				http.Error(w, string(encodedResponse), http.StatusInternalServerError)
				return
			}

			// Return a success response struct
			response := struct {
				AlertID string `json:"alert_id"`
				Error   string `json:"error"`
			}{
				AlertID: alert.AlertID,
				Error:   "",
			}

			encodedResponse, _ := json.Marshal(response)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write(encodedResponse)

		} else if r.Method == http.MethodGet {
			// Get Method
			serviceID := r.URL.Query().Get("service_id")
			startTS := r.URL.Query().Get("start_ts")
			endTS := r.URL.Query().Get("end_ts")

			if serviceID == "" || startTS == "" || endTS == "" {
				http.Error(w, "Query paramater missing", http.StatusBadRequest)
			}

			startTime, err := strconv.ParseInt(startTS, 10, 64)
			if err != nil {
				// log.Printf("Error parsing start_ts: %v\n", err)
				http.Error(w, "Invalid start_ts", http.StatusBadRequest)
				return
			}

			endTime, err := strconv.ParseInt(endTS, 10, 64)
			if err != nil {
				// log.Printf("Error parsing end_ts: %v\n", err)
				http.Error(w, "Invalid start_ts", http.StatusBadRequest)
				return
			}

			var validAlerts []Alert
			files, err := os.ReadDir("alerts")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			alertRes := AlertRes{
				ServiceID:   serviceID,
				ServiceName: serviceID[:len(serviceID)-3],
			}

			// Parse through files for valid Alerts
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

					// Check for valid time-range, append to alerts array
					if (alert.ServiceID == serviceID) && (alertTime >= startTime) && (alertTime <= endTime) {
						validAlerts = append(validAlerts, alert)
					}
				}
			}

			// Return all alerts with appropriate TS and serviceID
			alertRes.Alerts = validAlerts
			encodedAlerts, err := json.Marshal(alertRes)
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
