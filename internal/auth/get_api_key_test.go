package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "ApiKey testtest234")
	key, err := GetAPIKey(req.Header)
	want := "testtest234"
	if err != nil {
		t.Fatalf("Function output error %v", err)
	}
	if key != want {
		t.Fatalf("expected: %v, got: %v", want, key)
	}
}
