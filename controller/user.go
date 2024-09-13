package controller

import (
	"net/http"

	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	repoUser repositories.UserInterface
}

// NewUserController will create an object that represent the User.Article interface
func NewUserController(repoUser repositories.UserInterface) *User {
	return &User{repoUser}
}

// Register will create a new user
// @Summary      Register new user
// @Description  Register new user
// @Tags         User
// @Accept       multipart/form-data
// @Produce      json
// @Param        email           formData string true "Email address"
// @Param        name            formData string true "Full name"
// @Param        password        formData string true "Password"
// @Param        confirm_password formData string true "Confirm password"
// @Success      201  {object}  utils.SuccessResponseData{data=model.User}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user/register [post]
func (u *User) Register(ctx *gin.Context) {
	user := &model.UserCreate{}
	if err := ctx.ShouldBind(user); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	newUser, err := u.repoUser.Create(user)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, http.StatusCreated, newUser)
}

// Login will login user
// @Summary      Login user
// @Description  Login user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 request  body  model.UserLogin true "User data"
// @Success      200  {object}  utils.SuccessResponseData{data=string}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user/login [post]
func (u *User) Login(ctx *gin.Context) {
	var user model.UserLogin
	if err := ctx.ShouldBindJSON(&user); err != nil {
		_ = ctx.Error(utils.NewError(utils.ErrValidation, "email atau password tidak valid"))
		ctx.Next()
		return
	}

	userData, err := u.repoUser.Login(user.Email, user.Password)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	token, err := utils.GenerateToken(userData.ID, userData.Email, userData.Role)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, token)
}

// Update will update user data
// @Summary      Update user data
// @Description  Update user data
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer token"
// @Param 		 request  body  model.UserUpdate true "User data"
// @Success      200  {object}  utils.SuccessResponseData{data=model.User}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user [put]
func (u *User) Update(ctx *gin.Context) {
	user := &model.UserUpdate{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	// Ambil uuid yang disimpan di context dari middleware JWT
	claimsData := ctx.MustGet("id").(string)

	userData, err := u.repoUser.Update(claimsData, user)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, userData)
}

// FindOne will find user by id
// @Summary      Find user by id
// @Description  Find user by id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer token"
// @Success      200  {object}  utils.SuccessResponseData{data=model.User}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user [get]
func (u *User) FindOne(ctx *gin.Context) {
	id := ctx.MustGet("id").(string)
	user, err := u.repoUser.FindOne(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, http.StatusOK, user)
}

// Logout will logout user
// @Summary      Logout user
// @Description  Logout user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param 		 Authorization header string true "Bearer token"
// @Success      200  {object}  utils.SuccessResponseData{data=string}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /user/logout [get]
func (u *User) Logout(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		_ = ctx.Error(utils.NewError(utils.ErrUnauthorized, "Token is required"))
		ctx.Next()
		return
	}

	durationToken, err := utils.GetExpiredToken(token)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	if err := u.repoUser.Logout(token, durationToken); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, http.StatusOK, "Logout success")
}
