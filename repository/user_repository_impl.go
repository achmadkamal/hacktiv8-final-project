package repository

import (
	"hacktiv8_final_project/entity"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) Create(user entity.User) (entity.User, error) {
	r.db.Create(&user)
	return user, nil
}

func (r *UserRepositoryImpl) FindByID(id uint) (entity.User, error) {
	var user entity.User
	r.db.First(&user, id)
	return user, nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	var user entity.User
	r.db.Where("username = ?", username).First(&user)
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	r.db.Where("email = ?", email).First(&user)
	return user, nil
}
