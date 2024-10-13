package scraps

import (
	"monospec-api/triggers/auth/verify-auth-challenge-response/types"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyAppleIdToken(token string, publicKeys *types.ApplePublicKeys) error {
	parsedToken, err := jwt.Parse(token, keyFunc)

	println(parsedToken)

	if err != nil {
		return err
	}

	return nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	return nil, nil
}

// base64(header.json).base64(payload.json()).base64(hmac(SECRET, header+payload))
