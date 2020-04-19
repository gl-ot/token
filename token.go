package token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

const (
	authHeaderKey = "Authorization"
)

func ExtractClaims(r *http.Request) (map[string]interface{}, error) {
	tokenStr, err := ExtractToken(r)
	if err != nil {
		return nil, err
	}

	return ExtractTokenClaims(tokenStr)
}

func ExtractTokenClaims(token string) (map[string]interface{}, error) {
	if len(strings.TrimSpace(token)) == 0 {
		return nil, errors.New("empty token")
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, nil)
	if err != nil && err.Error() != "no Keyfunc was provided." {
		return nil, err
	}

	return claims, nil
}

func ExtractToken(r *http.Request) (string, error) {
	header := strings.TrimSpace(r.Header.Get(authHeaderKey))
	if strings.HasPrefix(header, "Bearer") {
		return strings.TrimSpace(header[6:]), nil
	} else {
		return "", errors.New("couldn't extract token from Authorization header")
	}
}