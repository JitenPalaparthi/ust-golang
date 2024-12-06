package database

import (
	"demo/models"
	"errors"
	"fmt"

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
	u.DB.AutoMigrate(models.User{}) // this creates the table at first. Any updates are there in the model , they are updated in the table

	tr := u.DB.Create(user)
	if tr.Error != nil {
		return nil, tr.Error
	}
	return user, nil
}

func (u *UserDB) GetByID(id uint) (*models.User, error) {
	if u.DB == nil {
		return nil, errors.New("invalid database connection")
	}
	user := new(models.User)
	tx := u.DB.First(user, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func (u *UserDB) DeleteByID(id uint) (uint, error) {
	if u.DB == nil {
		return 0, errors.New("invalid database connection")
	}

	tx := u.DB.Delete(models.User{}, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("no record to delete based on id " + fmt.Sprint(id))
	}
	return uint(tx.RowsAffected), nil
}

func (u *UserDB) UpdateByID(user *models.User) (*models.User, error) {
	if u.DB == nil {
		return nil, errors.New("invalid database connection")
	}

	fmt.Println("---->", user)
	tx := u.DB.Updates(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("no record to udpate based on id " + fmt.Sprint(user.Id))
	}
	return user, nil
}
