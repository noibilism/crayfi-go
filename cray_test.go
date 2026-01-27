package crayfi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	// Test default initialization (fails without API Key)
	os.Unsetenv("CRAY_API_KEY")
	_, err := New("")
	if err == nil {
		t.Error("Expected error when API Key is missing")
	}

	// Test with API Key
	client, err := New("test_key")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if client.baseURL != SandboxURL {
		t.Errorf("Expected default Sandbox URL, got %s", client.baseURL)
	}

	// Test with Env Option
	client, _ = New("test_key", WithEnv("live"))
	if client.baseURL != LiveURL {
		t.Errorf("Expected Live URL, got %s", client.baseURL)
	}

	// Test with BaseURL Option (should override Env)
	customURL := "https://custom.api.com"
	client, _ = New("test_key", WithEnv("live"), WithBaseURL(customURL))
	if client.baseURL != customURL {
		t.Errorf("Expected Custom URL, got %s", client.baseURL)
	}
}

func TestCardsInitiate(t *testing.T) {
	// Mock Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/initiate" {
			t.Errorf("Expected path /api/v2/initiate, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected method POST, got %s", r.Method)
		}
		if r.Header.Get("Authorization") != "Bearer test_key" {
			t.Errorf("Expected Authorization header, got %s", r.Header.Get("Authorization"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success", "id": "123"}`))
	}))
	defer server.Close()

	client, _ := New("test_key", WithBaseURL(server.URL))
	
	resp, err := client.Cards.Initiate(map[string]interface{}{"amount": 100})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := resp.(map[string]interface{})
	if result["status"] != "success" {
		t.Errorf("Expected status success, got %v", result["status"])
	}
}

func TestPayoutsBanks(t *testing.T) {
	// Mock Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/payout/banks" {
			t.Errorf("Expected path /payout/banks, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.URL.Query().Get("countryCode") != "NG" {
			t.Errorf("Expected countryCode=NG, got %s", r.URL.Query().Get("countryCode"))
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"data": []string{"Bank A", "Bank B"}})
	}))
	defer server.Close()

	client, _ := New("test_key", WithBaseURL(server.URL))

	resp, err := client.Payouts.Banks("NG")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	result := resp.(map[string]interface{})
	data := result["data"].([]interface{})
	if len(data) != 2 {
		t.Errorf("Expected 2 banks, got %d", len(data))
	}
}

func TestErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Invalid Amount"}`))
	}))
	defer server.Close()

	client, _ := New("test_key", WithBaseURL(server.URL))

	_, err := client.Cards.Initiate(nil)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	apiErr, ok := err.(*APIException)
	if !ok {
		t.Fatalf("Expected APIException, got %T", err)
	}

	if apiErr.StatusCode != 400 {
		t.Errorf("Expected status 400, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "Invalid Amount" {
		t.Errorf("Expected message 'Invalid Amount', got '%s'", apiErr.Message)
	}
}
