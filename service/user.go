package service

import (
	"blog/dto"
	"blog/entity"
	"blog/repository"
	"blog/utils"
	"errors"

	"github.com/jinzhu/copier"
)

type UserService interface {
	CheckUser(email string) (bool, error)
	RegisterUser(userDTO dto.UserRegister) (entity.User, error)
	Verify(email string, password string) error
	GetUserByEmail(email string) (entity.User, error)
	GetUserByID(userID uint64) (entity.User, error)
	DeleteUser(userID uint64) error
	UpdateUserNama(userID uint64, nama string) error
	UpdateUserPassword(userID uint64, password string) error
	VerifyPassword(userID uint64, password string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) CheckUser(email string) (bool, error) {
	res, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return false, err
	}

	if res.Email == "" {
		return false, nil
	}
	return true, nil
}

func (us *userService) RegisterUser(userDTO dto.UserRegister) (entity.User, error) {
	user := entity.User{}
	err := copier.Copy(&user, &userDTO)
	if err != nil {
		return user, err
	}
	return us.userRepository.RegisterUser(user)
}

func (us *userService) Verify(email string, password string) error {
	res, err := us.userRepository.GetUserByEmail(email)
	if err != nil {
		return err
	}
	CheckPassword, err := utils.CheckPassword(res.Password, []byte(password))
	if err != nil {
		return err
	}

	if !CheckPassword {
		return errors.New("password tidak sesuai")
	}
	return nil
}

func (us *userService) GetUserByEmail(email string) (entity.User, error) {
	return us.userRepository.GetUserByEmail(email)
}

func (us *userService) GetUserByID(userID uint64) (entity.User, error) {
	return us.userRepository.GetUserByID(userID)
}

func (us *userService) DeleteUser(userID uint64) error {
	return us.userRepository.DeleteUser(userID)
}

func (us *userService) UpdateUserNama(userID uint64, nama string) error {
	return us.userRepository.UpdateUserNama(userID, nama)
}

func (us *userService) UpdateUserPassword(userID uint64, password string) error {
	return us.userRepository.UpdateUserPassword(userID, password)
}

func (us *userService) VerifyPassword(userID uint64, password string) error {
	res, err := us.userRepository.GetPasswordByID(userID)
	if err != nil {
		return err
	}
	CheckPassword, err := utils.CheckPassword(res, []byte(password))
	if err != nil {
		return err
	}

	if !CheckPassword {
		return errors.New("password tidak sesuai")
	}
	return nil
}
