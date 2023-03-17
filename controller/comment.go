package controller

import (
	"blog/service"
	"blog/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(cs service.CommentService) CommentController {
	return &commentController{
		commentService: cs,
	}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
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

	var text string
	if err := ctx.ShouldBindJSON(&text); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengambil Request", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	comment, err := cc.commentService.CreateComment(text, userID, blogID)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Comment", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Comment", comment)
	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) UpdateComment(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	commentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := cc.commentService.VerifyComment(commentID, userID); err != nil {
		res := utils.BuildResponseFailed("Gagal Verifikasi", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var text string
	if err := ctx.ShouldBindJSON(&text); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengambil Request", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	comment, err := cc.commentService.UpdateComment(commentID, text)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Comment", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengupdate Comment", comment)
	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) DeleteComment(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	commentID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := cc.commentService.VerifyComment(commentID, userID); err != nil {
		res := utils.BuildResponseFailed("Gagal Verifikasi", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := cc.commentService.DeleteComment(commentID); err != nil {
		res := utils.BuildResponseFailed("Gagal Menghapus Comment", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menghapus Comment", commentID)
	ctx.JSON(http.StatusOK, res)
}
