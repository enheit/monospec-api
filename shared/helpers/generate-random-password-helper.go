package helpers

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"monospec-api/shared/problems"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"
	length  = 16
)

func GenerateRandomPassword() (string, error) {
	password := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := range password {
		index, err := rand.Int(rand.Reader, charsetLength)

		if err != nil {
			originalErrorMessage := err.Error()

			problem := problems.Problem{
				Id:             "42648602-8613-4f90-ae2c-b96a172d12cc",
				Message:        "Something went wrong while generating password",
				Description:    &originalErrorMessage,
				HttpStatusCode: 500,
			}

			return "", problem
		}

		password[i] = charset[index.Int64()]
	}

	return string(password), nil
}
