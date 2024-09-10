package controller

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

type Petak struct {
	repoPetak *repositories.Petak
}

func NewPetakController(repoPetak *repositories.Petak) *Petak {
	return &Petak{repoPetak}
}

// FindAll akan mengembalikan semua data petak
// @Summary      Menampilkan semua data petak
// @Description  Menampilkan semua data petak
// @Tags         Petak
// @Accept       json
// @Produce      json
// @Params 		 id_petak query string false "ID Petak"
// @Success      201  {object}  utils.SuccessResponseData{data=[]model.Petak}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /petak [get]
func (pc *Petak) FindAll(ctx *gin.Context) {
	idPetak := ctx.Query("id_petak")
	petaks, err := pc.repoPetak.FindAll(idPetak)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, petaks)
}

// FindByID akan mengembalikan data petak berdasarkan id
// @Summary      Menampilkan data petak berdasarkan id
// @Description  Menampilkan data petak berdasarkan id
// @Tags         Petak
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Petak"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Petak}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /petak/{id} [get]
func (pc *Petak) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	petak, err := pc.repoPetak.FindByID(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, petak)
}

// Create akan membuat data petak baru
// @Summary      Membuat data petak baru
// @Description  Membuat data petak baru
// @Tags         Petak
// @Accept       json
// @Produce      json
// @Param        petak     body    model.PetakInput     true  "Data Petak"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Petak}
// @Failure      422  {object}  utils.ErrorResponseData
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /petak [post]
func (pc *Petak) Create(ctx *gin.Context) {
	petak := &model.PetakInput{}
	if err := ctx.ShouldBindJSON(petak); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	if err := pc.repoPetak.Create(petak); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 201, petak)
}

// Update akan memperbarui data petak berdasarkan id
// @Summary      Memperbarui data petak berdasarkan id
// @Description  Memperbarui data petak berdasarkan id
// @Tags         Petak
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Petak"
// @Param        petak     body    model.PetakInput     true  "Data Petak"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Petak}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /petak/{id} [put]
func (pc *Petak) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	petak := &model.PetakInput{}
	if err := ctx.ShouldBindJSON(petak); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	if err := pc.repoPetak.Update(id, petak); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 201, petak)
}

// Delete akan menghapus data petak berdasarkan id
// @Summary      Menghapus data petak berdasarkan id
// @Description  Menghapus data petak berdasarkan id
// @Tags         Petak
// @Accept       json
// @Produce      json
// @Param        id     path    string     true  "ID Petak"
// @Success      201  {object}  utils.SuccessResponseData{data=model.Petak}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /petak/{id} [delete]
func (pc *Petak) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := pc.repoPetak.Delete(id); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	utils.SuccessResponse(ctx, 200, "Petak berhasil dihapus")
}
