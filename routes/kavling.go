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

// SetupRoutes will setup the routes for Kavling
func (p *Kavling) SetupRoutes(router *gin.RouterGroup) {
	router.GET("/kavling", p.ctrlKavling.FindAll)
	router.GET("/kavling/:id", p.ctrlKavling.FindByID)
	router.POST("/kavling", middleware.JWTMiddleware([]string{"admin"}), p.ctrlKavling.Create)
	router.PUT("/kavling/:id", middleware.JWTMiddleware([]string{"admin"}), p.ctrlKavling.Update)
	router.DELETE("/kavling/:id", middleware.JWTMiddleware([]string{"admin"}), p.ctrlKavling.Delete)
}
