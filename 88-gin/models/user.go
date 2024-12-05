package models

import "fmt"

type User struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Contact      string `json:"contact"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("invalid user name")
	}
	if u.Email == "" {
		return fmt.Errorf("invalid email")
	}
	return nil
}
