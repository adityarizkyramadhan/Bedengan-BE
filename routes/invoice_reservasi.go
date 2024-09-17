package routes

import (
	"github.com/adityarizkyramadhan/template-go-mvc/controller"
	"github.com/adityarizkyramadhan/template-go-mvc/middleware"
	"github.com/gin-gonic/gin"
)

type InvoiceReservasi struct {
	ctrlInvoiceReservasi *controller.InvoiceReservasi
}

func NewInvoiceReservasiRoutes(ctrlInvoiceReservasi *controller.InvoiceReservasi) *InvoiceReservasi {
	return &InvoiceReservasi{ctrlInvoiceReservasi}
}

// SetupRoutes will setup the routes for InvoiceReservasi
func (p *InvoiceReservasi) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/invoice-reservasi", middleware.JWTMiddleware([]string{"admin", "user"}), p.ctrlInvoiceReservasi.Create)
}
