package controller

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

type Kavling struct {
	repoKavling *repositories.Kavling
}

func NewKavlingController(repoKavling *repositories.Kavling) *Kavling {
	return &Kavling{repoKavling}
}

// FindAll akan mengembalikan semua data kavling
// @Summary      Menampilkan semua data kavling
// @Description  Menampilkan semua data kavling
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling [get]
func (kc *Kavling) FindAll(ctx *gin.Context) {
	kavlings, err := kc.repoKavling.FindAll()
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, kavlings)
}

// FindByID akan mengembalikan data kavling berdasarkan id
// @Summary      Menampilkan data kavling berdasarkan id
// @Description  Menampilkan data kavling berdasarkan id
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Kavling"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling/{id} [get]
func (kc *Kavling) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	kavling, err := kc.repoKavling.FindByID(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, kavling)
}

// Create akan membuat data kavling baru
// @Summary      Membuat data kavling baru
// @Description  Membuat data kavling baru
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Param        kavling     body    model.KavlingInput     true  "Data Kavling"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling [post]
func (kc *Kavling) Create(ctx *gin.Context) {
	kavling := &model.KavlingInput{}
	if err := ctx.ShouldBindJSON(kavling); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	if err := kc.repoKavling.Create(kavling); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 201, kavling)
}

// Update akan memperbarui data kavling berdasarkan id
// @Summary      Memperbarui data kavling berdasarkan id
// @Description  Memperbarui data kavling berdasarkan id
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Kavling"
// @Param        kavling     body    model.KavlingInput     true  "Data Kavling"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling/{id} [put]
func (kc *Kavling) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	kavling := &model.KavlingInput{}
	if err := ctx.ShouldBindJSON(kavling); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	if err := kc.repoKavling.Update(id, kavling); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 201, kavling)
}

// Delete akan menghapus data kavling berdasarkan id
// @Summary      Menghapus data kavling berdasarkan id
// @Description  Menghapus data kavling berdasarkan id
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Kavling"
// @Param        kavling     body    model.KavlingInput     true  "Data Kavling"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling/{id} [delete]
func (kc *Kavling) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := kc.repoKavling.Delete(id); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, "Kavling berhasil dihapus")
}
