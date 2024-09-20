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
	router.GET("/admin/invoice-reservasi", middleware.JWTMiddleware([]string{"admin"}), p.ctrlInvoiceReservasi.AdminFindAll)
	router.GET("/invoice-reservasi", middleware.JWTMiddleware([]string{"admin", "user"}), p.ctrlInvoiceReservasi.FindAll)
	router.GET("/invoice-reservasi/:id", middleware.JWTMiddleware([]string{"admin", "user"}), p.ctrlInvoiceReservasi.FindByID)
	router.PUT("/invoice-reservasi/:id", middleware.JWTMiddleware([]string{"admin", "user"}), p.ctrlInvoiceReservasi.Update)
	router.PUT("/invoice-reservasi/:id/file", middleware.JWTMiddleware([]string{"admin", "user"}), p.ctrlInvoiceReservasi.UpdateFile)
	router.PUT("/invoice-reservasi/:id/confirm", middleware.JWTMiddleware([]string{"admin"}), p.ctrlInvoiceReservasi.VerifikasiInvoice)
	router.PUT("/invoice-reservasi/:id/reject", middleware.JWTMiddleware([]string{"admin"}), p.ctrlInvoiceReservasi.TolakInvoice)
	router.DELETE("/invoice-reservasi/:id", middleware.JWTMiddleware([]string{"superadmin"}), p.ctrlInvoiceReservasi.Delete)
}
