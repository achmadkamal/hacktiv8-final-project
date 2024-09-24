package entity

import "time"

type Photo struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" binding:"required"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url" binding:"required"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
