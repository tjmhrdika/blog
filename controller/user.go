package controller

import (
	"errors"
	"net/http"

	"blog/dto"
	"blog/service"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUserNama(ctx *gin.Context)
	UpdateUserPassword(ctx *gin.Context)
}

type userController struct {
	jwtService  service.JWTService
	userService service.UserService
}

func NewUserController(us service.UserService, jwt service.JWTService) UserController {
	return &userController{
		jwtService:  jwt,
		userService: us,
	}
}

func (uc *userController) RegisterUser(ctx *gin.Context) {
	var user dto.UserRegister
	if err := ctx.ShouldBindJSON(&user); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengambil Request", err)
		ctx.JSON(http.StatusBadRequest, res)
	}

	if checkUser, _ := uc.userService.CheckUser(user.Email); checkUser {
		res := utils.BuildResponseFailed("Email Sudah Terdaftar", errors.New("Failed"))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	result, err := uc.userService.RegisterUser(user)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan User", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) LoginUser(ctx *gin.Context) {
	var userLoginDTO dto.UserLogin
	err := ctx.ShouldBindJSON(&userLoginDTO)
	if err != nil {
		response := utils.BuildResponseFailed("Gagal Mengambil Request", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	err = uc.userService.Verify(userLoginDTO.Email, userLoginDTO.Password)
	if err != nil {
		response := utils.BuildResponseFailed("Gagal Login", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := uc.userService.GetUserByEmail(userLoginDTO.Email)
	if err != nil {
		response := utils.BuildResponseFailed("Gagal Login", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	token := uc.jwtService.GenerateToken(user.ID)
	response := utils.BuildResponseSuccess("Berhasil Login", token)
	ctx.JSON(http.StatusOK, response)
}

func (uc *userController) GetUser(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	result, err := uc.userService.GetUserByID(userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan User", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan User", result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	if err := uc.userService.DeleteUser(userID); err != nil {
		res := utils.BuildResponseFailed("Gagal Menghapus User", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menghapus User", userID)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) UpdateUserNama(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	var nama string
	err := ctx.ShouldBindJSON(&nama)
	if err != nil {
		response := utils.BuildResponseFailed("Gagal Mengambil Request", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err := uc.userService.UpdateUserNama(userID, nama); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Nama User", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengupdate Nama User", nama)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) UpdateUserPassword(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	var password dto.UserUpdatePassword
	err := ctx.ShouldBindJSON(&password)
	if err != nil {
		response := utils.BuildResponseFailed("Gagal Mengambil Request", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = uc.userService.VerifyPassword(userID, password.PasswordLama)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Password User", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	if err := uc.userService.UpdateUserPassword(userID, password.PasswordBaru); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Password User", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengupdate Password User", password.PasswordBaru)
	ctx.JSON(http.StatusOK, res)
}
