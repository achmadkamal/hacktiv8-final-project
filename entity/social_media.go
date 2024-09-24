package entity

import "time"

type SocialMedia struct {
	Id             uint      `json:"id" gorm:"primaryKey"`
	Name           string    `json:"name" binding:"required"`
	SocialMediaUrl string    `json:"social_media_url" binding:"required"`
	UserId         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
