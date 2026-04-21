package repository

import (
	"golang-rest-api/domain"

	"gorm.io/gorm"
)

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) domain.RefreshTokenRepository {
	return &refreshTokenRepository{db}
}

func (r *refreshTokenRepository) SaveRefreshToken(token *domain.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *refreshTokenRepository) FindByToken(token string) (*domain.RefreshToken, error) {
	var refreshToken domain.RefreshToken

	err := r.db.Where("token = ?", token).First(&refreshToken).Error

	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}
