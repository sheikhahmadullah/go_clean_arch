package usecase

import (
	"golang-rest-api/domain"

	"golang.org/x/crypto/bcrypt"
)

// this struct holds a dependency , instead of holding the database connection
// it holds an interface (domain.UserRepository)
type SignupUsecase struct {
	repo domain.UserRepository
}

// it relies on an interface rather than the actual gorm struct
// it doesn't care if we are using postgres , mongodb or something else
// it only knows that the repo has a Create() method

func NewSignupUsecase(r domain.UserRepository) *SignupUsecase {
	return &SignupUsecase{r}
}

func (u *SignupUsecase) Signup(user *domain.User) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	user.Password = string(hashed)

	return u.repo.Create(user)
}
