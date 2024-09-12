package domain

type Signup struct {
	User
	JwtToken
}

type SignupRepository interface {
	GetUserByEmail(email string) (User, error)
	Create(*User) error
}

func (s *Signup) IsSignupValid() []string {
	errors := s.User.IsValidUser()
	return errors
}
