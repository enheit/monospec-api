package types

type RequestBody struct {
	IdentityToken string `json:"identityToken" validate:"required"`

	// NOTE: It seems like Apple does not include user's first name in the identity token, but we recive it in iOS app. So, I think it's safe to pass user's name along the identity token, so we can set his name @ Roman
	FirstName string `json:"firstName"`
}
