package domain

type Task struct {
	Id     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	UserId uint   `json:"user_id"`
}
