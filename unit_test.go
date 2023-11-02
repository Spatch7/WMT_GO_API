package main

import (
	"bytes"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestPost1(t *testing.T) {
	serverURL := "http://localhost:8080"
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
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestPost2(t *testing.T) {
	serverURL := "http://localhost:8080"
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
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, resp.StatusCode)
	}
}

func TestPost3(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Create a request with the desired JSON payload
	requestBody := `{
		"alert_id": "cc7f8b55e62c0768763a19d3d0e43725",
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
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d", http.StatusCreated, resp.StatusCode)
	}
}

// Test Error 400
func TestPost4(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Create a request with the desired JSON payload
	requestBody := `{
		"alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
		"service_id": "my_test_service_id",
		"service_name": "my_test_service",
		"model": "my_test_model",
		"alert_type": "anomaly",
		"alert_ts": "invalid_ts", // Trigger an error by providing an invalid timestamp
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
	if http.StatusBadRequest != resp.StatusCode {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

// Test Error 500
func TestPost5(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Create a request with an incomplete JSON payload (missing required fields)
	requestBody := `{
		"alert_id": "b950482e9911ec7e41f7ca5e5d9a424f"
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
	if http.StatusInternalServerError != resp.StatusCode {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func TestGet1(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Create a GET request with the desired query parameters
	req, err := http.NewRequest("GET", serverURL+"/alerts?service_id=my_test_service_id&start_ts=0&end_ts=9999999999", nil)
	if err != nil {
		t.Errorf("Error creating GET request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestGet2(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Create a GET request with the desired query parameters
	req, err := http.NewRequest("GET", serverURL+"/alerts?service_id=my_test_service_id&start_ts=0&end_ts=9999999999", nil)
	if err != nil {
		t.Errorf("Error creating GET request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestGetFilesCreatedInPast45Minutes(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Calculate ts for the past 45 minutes
	endTime := time.Now().Unix()
	startTime := endTime - 2700

	// Create a GET request with the desired query parameters
	req, err := http.NewRequest("GET", serverURL+"/alerts?service_id=my_test_service_id&start_ts="+strconv.FormatInt(startTime, 10)+"&end_ts="+strconv.FormatInt(endTime, 10), nil)
	if err != nil {
		t.Errorf("Error creating GET request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}

// Error status 400
func TestGetBadRequest(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Create a GET request with missing or incorrect query parameters
	req, err := http.NewRequest("GET", serverURL+"/alerts", nil)
	if err != nil {
		t.Errorf("Error creating GET request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

// Error status 404
func TestGetNotFound(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Create a GET request for a non-existent resource
	req, err := http.NewRequest("GET", serverURL+"/nonexistent_resource", nil)
	if err != nil {
		t.Errorf("Error creating GET request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, resp.StatusCode)
	}
}

// Error status 500
func TestGetInternalServerError(t *testing.T) {
	serverURL := "http://localhost:8080"
	client := &http.Client{}

	// Define the path to the "alerts" folder
	alertsFolder := "alerts"

	// Delete the "alerts" folder
	if err := os.RemoveAll(alertsFolder); err != nil {
		t.Errorf("Error deleting the 'alerts' folder: %v", err)
		return
	}

	// Create a GET request that intentionally triggers an internal server error
	req, err := http.NewRequest("GET", serverURL+"/alerts?service_id=my_test_service_id&start_ts=0&end_ts=9999999999", nil)
	if err != nil {
		t.Errorf("Error creating GET request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error sending GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Verify the response status code
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, resp.StatusCode)
	}

	// Recreate the "alerts" folder
	if err := os.MkdirAll(alertsFolder, os.ModePerm); err != nil {
		t.Errorf("Error recreating the 'alerts' folder: %v", err)
	}
}

func TestMain(m *testing.M) {

	exitCode := m.Run()
	os.Exit(exitCode)
}
