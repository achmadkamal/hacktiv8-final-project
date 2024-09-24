package repository

import "hacktiv8_final_project/entity"

type SocialMediaRepository interface {
	FindAll() ([]entity.SocialMedia, error)
	FindByID(id string, socialMedia *entity.SocialMedia) error
	Create(socialMedia entity.SocialMedia) (entity.SocialMedia, error)
	Update(socialMedia entity.SocialMedia) error
	Delete(id string) error
}
