package routes

import (
	"github.com/adityarizkyramadhan/template-go-mvc/controller"
	"github.com/adityarizkyramadhan/template-go-mvc/middleware"
	"github.com/gin-gonic/gin"
)

type SubGround struct {
	ctrlSubGround *controller.SubGroundController
}

func NewSubGroundRoutes(ctrlSubGround *controller.SubGroundController) *SubGround {
	return &SubGround{ctrlSubGround}
}

// SetupRoutes will setup the routes for SubGround
func (p *SubGround) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/sub-ground", middleware.JWTMiddleware([]string{"admin"}), p.ctrlSubGround.Create)
	router.GET("/sub-ground", p.ctrlSubGround.FindAll)
	router.GET("/sub-ground/:id", p.ctrlSubGround.FindByID)
	router.PUT("/sub-ground/:id", middleware.JWTMiddleware([]string{"admin"}), p.ctrlSubGround.Update)
	router.DELETE("/sub-ground/:id", middleware.JWTMiddleware([]string{"admin"}), p.ctrlSubGround.Delete)
}
