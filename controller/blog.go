package controller

import (
	"blog/dto"
	"blog/service"
	"blog/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController interface {
	CreateBlog(ctx *gin.Context)
	GetBlogByID(ctx *gin.Context)
	UpdateBlog(ctx *gin.Context)
	DeleteBlog(ctx *gin.Context)
}

type blogController struct {
	blogService service.BlogService
}

func NewBlogController(bs service.BlogService) BlogController {
	return &blogController{
		blogService: bs,
	}
}

func (bc *blogController) CreateBlog(ctx *gin.Context) {
	temp, success := ctx.Get("userID")
	if !success {
		res := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("gagal mendapat user_id"))
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	userID := temp.(uint64)

	var blogDTO dto.BlogCreate
	if err := ctx.ShouldBindJSON(&blogDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengambil Request", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	blog, err := bc.blogService.CreateBlog(userID, blogDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Blog", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Blog", blog)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) GetBlogByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := bc.blogService.GetBlogByID(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Blog", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) UpdateBlog(ctx *gin.Context) {
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

	if err := bc.blogService.VerifyBlog(blogID, userID); err != nil {
		res := utils.BuildResponseFailed("Gagal Verifikasi", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var blogDTO dto.BlogUpdate
	if err := ctx.ShouldBindJSON(&blogDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengambil Request", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	blog, err := bc.blogService.UpdateBlog(blogID, blogDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Blog", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengupdate Blog", blog)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) DeleteBlog(ctx *gin.Context) {
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

	if err := bc.blogService.VerifyBlog(blogID, userID); err != nil {
		res := utils.BuildResponseFailed("Gagal Verifikasi", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := bc.blogService.DeleteBlog(blogID); err != nil {
		res := utils.BuildResponseFailed("Gagal Menghapus Blog", err)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res := utils.BuildResponseSuccess("Berhasil Menghapus Blog", blogID)
	ctx.JSON(http.StatusOK, res)
}
