package repository

import (
	"hacktiv8_final_project/entity"

	"gorm.io/gorm"
)

type PhotoRepositoryImpl struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &PhotoRepositoryImpl{db}
}

func (r *PhotoRepositoryImpl) FindAll() ([]entity.Photo, error) {
	var photos []entity.Photo
	r.db.Find(&photos)
	return photos, nil
}

func (r *PhotoRepositoryImpl) FindByID(id string, photo *entity.Photo) error {
	return r.db.First(photo, id).Error
}

func (r *PhotoRepositoryImpl) Create(photo entity.Photo) (entity.Photo, error) {
	r.db.Create(&photo)
	return photo, nil
}

func (r *PhotoRepositoryImpl) Update(photo entity.Photo) error {
	return r.db.Save(&photo).Error
}

func (r *PhotoRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Photo{}, id).Error
}
