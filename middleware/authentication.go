package middleware

import (
	"errors"
	"net/http"
	"strings"

	"blog/service"
	"blog/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("token tidak ditemukan"))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			response := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("token tidak valid"))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed("Gagal Memproses Request", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		if !token.Valid {
			response := utils.BuildResponseFailed("Gagal Memproses Request", errors.New("token tidak valid"))
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID, err := jwtService.GetUserIDByToken(authHeader)
		if err != nil {
			response := utils.BuildResponseFailed("Gagal Memproses Request", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		ctx.Set("userID", userID)
		ctx.Next()
	}
}
