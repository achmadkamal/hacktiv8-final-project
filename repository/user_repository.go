package repository

import "hacktiv8_final_project/entity"

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	FindByID(id uint) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	FindByEmail(username string) (entity.User, error)
}
