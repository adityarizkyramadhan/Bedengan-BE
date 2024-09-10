package routes

import (
	"github.com/adityarizkyramadhan/template-go-mvc/controller"
	"github.com/adityarizkyramadhan/template-go-mvc/middleware"
	"github.com/gin-gonic/gin"
)

type Kavling struct {
	ctrlKavling *controller.Kavling
}

func NewKavlingRoutes(ctrlKavling *controller.Kavling) *Kavling {
	return &Kavling{ctrlKavling}
}

// SetupRoutes will setup the routes for kavling
func (k *Kavling) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/kavling", k.ctrlKavling.FindAll)
	router.GET("/kavling/:id", k.ctrlKavling.FindByID)
	router.POST("/kavling", middleware.JWTMiddleware([]string{"admin"}), k.ctrlKavling.Create)
	router.PUT("/kavling/:id", middleware.JWTMiddleware([]string{"admin"}), k.ctrlKavling.Update)
	router.DELETE("/kavling/:id", middleware.JWTMiddleware([]string{"admin"}), k.ctrlKavling.Delete)
}
