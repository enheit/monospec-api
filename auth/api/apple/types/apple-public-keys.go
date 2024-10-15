package types

type ApplePublicKeys struct {
	Keys []ApplePublicKey `json:"keys"`
}

type ApplePublicKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
}
