package repository

import (
	"blog/entity"

	"gorm.io/gorm"
)

type LikeRepository interface {
	CreateLike(like entity.Like) error
	DeleteLike(like entity.Like) error
}

type likeRepository struct {
	DB *gorm.DB
}

func NewLikeRepository(db *gorm.DB) LikeRepository {
	return &likeRepository{
		DB: db,
	}
}

func (lr *likeRepository) CreateLike(like entity.Like) error {
	return lr.DB.Create(&like).Error
}

func (lr *likeRepository) DeleteLike(like entity.Like) error {
	return lr.DB.Delete(&like).Error
}
