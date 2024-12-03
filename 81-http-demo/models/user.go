package models

import "fmt"

type User struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("invalid user name")
	}

	if u.Email == "" {
		return fmt.Errorf("invalid user email")
	}
	return nil
}
