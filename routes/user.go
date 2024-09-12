package routes

import (
	"github.com/adityarizkyramadhan/template-go-mvc/controller"
	"github.com/adityarizkyramadhan/template-go-mvc/middleware"
	"github.com/gin-gonic/gin"
)

type User struct {
	ctrlUser *controller.User
}

func NewUserRoutes(ctrlUser *controller.User) *User {
	return &User{ctrlUser}
}

// SetupRoutes will setup the routes for user
func (u *User) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/user/register", u.ctrlUser.Register)
	router.POST("/user/login", u.ctrlUser.Login)
	router.GET("/user/logout", middleware.JWTMiddleware([]string{"admin", "user"}), u.ctrlUser.Logout)
	router.PUT("/user", middleware.JWTMiddleware([]string{"admin", "user"}), u.ctrlUser.Update)
	router.GET("/user", middleware.JWTMiddleware([]string{"admin", "user"}), u.ctrlUser.FindOne)
}
