package entity

import (
	"blog/utils"

	"gorm.io/gorm"
)

type User struct {
	ID       uint64    `json:"id" gorm:"primaryKey"`
	Nama     string    `json:"nama"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Blogs    []Blog    `json:"blogs,omitempty"`
	Likes    []Like    `json:"likes,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	var err error
	if u.Password != "" {
		u.Password, err = utils.HashPassword(u.Password)
	}
	if err != nil {
		return err
	}
	return nil
}
