package repository

import (
	"hacktiv8_final_project/entity"

	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{db}
}

func (r *CommentRepositoryImpl) FindAll() ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := r.db.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepositoryImpl) FindByID(id string, comment *entity.Comment) error {
	return r.db.First(comment, id).Error
}

func (r *CommentRepositoryImpl) Create(comment entity.Comment) (entity.Comment, error) {
	if err := r.db.Create(&comment).Error; err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}

func (r *CommentRepositoryImpl) Update(comment entity.Comment) error {
	return r.db.Save(&comment).Error
}

func (r *CommentRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Comment{}, id).Error
}
