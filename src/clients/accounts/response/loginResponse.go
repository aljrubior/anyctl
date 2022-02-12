package response

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	RedirectUrl string `json:"redirectUrl"`
}
