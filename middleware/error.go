package middleware

import (
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

// ErrorHandler adalah middleware untuk menangkap dan menangani error
func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Jalankan permintaan dan tangkap semua error
		ctx.Next()
		// Tangkap error jika ada
		err := ctx.Errors.Last()
		if err != nil {
			// Tentukan status code dan pesan error yang sesuai
			errParse := utils.ParseError(err.Error())
			utils.ErrorResponse(ctx, errParse.StatusCode, errParse.Message)
		}
	}
}
