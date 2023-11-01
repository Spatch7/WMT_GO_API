package main

import (
	"bytes"
	"net/http"
	"os"
	"testing"
)

func TestPost1(t *testing.T) {
	serverURL := "http://localhost:8080" // Change to match your server's URL
	client := &http.Client{}

	// Create a request with the desired JSON payload
	requestBody := `{
		"alert_id": "a950482e9911ec7e41f7ca5e5d9a424f",
		"service_id": "my_test_service_id",
		"service_name": "my_test_service",
		"model": "my_test_model",
		"alert_type": "anomaly",
		"alert_ts": "1695644160",
		"severity": "warning",
		"team_slack": "slack_ch"
	}`

	req, err := http.NewRequest("POST", serverURL+"/alerts", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		t.Errorf("Error creating HTTP request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending POST request: %v", err)
		return
	}

	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestPost2(t *testing.T) {
	serverURL := "http://localhost:8080" // Change to match your server's URL
	client := &http.Client{}

	// Create a request with the desired JSON payload
	requestBody := `{
		"alert_id": "bc7f8b55e62c0768763a19d3d0e43725",
		"service_id": "my_test_service_id",
		"service_name": "my_test_service",
		"model": "my_test_model",
		"alert_type": "anomaly",
		"alert_ts": "1695644160",
		"severity": "warning",
		"team_slack": "slack_ch"
	  }`

	req, err := http.NewRequest("POST", serverURL+"/alerts", bytes.NewBuffer([]byte(requestBody)))
	if err != nil {
		t.Errorf("Error creating HTTP request: %v", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending POST request: %v", err)
		return
	}

	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestMain(m *testing.M) {

	exitCode := m.Run()
	os.Exit(exitCode)
}
