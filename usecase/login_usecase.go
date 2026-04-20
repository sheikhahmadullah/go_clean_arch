package usecase

import (
	"errors"
	"golang-rest-api/domain"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	repo domain.UserRepository
}

func NewLoginUsecase(r domain.UserRepository) *LoginUsecase {
	return &LoginUsecase{r}
}

func (u *LoginUsecase) Login(req *domain.LoginRequest) (*domain.User, error) {
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	user, err := u.repo.FindByEmail(req.Email)

	if err != nil {
		return nil, errors.New("User not found!")
	}

	// debugging

	// fmt.Printf("Attempting login for: %s\n", req.Email)

	// fmt.Printf("Hash from DB: %s\n", user.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, errors.New("Invalid password")
	}

	return user, nil
}
