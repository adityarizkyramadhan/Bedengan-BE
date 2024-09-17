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
