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

// FindAll akan mengembalikan semua data Kavling
// @Summary      Menampilkan semua data Kavling
// @Description  Menampilkan semua data Kavling
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Params 		 ground_id query string false "ID Kavling"
// @Success      201  {object}  utils.SuccessResponseData{data=[]model.Kavling}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling [get]
func (pc *Kavling) FindAll(ctx *gin.Context) {
	idKavling := ctx.Query("ground_id")
	Kavlings, err := pc.repoKavling.FindAll(idKavling)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, Kavlings)
}

// FindByID akan mengembalikan data Kavling berdasarkan id
// @Summary      Menampilkan data Kavling berdasarkan id
// @Description  Menampilkan data Kavling berdasarkan id
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Kavling"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling/{id} [get]
func (pc *Kavling) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	Kavling, err := pc.repoKavling.FindByID(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, Kavling)
}

// Create akan membuat data Kavling baru
// @Summary      Membuat data Kavling baru
// @Description  Membuat data Kavling baru
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Param        Kavling     body    model.KavlingInput     true  "Data Kavling"
// @Param 		 Authorization header string true "Bearer token"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling [post]
func (pc *Kavling) Create(ctx *gin.Context) {
	Kavling := &model.KavlingInput{}
	if err := ctx.ShouldBindJSON(Kavling); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	kavling, err := pc.repoKavling.Create(Kavling)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 201, kavling)
}

// Update akan memperbarui data Kavling berdasarkan id
// @Summary      Memperbarui data Kavling berdasarkan id
// @Description  Memperbarui data Kavling berdasarkan id
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Kavling"
// @Param 		 Authorization header string true "Bearer token"
// @Param        Kavling     body    model.KavlingInput     true  "Data Kavling"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling/{id} [put]
func (pc *Kavling) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	Kavling := &model.KavlingInput{}
	if err := ctx.ShouldBindJSON(Kavling); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	if err := pc.repoKavling.Update(id, Kavling); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 201, Kavling)
}

// Delete akan menghapus data Kavling berdasarkan id
// @Summary      Menghapus data Kavling berdasarkan id
// @Description  Menghapus data Kavling berdasarkan id
// @Tags         Kavling
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Kavling"
// @Param 		 Authorization header string true "Bearer token"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Kavling}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /kavling/{id} [delete]
func (pc *Kavling) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := pc.repoKavling.Delete(id); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, "Kavling berhasil dihapus")
}
