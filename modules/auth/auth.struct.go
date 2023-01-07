package auth

type LoginBody struct {
	Username string
	Password string
}

type TokenResponse struct {
	AccessToken string `json:"accessToken"`
}
