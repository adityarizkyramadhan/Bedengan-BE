package routes

import (
	"github.com/adityarizkyramadhan/template-go-mvc/controller"
	"github.com/adityarizkyramadhan/template-go-mvc/middleware"
	"github.com/gin-gonic/gin"
)

type Perlengkapan struct {
	ctrlPerlengkapan *controller.Perlengkapan
}

func NewPerlengkapanRoutes(ctrlPerlengkapan *controller.Perlengkapan) *Perlengkapan {
	return &Perlengkapan{ctrlPerlengkapan}
}

// SetupRoutes will setup the routes for perlengkapan
func (p *Perlengkapan) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/perlengkapan", p.ctrlPerlengkapan.FindAll)
	router.GET("/perlengkapan/:id", p.ctrlPerlengkapan.FindByID)
	router.POST("/perlengkapan", middleware.JWTMiddleware([]string{"admin"}), p.ctrlPerlengkapan.Create)
	router.PUT("/perlengkapan/:id", middleware.JWTMiddleware([]string{"admin"}), p.ctrlPerlengkapan.Update)
	router.DELETE("/perlengkapan/:id", middleware.JWTMiddleware([]string{"admin"}), p.ctrlPerlengkapan.Delete)
}
