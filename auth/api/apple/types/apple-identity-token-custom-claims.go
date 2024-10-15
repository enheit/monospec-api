package types

import "github.com/golang-jwt/jwt/v5"

// HINT: Those claims were taked from the parsed Apple identity token @ Roman
// NOTE: I'm not sure if apple can return other claims or not, this is the only one I found @ Roman
type AppleIdentityTokenCustomClaims struct {
	CHash          string `json:"c_hash"`
	Email          string `json:"email"`
	EmailVerified  bool   `json:"email_verified"`
	AuthTime       int64  `json:"auth_time"`
	NonceSupported bool   `json:"nonce_supported"`
	RealUserStatus int64  `json:"real_user_status"`
	jwt.RegisteredClaims
}
