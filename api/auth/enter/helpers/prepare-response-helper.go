package helpers

import (
	"encoding/json"
	"monospec-api/api/auth/enter/types"
)

func PrepareResponseBody(user *types.User) (*string, error) {
	responseBody := types.ResponseBody{
		User: user,
	}

	rawResponseBodyBytes, err := json.Marshal(responseBody)

	if err != nil {
		return nil, err
	}

	rawResponseBody := string(rawResponseBodyBytes)

	return &rawResponseBody, nil
}
