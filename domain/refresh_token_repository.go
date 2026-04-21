package domain

type RefreshTokenRepository interface {
	SaveRefreshToken(token *RefreshToken) error
	FindByToken(token string) (*RefreshToken, error)
}
