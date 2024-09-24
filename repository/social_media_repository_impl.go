package repository

import (
	"hacktiv8_final_project/entity"

	"gorm.io/gorm"
)

type SocialMediaRepositoryImpl struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryImpl{db}
}

func (r *SocialMediaRepositoryImpl) FindAll() ([]entity.SocialMedia, error) {
	var socialMedias []entity.SocialMedia
	r.db.Find(&socialMedias)
	return socialMedias, nil
}

func (r *SocialMediaRepositoryImpl) FindByID(id string, socialMedia *entity.SocialMedia) error {
	return r.db.First(socialMedia, id).Error
}

func (r *SocialMediaRepositoryImpl) Create(socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	r.db.Create(&socialMedia)
	return socialMedia, nil
}

func (r *SocialMediaRepositoryImpl) Update(socialMedia entity.SocialMedia) error {
	return r.db.Save(&socialMedia).Error
}

func (r *SocialMediaRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.SocialMedia{}, id).Error
}
