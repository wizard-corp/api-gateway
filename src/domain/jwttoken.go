package domain

type JwtTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type JwtToken struct {
	AccessTokenSecret      string
	AccessTokenExpiryHour  int
	RefreshTokenSecret     string
	RefreshTokenExpiryHour int
}
