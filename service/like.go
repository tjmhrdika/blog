package service

import (
	"blog/entity"
	"blog/repository"
)

type LikeService interface {
	CreateLike(blogID uint64, userID uint64) (entity.Like, error)
	DeleteLike(blogID uint64, userID uint64) (entity.Like, error)
}

type likeService struct {
	likeRepository repository.LikeRepository
}

func NewLikeService(lr repository.LikeRepository) LikeService {
	return &likeService{
		likeRepository: lr,
	}
}

func (ls *likeService) CreateLike(blogID uint64, userID uint64) (entity.Like, error) {
	var like = entity.Like{
		BlogID: blogID,
		UserID: userID,
	}
	if err := ls.likeRepository.CreateLike(like); err != nil {
		return entity.Like{}, err
	}
	return like, nil
}

func (ls *likeService) DeleteLike(blogID uint64, userID uint64) (entity.Like, error) {
	var like = entity.Like{
		BlogID: blogID,
		UserID: userID,
	}
	if err := ls.likeRepository.DeleteLike(like); err != nil {
		return entity.Like{}, err
	}
	return like, nil
}
