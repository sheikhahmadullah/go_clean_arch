package domain

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"-"` // Keep this! It protects your data on output.
}
