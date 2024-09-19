package controller

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

type Perlengkapan struct {
	repoPerlengkapan *repositories.Perlengkapan
}

func NewPerlengkapanController(repoPerlengkapan *repositories.Perlengkapan) *Perlengkapan {
	return &Perlengkapan{repoPerlengkapan}
}

// FindAll akan mengembalikan semua data perlengkapan
// @Summary      Menampilkan semua data perlengkapan
// @Description  Menampilkan semua data perlengkapan
// @Tags         Perlengkapan
// @Accept       json
// @Produce      json
// @Success      201  {object}  utils.SuccessResponseData{data=[]model.Perlengkapan}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /perlengkapan [get]
func (pc *Perlengkapan) FindAll(ctx *gin.Context) {
	perlengkapans, err := pc.repoPerlengkapan.FindAll()
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, perlengkapans)
}

// FindByID akan mengembalikan data perlengkapan berdasarkan id
// @Summary      Menampilkan data perlengkapan berdasarkan id
// @Description  Menampilkan data perlengkapan berdasarkan id
// @Tags         Perlengkapan
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Perlengkapan"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Perlengkapan}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /perlengkapan/{id} [get]
func (pc *Perlengkapan) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	perlengkapan, err := pc.repoPerlengkapan.FindByID(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, perlengkapan)
}

// Create akan membuat data perlengkapan baru
// @Summary      Membuat data perlengkapan baru
// @Description  Membuat data perlengkapan baru
// @Tags         Perlengkapan
// @Accept       multipart/form-data
// @Produce      json
// @Param 		 Authorization header string true "Bearer token"
// @Param 		 nama formData string true "Nama Perlengkapan"
// @Param 		 deskripsi formData string true "Deskripsi Perlengkapan"
// @Param 		 harga formData int true "Harga Perlengkapan"
// @Param 		 stok formData int true "Stok Perlengkapan"
// @Param 		 image formData file true "Image Perlengkapan"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Perlengkapan}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /perlengkapan [post]
func (pc *Perlengkapan) Create(ctx *gin.Context) {
	perlengkapan := &model.PerlengkapanInput{}
	if err := ctx.ShouldBind(perlengkapan); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	perlengkapanData, err := pc.repoPerlengkapan.Create(perlengkapan)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 201, perlengkapanData)
}

// Update akan memperbarui data perlengkapan berdasarkan id
// @Summary      Memperbarui data perlengkapan berdasarkan id
// @Description  Memperbarui data perlengkapan berdasarkan id
// @Tags         Perlengkapan
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Perlengkapan"
// @Param 		 Authorization header string true "Bearer token"
// @Param 		 nama formData string true "Nama Perlengkapan"
// @Param 		 deskripsi formData string true "Deskripsi Perlengkapan"
// @Param 		 harga formData int true "Harga Perlengkapan"
// @Param 		 stok formData int true "Stok Perlengkapan"
// @Param 		 image formData file true "Image Perlengkapan"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Perlengkapan}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /perlengkapan/{id} [put]
func (pc *Perlengkapan) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	perlengkapan := &model.PerlengkapanInput{}
	if err := ctx.ShouldBind(perlengkapan); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	perlengkapanData, err := pc.repoPerlengkapan.Update(id, perlengkapan)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, perlengkapanData)
}

// Delete akan menghapus data perlengkapan berdasarkan id
// @Summary      Menghapus data perlengkapan berdasarkan id
// @Description  Menghapus data perlengkapan berdasarkan id
// @Tags         Perlengkapan
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Perlengkapan"
// @Param 		 Authorization header string true "Bearer token"
// @Success      200  {object}  utils.SuccessResponseData{data=string}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /perlengkapan/{id} [delete]
func (pc *Perlengkapan) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := pc.repoPerlengkapan.Delete(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, "perlengkapan berhasil dihapus")
}
