package types

import "time"

type RefreshToken struct {
	Header  RefreshTokenHeader
	Payload RefreshTokenPayload
}

type RefreshTokenHeader struct {
	Algorithm string
}

type RefreshTokenPayload struct {
	Subject   string
	IssuedAt  time.Time
	Issuer    string
	Audience  string
	ExpiresAt time.Time
}
