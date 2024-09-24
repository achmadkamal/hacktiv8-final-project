package repository

import (
	"hacktiv8_final_project/entity"
)

type CommentRepository interface {
	FindAll() ([]entity.Comment, error)
	FindByID(id string, comment *entity.Comment) error
	Create(comment entity.Comment) (entity.Comment, error)
	Update(comment entity.Comment) error
	Delete(id string) error
}
