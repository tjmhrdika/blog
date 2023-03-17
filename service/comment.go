package service

import (
	"blog/entity"
	"blog/repository"
	"errors"
)

type CommentService interface {
	CreateComment(text string, userID uint64, blogID uint64) (entity.Comment, error)
	VerifyComment(commentID uint64, userID uint64) error
	UpdateComment(commentID uint64, text string) (entity.Comment, error)
	DeleteComment(commentID uint64) error
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(cr repository.CommentRepository) CommentService {
	return &commentService{
		commentRepository: cr,
	}
}

func (cs *commentService) CreateComment(text string, userID uint64, blogID uint64) (entity.Comment, error) {
	var comment = entity.Comment{
		Text:   text,
		UserID: userID,
		BlogID: blogID,
	}
	return cs.commentRepository.CreateComment(comment)

}

func (cs *commentService) VerifyComment(commentID uint64, userID uint64) error {
	res, err := cs.commentRepository.GetUserIDByCommentID(commentID)
	if err != nil {
		return err
	}
	if userID != res {
		return errors.New("user bukan pemilik comment")
	}
	return nil
}

func (cs *commentService) UpdateComment(commentID uint64, text string) (entity.Comment, error) {
	var comment = entity.Comment{
		ID:   commentID,
		Text: text,
	}
	return cs.commentRepository.UpdateComment(comment)
}

func (cs *commentService) DeleteComment(commentID uint64) error {
	return cs.commentRepository.DeleteComment(commentID)
}
