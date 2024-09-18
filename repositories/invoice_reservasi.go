package repositories

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"gorm.io/gorm"
)

type InvoiceReservasi struct {
	db *gorm.DB
}

func NewInvoiceReservasiRepository(db *gorm.DB) *InvoiceReservasi {
	return &InvoiceReservasi{db}
}

func (i *InvoiceReservasi) Create(userID string, inputInvoiceReservasi *model.InputInvoiceReservasi) (*model.InvoiceReservasi, error) {
	invoiceReservasi := inputInvoiceReservasi.ToInvoiceReservasi()
	invoiceReservasi.BeforeCreate()
	invoiceReservasi.UserID = userID
	tx := i.db.Begin()
	// Jika kavling sudah ada yang reservasi maka tidak bisa reservasi
	var kavlingsID []string
	for _, r := range inputInvoiceReservasi.Reservasi {
		if r.KavlingID != nil {
			kavlingsID = append(kavlingsID, *r.KavlingID)
		}
	}

	var reservasiCount int64
	// Pada tanggal_kedatangan sampai tanggal_kepulangan, kavling sudah ada yang reservasi
	if err := tx.Preload("InvoiceReservasi").Model(&model.Reservasi{}).
		Joins("JOIN invoice_reservasis ON invoice_reservasis.id = reservasis.invoice_reservasi_id").
		Where("reservasis.kavling_id IN (?)", kavlingsID).
		Where("invoice_reservasis.tanggal_kedatangan <= ?", inputInvoiceReservasi.TanggalKepulangan).
		Where("invoice_reservasis.tanggal_kepulangan >= ?", inputInvoiceReservasi.TanggalKedatangan).
		Count(&reservasiCount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if reservasiCount > 0 {
		tx.Rollback()
		return nil, utils.NewError(utils.ErrBadRequest, "Kavling sudah ada yang reservasi")
	}

	// ambil pada perlengkapan dan kavling untuk mendapatkan harga
	var perlengkapanID []string
	for _, r := range inputInvoiceReservasi.Reservasi {
		if r.PerlengkapanID != nil {
			perlengkapanID = append(perlengkapanID, *r.PerlengkapanID)
		}
	}

	var perlengkapan []model.Perlengkapan
	if err := tx.Find(&perlengkapan, "id IN (?)", perlengkapanID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var kavling []model.Kavling
	if err := tx.Find(&kavling, "id IN (?)", kavlingsID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Hitung total harga
	var totalHargaPerHari int
	for _, r := range inputInvoiceReservasi.Reservasi {
		if r.KavlingID != nil {
			for _, k := range kavling {
				if k.ID == *r.KavlingID {
					totalHargaPerHari += k.Harga * r.Jumlah
				}
			}
		}
		if r.PerlengkapanID != nil {
			for _, p := range perlengkapan {
				if p.ID == *r.PerlengkapanID {
					totalHargaPerHari += p.Harga * r.Jumlah
				}
			}
		}
	}

	// Lama hari = tanggal kepulangan - tanggal kedatangan
	lamaHari, err := inputInvoiceReservasi.CalculateLamaHari()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	invoiceReservasi.Jumlah = totalHargaPerHari * lamaHari

	if err := tx.Create(invoiceReservasi).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	var reservasiData []*model.Reservasi
	for _, r := range inputInvoiceReservasi.Reservasi {
		reservasi := r.ToReservasi(invoiceReservasi)
		reservasi.BeforeCreate()
		reservasi.UserID = userID
		reservasiData = append(reservasiData, reservasi)
	}
	if err := tx.Create(&reservasiData).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return invoiceReservasi, nil
}

func (i *InvoiceReservasi) FindAll(userID string) ([]model.InvoiceReservasi, error) {
	var invoices []model.InvoiceReservasi
	if err := i.db.Where("user_id = ?", userID).Preload("Reservasi").Find(&invoices).Error; err != nil {
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
