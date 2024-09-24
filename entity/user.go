package entity

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"uniqueIndex" binding:"required,email"`
	Username  string    `json:"username" gorm:"uniqueIndex" binding:"required"`
	Password  string    `json:"password" binding:"required,min=6"`
	Age       uint      `json:"age" binding:"required,gt=8"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
