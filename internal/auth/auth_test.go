package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIeySuccess(t *testing.T){
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey testapikey123")

	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if apiKey != "testapikey123" {
		t.Errorf("expected 'testapikey123', got %s", apiKey)
	}
}

func TestGetAPIKeyNoAuthHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKeyMalformedAuthHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer sometoken")

	_, err := GetAPIKey(headers)
	if err == nil || err.Error() != "malformed authorization header" {
		t.Errorf("expected 'malformed authorization header' error, got %v", err)
	}
}