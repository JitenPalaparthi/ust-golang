package database

import (
	"demo/models"
	"errors"

	"gorm.io/gorm"
)

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Create(user *models.User) (*models.User, error) {

	if u.DB == nil {
		return nil, errors.New("invalid database connection")
	}
	u.DB.AutoMigrate(models.User{}) // this creates the table at first. Any updates are there in the model , they are taken care

	tr := u.DB.Create(user)
	if tr.Error != nil {
		return nil, tr.Error
	}
	return user, nil
}
