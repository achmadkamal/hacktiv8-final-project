package database

import (
	"fmt"
	"hacktiv8_final_project/entity"
	"hacktiv8_final_project/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "postgres"
	PORT     = "5432"
	DATABASE = "hacktiv8_final_project"
)

var (
	db  *gorm.DB
	err error
)

func NewDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		HOST,
		USER,
		PASSWORD,
		DATABASE,
		PORT,
	)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	helper.PanicIfErr(err)

	err = db.AutoMigrate(&entity.User{}, &entity.Comment{}, &entity.Photo{}, &entity.SocialMedia{})
	helper.PanicIfErr(err)

	return db
}
