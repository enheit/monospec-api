package services

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"monospec-api/auth/api/apple/types"
	"monospec-api/shared/problems"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AppleIdentityTokenService struct {
}

func NewAppleIdentityTokenService() *AppleIdentityTokenService {
	return &AppleIdentityTokenService{}
}

func (s *AppleIdentityTokenService) VerifyIdentityToken(rawIdentityToken string, publicKeys *types.ApplePublicKeys) (*types.AppleIdentityToken, error) {
	println("rawIdentityToken: ", rawIdentityToken)

	// HERE

	token, err := jwt.ParseWithClaims(rawIdentityToken, &types.AppleIdentityTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, unexpectedSigningMethod(token.Header["alg"].(string))
		}

		for _, key := range publicKeys.Keys {
			if key.Kid == token.Header["kid"] {
				parsedKey, err := parsePublicKey(key)

				if err != nil {
					return nil, err
				}

				return parsedKey, nil
			}
		}

		return nil, kidNotFound(token.Header["kid"].(string))
	})

	println("token: ", token)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*types.AppleIdentityTokenCustomClaims)

	if !ok {
		return nil, invalidAppleIdentityToken()
	}

	if err != nil {
		return nil, err
	}

	keyId, ok := token.Header["kid"].(string)

	if !ok {
		return nil, kidNotFound(token.Header["kid"].(string))
	}

	issuer, err := token.Claims.GetIssuer()

	if err != nil {
		return nil, err
	}

	if issuer != "https://appleid.apple.com" {
		return nil, NewInvalidIssuer(issuer)
	}

	audience, err := token.Claims.GetAudience()

	if err != nil {
		return nil, err
	}

	expiresAt, err := token.Claims.GetExpirationTime()

	if err != nil {
		return nil, err
	}

	issuedAt, err := token.Claims.GetIssuedAt()

	if err != nil {
		return nil, err
	}

	if expiresAt.Before(time.Now()) {
		return nil, tokenExpired()
	}

	// HINT: It seems like this app is the only audience for apple's jwt @ Roman
	if audience[0] != "app.monospec.Monospec" {
		return nil, invalidAudience(audience[0])
	}

	appleIdentityToken := &types.AppleIdentityToken{
		Header: types.AppleIdentityTokenHeader{
			Algorithm: token.Method.Alg(),
			KeyId:     keyId,
		},
		Payload: types.AppleIdentityTokenPayload{
			Subject:         claims.Subject,
			Email:           claims.Email,
			IsEmailVerified: claims.EmailVerified,
			AuthenticatedAt: time.Unix(claims.AuthTime, 0),
			IssuedAt:        issuedAt.Time,
		},
	}

	return appleIdentityToken, nil
}

func parsePublicKey(applePublicKey types.ApplePublicKey) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(applePublicKey.N)

	if err != nil {
		return nil, err
	}

	eBytes, err := base64.RawURLEncoding.DecodeString(applePublicKey.E)

	if err != nil {
		return nil, err
	}

	pubKey := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nBytes),
		E: int(new(big.Int).SetBytes(eBytes).Int64()),
	}

	return pubKey, nil
}

func (s *AppleIdentityTokenService) GetPublicKeys() (*types.ApplePublicKeys, error) {
	response, err := http.Get("https://appleid.apple.com/auth/keys")

	if err != nil {
		originalErrorMessage := err.Error()

		problem := problems.Problem{
			Id:             "2204db26-1035-43d3-9e1d-46fedf7c6177",
			Message:        "Failed to get apple public keys",
			Description:    &originalErrorMessage,
			HttpStatusCode: 503,
		}

		return nil, problem
	}

	defer response.Body.Close()

	var publicKeys types.ApplePublicKeys

	if err := json.NewDecoder(response.Body).Decode(&publicKeys); err != nil {
		return nil, err
	}

	return &publicKeys, nil
}

func tokenExpired() error {
	description := "Identity token expired"

	problem := problems.Problem{
		Id:             "d3c9f5b8-3d9f-4c7f-8d9f-9d9f9d9f9d9f",
		Message:        description, // "Failed to verify apple identity token",
		Description:    &description,
		HttpStatusCode: 401,
	}

	return problem
}

func invalidAudience(audience string) error {
	originalErrorMessage := fmt.Sprintf("Unexpected audience: %s", audience)

	problem := problems.Problem{
		Id:             "c35470ea-6aea-4ae1-93e4-503d1cae88f8",
		Message:        originalErrorMessage, // "Failed to verify apple identity token",
		Description:    &originalErrorMessage,
		HttpStatusCode: 401,
	}

	return problem
}

type InvalidIssuer struct {
	problem problems.Problem
}

func (e *InvalidIssuer) Error() string {
	return e.problem.Message
}

func NewInvalidIssuer(issuer string) *InvalidIssuer {
	description := fmt.Sprintf("Unexpected issuer: %s", issuer)

	problem := problems.Problem{
		Id:             "c746ec2e-ff93-4b1e-a6fc-046e9daa7832",
		Message:        description, // "Failed to verify apple identity token",
		Description:    &description,
		HttpStatusCode: 401,
	}

	return &InvalidIssuer{problem: problem}
}

// func invalidIssuer(issuer string) error {
// 	originalErrorMessage := fmt.Sprintf("Unexpected issuer: %s", issuer)
//
// 	problem := problems.Problem{
// 		Id:             "c746ec2e-ff93-4b1e-a6fc-046e9daa7832",
// 		Message:        "Failed to verify apple identity token",
// 		Description:    &originalErrorMessage,
// 		HttpStatusCode: 401,
// 	}
//
// 	return problem
// }

func unexpectedSigningMethod(alg string) error {
	description := fmt.Sprintf("Unexpected signing method: %v", alg)

	problem := problems.Problem{
		Id:             "8ab2be7c-3e8c-48cd-a670-d8f725ebe5c3",
		Message:        description, // "Failed to verify apple identity token",
		Description:    &description,
		HttpStatusCode: 401,
	}

	return problem
}

func kidNotFound(kid string) error {
	originalErrorMessage := fmt.Sprintf("Key with kid %s not found", kid)

	problem := problems.Problem{
		Id:             "ccf6182c-54fe-4608-8f14-9b400cee6f15",
		Message:        originalErrorMessage, // "Failed to verify apple identity token",
		Description:    &originalErrorMessage,
		HttpStatusCode: 401,
	}

	return problem
}

func invalidAppleIdentityToken() error {
	description := "Failed to extract claims fromt apple's identity token"

	problem := problems.Problem{
		Id:             "4ac758ae-eb73-426e-ad55-4440216c203e",
		Message:        description, // "Failed to verify apple identity token",
		Description:    &description,
		HttpStatusCode: 401,
	}

	return problem
}
