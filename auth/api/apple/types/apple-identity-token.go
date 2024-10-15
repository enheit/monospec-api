package types

import "time"

type AppleIdentityToken struct {
	Header  AppleIdentityTokenHeader
	Payload AppleIdentityTokenPayload
}

type AppleIdentityTokenHeader struct {
	KeyId     string // HINT: As far as I understand this is an id of a public key provided by apple @ Roman
	Algorithm string
}

type AppleIdentityTokenPayload struct {
	Subject         string
	Email           string
	IsEmailVerified bool
	AuthenticatedAt time.Time // HINT: Last time when user was authenticated (info from ChatGPT) @ Roman
	IssuedAt        time.Time
}
