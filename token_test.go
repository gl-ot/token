package token

import (
	"net/http"
	"reflect"
	"testing"
)

const (
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
)

var tokenClaims = map[string]interface{} {
	"sub": "1234567890",
	"name": "John Doe",
	"iat": 1516239022,
}

func TestExtractClaims(t *testing.T) {
	claims, err := ExtractClaims(buildRequest())
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(claims, tokenClaims) {
		t.Errorf("extracted claims: got=%s, expected=%s", claims, tokenClaims)
	}
}

func TestExtractTokenClaims(t *testing.T) {
	claims, err := ExtractTokenClaims(token)
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(claims, tokenClaims) {
		t.Errorf("extracted claims: got=%s, expected=%s", claims, tokenClaims)
	}
}

func TestExtractToken(t *testing.T) {
	got, err := ExtractToken(buildRequest())
	if err != nil {
		t.Error(err)
	}

	if got != token {
		t.Errorf("extracted token: got=%s, expected=%s", got, token)
	}
}

func buildRequest() *http.Request {
	r, _ := http.NewRequest("GET", "http://example.com", nil)
	r.Header.Set(authHeaderKey, "Bearer " + token)
	return r
}
