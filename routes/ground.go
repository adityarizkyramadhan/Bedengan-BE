package routes

import (
	"github.com/adityarizkyramadhan/template-go-mvc/controller"
	"github.com/adityarizkyramadhan/template-go-mvc/middleware"
	"github.com/gin-gonic/gin"
)

type Ground struct {
	ctrlGround *controller.Ground
}

func NewGroundRoutes(ctrlGround *controller.Ground) *Ground {
	return &Ground{ctrlGround}
}

// SetupRoutes will setup the routes for Ground
func (k *Ground) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/ground", k.ctrlGround.FindAll)
	router.GET("/ground/:id", k.ctrlGround.FindByID)
	router.POST("/ground", middleware.JWTMiddleware([]string{"admin"}), k.ctrlGround.Create)
	router.PUT("/ground/:id", middleware.JWTMiddleware([]string{"admin"}), k.ctrlGround.Update)
	router.DELETE("/ground/:id", middleware.JWTMiddleware([]string{"admin"}), k.ctrlGround.Delete)
}
