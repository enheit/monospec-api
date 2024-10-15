package types

import "time"

type AccessToken struct {
	Header  AccessTokenHeader
	Payload AccessTokenPayload
}

type AccessTokenHeader struct {
	Algorithm string
}

type AccessTokenPayload struct {
	Issuer    string
	Audience  string
	Subject   string
	IssuedAt  time.Time
	ExpiresAt time.Time
}
