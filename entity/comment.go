package entity

import "time"

type Comment struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	UserId    uint      `json:"user_id"`
	PhotoId   uint      `json:"photo_id"`
	Message   string    `json:"message" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
