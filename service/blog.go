package service

import (
	"blog/dto"
	"blog/entity"
	"blog/repository"
	"errors"

	"github.com/jinzhu/copier"
)

type BlogService interface {
	CreateBlog(userID uint64, blogDTO dto.BlogCreate) (entity.Blog, error)
	GetBlogByID(userID uint64) (entity.Blog, error)
	VerifyBlog(blogID uint64, userID uint64) error
	UpdateBlog(blogID uint64, blogDTO dto.BlogUpdate) error
	DeleteBlog(blogID uint64) error
}

type blogService struct {
	blogRepository repository.BlogRepository
}

func NewBlogService(br repository.BlogRepository) BlogService {
	return &blogService{
		blogRepository: br,
	}
}

func (bs *blogService) CreateBlog(userID uint64, blogDTO dto.BlogCreate) (entity.Blog, error) {
	blog := entity.Blog{}
	if err := copier.Copy(&blog, &blogDTO); err != nil {
		return entity.Blog{}, err
	}
	blog.UserID = userID
	return bs.blogRepository.CreateBlog(blog)
}

func (bs *blogService) GetBlogByID(userID uint64) (entity.Blog, error) {
	return bs.blogRepository.GetBlogByID(userID)
}

func (bs *blogService) VerifyBlog(blogID uint64, userID uint64) error {
	res, err := bs.blogRepository.GetUserIDByBlogID(blogID)
	if err != nil {
		return err
	}
	if userID != res {
		return errors.New("user bukan pemilik blog")
	}
	return nil
}

func (bs *blogService) UpdateBlog(blogID uint64, blogDTO dto.BlogUpdate) error {
	blog := entity.Blog{}
	if err := copier.Copy(&blog, &blogDTO); err != nil {
		return err
	}
	blog.ID = blogID
	return bs.blogRepository.UpdateBlog(blog)
}

func (bs *blogService) DeleteBlog(blogID uint64) error {
	return bs.blogRepository.DeleteBlog(blogID)
}
