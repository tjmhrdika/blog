package repository

import (
	"blog/entity"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment entity.Comment) (entity.Comment, error)
	GetUserIDByCommentID(commentID uint64) (uint64, error)
	UpdateComment(comment entity.Comment) error
	DeleteComment(commentID uint64) error
}

type commentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

func (cr *commentRepository) CreateComment(comment entity.Comment) (entity.Comment, error) {
	if err := cr.DB.Create(&comment).Error; err != nil {
		return entity.Comment{}, err
	}
	return comment, nil
}

func (cr *commentRepository) GetUserIDByCommentID(commentID uint64) (uint64, error) {
	var comment entity.Comment
	if err := cr.DB.Where("id = ?", commentID).Take(&comment).Error; err != nil {
		return 0, err
	}
	return comment.UserID, nil
}

func (cr *commentRepository) UpdateComment(comment entity.Comment) error {
	if err := cr.DB.Updates(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *commentRepository) DeleteComment(commentID uint64) error {
	if err := cr.DB.Delete(&entity.Comment{}, "id = ?", commentID).Error; err != nil {
		return err
	}
	return nil
}
