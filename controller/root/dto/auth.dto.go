package dto

type AuthDto struct {
	AccessToken string `json:"access_token"`
}

func NewAuthDto(accessToken string) *AuthDto {
	return &AuthDto{AccessToken: accessToken}
}
