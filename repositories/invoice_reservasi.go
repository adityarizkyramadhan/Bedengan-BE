package repositories

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"gorm.io/gorm"
)

type InvoiceReservasi struct {
	db *gorm.DB
}

func NewInvoiceReservasiRepository(db *gorm.DB) *InvoiceReservasi {
	return &InvoiceReservasi{db}
}

func (i *InvoiceReservasi) Create(userID string, inputInvoiceReservasi *model.InputInvoiceReservasi, inputReservasi []*model.InputReservasi) (*model.InvoiceReservasi, error) {
	reservasi := inputInvoiceReservasi.ToInvoiceReservasi()
	reservasi.BeforeCreate()
	reservasi.UserID = userID
	tx := i.db.Begin()
	if err := tx.Create(reservasi).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	var reservasiData []*model.Reservasi
	for _, r := range inputReservasi {
		reservasi := r.ToReservasi()
		reservasi.BeforeCreate()
		reservasi.InvoiceReservasiID = reservasi.ID
		reservasi.UserID = userID
		reservasiData = append(reservasiData, reservasi)
	}
	if err := tx.Create(&reservasiData).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return reservasi, nil
}

func (i *InvoiceReservasi) FindAll(userID string) ([]model.InvoiceReservasi, error) {
	var invoices []model.InvoiceReservasi
	if err := i.db.Where("user_id = ?", userID).Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (i *InvoiceReservasi) FindByID(userID, id string) (*model.InvoiceReservasi, error) {
	var invoice model.InvoiceReservasi
	if err := i.db.Where("user_id = ? AND id = ?", userID, id).First(&invoice).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (i *InvoiceReservasi) Update(userID, id string, inputInvoiceReservasi *model.InputInvoiceReservasi) (*model.InvoiceReservasi, error) {
	invoice := inputInvoiceReservasi.ToInvoiceReservasi()
	invoice.BeforeCreate()
	invoice.UserID = userID
	if err := i.db.Where("user_id = ? AND id = ?", userID, id).Updates(invoice).Error; err != nil {
		return nil, err
	}
	return invoice, nil
}

func (i *InvoiceReservasi) Delete(userID, id string) error {
	tx := i.db.Begin()
	// Hapus reservasi
	if err := tx.Where("user_id = ? AND invoice_reservasi_id = ?", userID, id).Delete(&model.Reservasi{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	// Hapus invoice
	if err := tx.Where("user_id = ? AND id = ?", userID, id).Delete(&model.InvoiceReservasi{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
