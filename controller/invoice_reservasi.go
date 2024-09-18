package controller

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

type InvoiceReservasi struct {
	repo *repositories.InvoiceReservasi
}

func NewInvoiceReservasiController(repo *repositories.InvoiceReservasi) *InvoiceReservasi {
	return &InvoiceReservasi{repo}
}

// Create akan membuat data InvoiceReservasi baru
// @Summary      Membuat data InvoiceReservasi baru
// @Description  Membuat data InvoiceReservasi baru
// @Tags         InvoiceReservasi
// @Accept       json
// @Produce      json
// @Success      201  {object}  utils.SuccessResponseData{data=model.InvoiceReservasi}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /invoice-reservasi [post]
func (pc *InvoiceReservasi) Create(ctx *gin.Context) {
	var input model.InputInvoiceReservasi
	if err := ctx.ShouldBindJSON(&input); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	userID := ctx.MustGet("id").(string)

	newInvoiceReservasi, err := pc.repo.Create(userID, &input)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 201, newInvoiceReservasi)
}

// FindAll akan mengambil semua data InvoiceReservasi
// @Summary      Mengambil semua data InvoiceReservasi
// @Description  Mengambil semua data InvoiceReservasi
// @Tags         InvoiceReservasi
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.SuccessResponseData{data=[]model.InvoiceReservasi}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /invoice-reservasi [get]
func (pc *InvoiceReservasi) FindAll(ctx *gin.Context) {
	userID := ctx.MustGet("id").(string)
	invoiceReservasi, err := pc.repo.FindAll(userID)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, invoiceReservasi)
}

// FindByID akan mengambil data InvoiceReservasi berdasarkan id
// @Summary      Mengambil data InvoiceReservasi berdasarkan id
// @Description  Mengambil data InvoiceReservasi berdasarkan id
// @Tags         InvoiceReservasi
// @Accept       json
// @Produce      json
// @Param 		 id path string true "ID InvoiceReservasi"
// @Success      200  {object}  utils.SuccessResponseData{data=model.InvoiceReservasi}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /invoice-reservasi/{id} [get]
func (pc *InvoiceReservasi) FindByID(ctx *gin.Context) {
	userID := ctx.MustGet("id").(string)
	id := ctx.Param("id")

	invoiceReservasi, err := pc.repo.FindByID(userID, id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, invoiceReservasi)
}

// Delete akan menghapus data InvoiceReservasi berdasarkan id
// @Summary      Menghapus data InvoiceReservasi berdasarkan id
// @Description  Menghapus data InvoiceReservasi berdasarkan id
// @Tags         InvoiceReservasi
// @Accept       json
// @Produce      json
// @Param 		 id path string true "ID InvoiceReservasi"
// @Success      200  {object}  utils.SuccessResponseData{data=string}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /invoice-reservasi/{id} [delete]
func (pc *InvoiceReservasi) Delete(ctx *gin.Context) {
	userID := ctx.MustGet("id").(string)
	id := ctx.Param("id")

	if err := pc.repo.Delete(userID, id); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, "Invoice reservasi berhasil dihapus")
}

// Update akan mengupdate data InvoiceReservasi berdasarkan id
// @Summary      Mengupdate data InvoiceReservasi berdasarkan id
// @Description  Mengupdate data InvoiceReservasi berdasarkan id
// @Tags         InvoiceReservasi
// @Accept       json
// @Produce      json
// @Param 		 id path string true "ID InvoiceReservasi"
// @Success      200  {object}  utils.SuccessResponseData{data=model.InvoiceReservasi}
// @Failure      500  {object}  utils.ErrorResponseData
// @Router       /invoice-reservasi/{id} [put]
func (pc *InvoiceReservasi) Update(ctx *gin.Context) {
	userID := ctx.MustGet("id").(string)
	id := ctx.Param("id")

	var input model.InputInvoiceReservasi
	if err := ctx.ShouldBindJSON(&input); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	invoiceReservasi, err := pc.repo.Update(userID, id, &input)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, invoiceReservasi)
}
