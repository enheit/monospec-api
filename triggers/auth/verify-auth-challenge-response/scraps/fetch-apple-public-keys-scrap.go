package scraps

import (
	"encoding/json"
	"fmt"
	"net/http"

	"monospec-api/triggers/auth/verify-auth-challenge-response/types"
)

func FetchApplePublicKeys() (*types.ApplePublicKeys, error) {
	res, err := http.Get("https://appleid.apple.com/auth/keys")

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch keys: status code %d", res.StatusCode)
	}

	var applePublicKeys types.ApplePublicKeys
	err = json.NewDecoder(res.Body).Decode(&applePublicKeys)

	if err != nil {
		return nil, err
	}

	return &applePublicKeys, nil
}
