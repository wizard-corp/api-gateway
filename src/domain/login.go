package domain

type Login struct {
	Email    string
	Password string
	JwtToken
}

type LoginRepository interface {
	GetUserByEmail(email string) (User, error)
}

func (l *Login) IsLoginValid() []string {
	var errors []string

	if l.Email == "" {
		errors = append(errors, IS_EMPTY)
	}

	if l.Password == "" {
		errors = append(errors, IS_EMPTY)
	}

	if IsValidEmail(l.Email) {
		errors = append(errors, INVALID_FORMAT)
	}

	return errors
}
