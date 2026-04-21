package usecase

import (
	"golang-rest-api/domain"
	"golang-rest-api/internal/tokenutil"
)

type RefreshUsecase struct {
	repo domain.RefreshTokenRepository
}

func NewRefreshUsecase(r domain.RefreshTokenRepository) *RefreshUsecase {
	return &RefreshUsecase{r}
}

// refresh will give us new token by taking old token

func (u *RefreshUsecase) Refresh(oldToken string) (string, error) {
	rt, err := u.repo.FindByToken(oldToken)

	if err != nil {
		return "", err
	}

	return tokenutil.GenerateToken(rt.UserID)
}
