package repository

import (
	"blog/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user entity.User) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	GetUserByID(userID uint64) (entity.User, error)
	GetPasswordByID(userID uint64) (string, error)
	DeleteUser(userID uint64) error
	UpdateUserNama(userID uint64, nama string) (entity.User, error)
	UpdateUserPassword(userID uint64, password string) (entity.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) RegisterUser(user entity.User) (entity.User, error) {
	if err := ur.DB.Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByEmail(email string) (entity.User, error) {
	var user entity.User
	if err := ur.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetUserByID(userID uint64) (entity.User, error) {
	var user entity.User
	if err := ur.DB.Where("id = ?", userID).Preload("Blogs").Preload("Comments").Preload("Likes").Take(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ur *userRepository) GetPasswordByID(userID uint64) (string, error) {
	var user entity.User
	if err := ur.DB.Where("id = ?", userID).Take(&user).Error; err != nil {
		return "", err
	}
	return user.Password, nil
}

func (ur *userRepository) DeleteUser(userID uint64) error {
	if err := ur.DB.Delete(&entity.User{}, "id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUserNama(userID uint64, nama string) (entity.User, error) {
	var user entity.User
	user.ID = userID
	user.Nama = nama
	if err := ur.DB.Updates(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (ur *userRepository) UpdateUserPassword(userID uint64, password string) (entity.User, error) {
	var user entity.User
	user.ID = userID
	user.Password = password
	if err := ur.DB.Updates(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
