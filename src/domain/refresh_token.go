package domain

type RefreshToken struct {
	RefreshToken string
	JwtToken
}

type RefreshTokenRepository interface {
	GetUserByID(id string) (User, error)
}

func (rtc *RefreshToken) IsRefreshTokenValid() []string {
	var errors []string

	if rtc.RefreshToken == "" {
		errors = append(errors, IS_EMPTY)
	}

	return errors
}
