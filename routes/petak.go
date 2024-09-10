package routes

import (
	"github.com/adityarizkyramadhan/template-go-mvc/controller"
	"github.com/adityarizkyramadhan/template-go-mvc/middleware"
	"github.com/gin-gonic/gin"
)

type Petak struct {
	ctrlPetak *controller.Petak
}

func NewPetakRoutes(ctrlPetak *controller.Petak) *Petak {
	return &Petak{ctrlPetak}
}

// SetupRoutes will setup the routes for petak
func (p *Petak) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/petak", p.ctrlPetak.FindAll)
	router.GET("/petak/:id", p.ctrlPetak.FindByID)
	router.POST("/petak", middleware.JWTMiddleware([]string{"admin"}), p.ctrlPetak.Create)
	router.PUT("/petak/:id", middleware.JWTMiddleware([]string{"admin"}), p.ctrlPetak.Update)
	router.DELETE("/petak/:id", middleware.JWTMiddleware([]string{"admin"}), p.ctrlPetak.Delete)
}
