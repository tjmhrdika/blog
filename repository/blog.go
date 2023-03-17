package repository

import (
	"blog/entity"

	"gorm.io/gorm"
)

type BlogRepository interface {
	CreateBlog(blog entity.Blog) (entity.Blog, error)
	GetBlogByID(blogID uint64) (entity.Blog, error)
	GetUserIDByBlogID(blogID uint64) (uint64, error)
	UpdateBlog(blog entity.Blog) (entity.Blog, error)
	DeleteBlog(blogID uint64) error
}

type blogRepository struct {
	DB *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{
		DB: db,
	}
}

func (br *blogRepository) CreateBlog(blog entity.Blog) (entity.Blog, error) {
	if err := br.DB.Create(&blog).Error; err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}

func (br *blogRepository) GetBlogByID(blogID uint64) (entity.Blog, error) {
	var blog entity.Blog
	if err := br.DB.Preload("Likes").Preload("Comments").Where("id = ?", blogID).Take(&blog).Error; err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}

func (br *blogRepository) GetUserIDByBlogID(blogID uint64) (uint64, error) {
	var blog entity.Blog
	if err := br.DB.Where("id = ?", blogID).Take(&blog).Error; err != nil {
		return 0, err
	}
	return blog.UserID, nil
}

func (br *blogRepository) UpdateBlog(blog entity.Blog) (entity.Blog, error) {
	if err := br.DB.Updates(&blog).Error; err != nil {
		return entity.Blog{}, err
	}
	return blog, nil
}

func (br *blogRepository) DeleteBlog(blogID uint64) error {
	if err := br.DB.Delete(&entity.Blog{}, "id = ?", blogID).Error; err != nil {
		return err
	}
	return nil
}
