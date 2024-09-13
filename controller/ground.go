package controller

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

type Ground struct {
	repoGround *repositories.Ground
}

func NewGroundController(repoGround *repositories.Ground) *Ground {
	return &Ground{repoGround}
}

// FindAll akan mengembalikan semua data Ground
// @Summary      Menampilkan semua data Ground
// @Description  Menampilkan semua data Ground
// @Tags         Ground
// @Accept       json
// @Produce      json
// @Success      201  {object}  utils.SuccessResponseData{data=[]model.Ground}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /ground [get]
func (kc *Ground) FindAll(ctx *gin.Context) {
	Grounds, err := kc.repoGround.FindAll()
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, Grounds)
}

// FindByID akan mengembalikan data Ground berdasarkan id
// @Summary      Menampilkan data Ground berdasarkan id
// @Description  Menampilkan data Ground berdasarkan id
// @Tags         Ground
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Ground"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Ground}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /ground/{id} [get]
func (kc *Ground) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	Ground, err := kc.repoGround.FindByID(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, Ground)
}

// Create akan membuat data Ground baru
// @Summary      Membuat data Ground baru
// @Description  Membuat data Ground baru
// @Tags         Ground
// @Accept       multipart/form-data
// @Produce      json
// @Param        name           formData string true "Nama Ground"
// @Param        image          formData file   true "Gambar Ground"
// @Param 		   Authorization header string true "Bearer token"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Ground}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /ground [post]
func (kc *Ground) Create(ctx *gin.Context) {
	Ground := &model.GroundInput{}
	if err := ctx.ShouldBind(Ground); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	groundData, err := kc.repoGround.Create(Ground)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 201, groundData)
}

// Update akan memperbarui data Ground berdasarkan id
// @Summary      Memperbarui data Ground berdasarkan id
// @Description  Memperbarui data Ground berdasarkan id
// @Tags         Ground
// @Accept       multipart/form-data
// @Produce      json
// @Param        id     path    string     true  "ID Ground"
// @Param        name   formData string     true  "Nama Ground"
// @Param        image  formData file       true  "Gambar Ground"
// @Param 		 Authorization header string true "Bearer token"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Ground}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /ground/{id} [put]
func (kc *Ground) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	Ground := &model.GroundInput{}
	if err := ctx.ShouldBind(Ground); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	GroundData, err := kc.repoGround.Update(id, Ground)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 201, GroundData)
}

// Delete akan menghapus data Ground berdasarkan id
// @Summary      Menghapus data Ground berdasarkan id
// @Description  Menghapus data Ground berdasarkan id
// @Tags         Ground
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Ground"
// @Param 		 Authorization header string true "Bearer token"
// @Param        Ground     body    model.GroundInput     true  "Data Ground"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Ground}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /ground/{id} [delete]
func (kc *Ground) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := kc.repoGround.Delete(id); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, "Ground berhasil dihapus")
}
