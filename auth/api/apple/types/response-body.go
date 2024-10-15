package types

type ResponseBody struct {
	User         *User   `json:"user"`
	AccessToken  *string `json:"accessToken"`
	RefreshToken *string `json:"refreshToken"`
}
