package repository

import "hacktiv8_final_project/entity"

type PhotoRepository interface {
	FindAll() ([]entity.Photo, error)
	FindByID(id string, photo *entity.Photo) error
	Create(photo entity.Photo) (entity.Photo, error)
	Update(photo entity.Photo) error
	Delete(id string) error
}
