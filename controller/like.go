package controller

import (
	"blog/service"
	"blog/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LikeController interface {
	CreateLike(ctx *gin.Context)
	DeleteLike(ctx *gin.Context)
}

type likeController struct {
	likeService service.LikeService
}

func NewLikeController(ls service.LikeService) LikeController {
	return &likeController{
		likeService: ls,
	}
}

func (lc *likeController) CreateLike(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	blogID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	like, err := lc.likeService.CreateLike(blogID, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Like", err)
		ctx.JSON(http.StatusBadRequest, res)
	}
	res := utils.BuildResponseSuccess("Berhasil Menambahkan Like", like)
	ctx.JSON(http.StatusOK, res)
}

func (lc *likeController) DeleteLike(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	blogID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	like, err := lc.likeService.DeleteLike(blogID, userID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menghapus Like", err)
		ctx.JSON(http.StatusBadRequest, res)
	}
	res := utils.BuildResponseSuccess("Berhasil Menghapus Like", like)
	ctx.JSON(http.StatusOK, res)
}
