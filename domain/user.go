package domain

type User struct {
	Id    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email"`

	Password string `json:"-"`
}
