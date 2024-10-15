package services

import (
	"fmt"
	"monospec-api/auth/types"
	"monospec-api/shared/problems"
	"time"

	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	PrivateKey string

	issuer   string
	audience string
}

func NewTokenService(privateKey string) *TokenService {
	issuer := "https://monospec.app"
	audience := "https://api.monospec.app"

	return &TokenService{
		PrivateKey: privateKey,

		issuer:   issuer,
		audience: audience,
	}
}

func (t *TokenService) CreateAccessToken(userId string) (*string, error) {
	createdAt := time.Now()
	expiresAt := createdAt.Add(time.Minute * 5)

	claims := jwt.MapClaims{
		"iss": t.issuer,
		"sub": userId,
		"jti": uuid.New().String(),
		"aud": t.audience,
		"iat": createdAt.Unix(),
		"exp": expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	rawAccessToken, err := token.SignedString(t.PrivateKey)

	if err != nil {
		originalErrorMessage := err.Error()

		problem := problems.Problem{
			Id:             "bb563e4e-e2eb-420b-8293-804de8171171",
			Message:        "Failed to sign access token",
			Description:    &originalErrorMessage,
			HttpStatusCode: 500,
		}

		return nil, problem
	}

	return &rawAccessToken, nil
}

func (t *TokenService) CreateRefreshToken(userId string) (*string, error) {
	createdAt := time.Now()
	expiresAt := createdAt.Add(time.Hour * 1)

	claims := jwt.MapClaims{
		"iss": t.issuer,
		"sub": userId,
		"jti": uuid.New().String(),
		"aud": t.audience,
		"iat": createdAt.Unix(),
		"exp": expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	rawRefreshToken, err := token.SignedString(t.PrivateKey)

	if err != nil {
		originalErrorMessage := err.Error()

		problem := problems.Problem{
			Id:             "0bbb3cbf-84de-4ad0-b153-2ff74f111e12",
			Message:        "Failed to sign refresh token",
			Description:    &originalErrorMessage,
			HttpStatusCode: 500,
		}

		return nil, problem
	}

	return &rawRefreshToken, nil
}

func (t *TokenService) VerifyAccessToken(rawAccessToken string) (*types.AccessToken, error) {
	token, err := jwt.Parse(rawAccessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errorDescripion := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])

			problem := problems.Problem{
				Id:             "67ff9253-befd-495d-8cdd-225e09f46da9",
				Message:        "Failed to verify access token",
				Description:    &errorDescripion,
				HttpStatusCode: 401,
			}

			return nil, problem
		}

		return []byte(t.PrivateKey), nil
	})

	subject, err := token.Claims.GetSubject()

	if err != nil {
		return nil, err
	}

	issuedAt, err := token.Claims.GetIssuedAt()

	if err != nil {
		return nil, err
	}

	expirationTime, err := token.Claims.GetExpirationTime()

	if err != nil {
		return nil, err
	}

	issuer, err := token.Claims.GetIssuer()

	if err != nil {
		return nil, err
	}

	audience, err := token.Claims.GetAudience()

	if err != nil {
		return nil, err
	}

	accessToken := &types.AccessToken{
		Header: types.AccessTokenHeader{
			Algorithm: token.Method.Alg(),
		},
		Payload: types.AccessTokenPayload{
			Subject:   subject,
			IssuedAt:  issuedAt.Time,
			Issuer:    issuer,
			Audience:  audience[0], // HINT: Our audience is a single string, not an array @ Roman
			ExpiresAt: expirationTime.Time,
		},
	}

	return accessToken, nil
}

func (t *TokenService) VerifyRefreshToken(rawRefreshToken string) (*types.RefreshToken, error) {
	token, err := jwt.Parse(rawRefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errorDescripion := fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"])

			problem := problems.Problem{
				Id:             "471260b7-043c-4f3e-bdb6-e4f0bff6c47d",
				Message:        "Failed to verify refresh token",
				Description:    &errorDescripion,
				HttpStatusCode: 401,
			}

			return nil, problem
		}

		return []byte(t.PrivateKey), nil
	})

	subject, err := token.Claims.GetSubject()

	if err != nil {
		return nil, err
	}

	issuedAt, err := token.Claims.GetIssuedAt()

	if err != nil {
		return nil, err
	}

	expirationTime, err := token.Claims.GetExpirationTime()

	if err != nil {
		return nil, err
	}

	issuer, err := token.Claims.GetIssuer()

	if err != nil {
		return nil, err
	}

	audience, err := token.Claims.GetAudience()

	if err != nil {
		return nil, err
	}

	refreshToken := &types.RefreshToken{
		Header: types.RefreshTokenHeader{
			Algorithm: token.Method.Alg(),
		},
		Payload: types.RefreshTokenPayload{
			Subject:   subject,
			IssuedAt:  issuedAt.Time,
			Issuer:    issuer,
			Audience:  audience[0], // HINT: Our audience is a single string, not an array @ Roman
			ExpiresAt: expirationTime.Time,
		},
	}

	return refreshToken, nil
}
