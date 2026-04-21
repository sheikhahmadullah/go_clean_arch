package domain

type RefreshToken struct {
	ID     uint `gorm:"primaryKey"`
	Token  string
	UserID uint
}
 