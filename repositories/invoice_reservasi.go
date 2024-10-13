package repositories

import (
	"time"

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
	invoiceReservasi.Status = "menunggu_pembayaran"
	invoiceReservasi.UserID = &userID
	invoiceReservasi.Tipe = inputInvoiceReservasi.Tipe
	tx := i.db.Begin()
	// Jika kavling sudah ada yang reservasi maka tidak bisa reservasi
	var kavlingsID []string
	for _, r := range inputInvoiceReservasi.Reservasi {
		if r.KavlingID != nil {
			kavlingsID = append(kavlingsID, *r.KavlingID)
		}
	}

	// ambil user ambil rolenya
	var user model.User
	if err := tx.First(&user, "id = ?", userID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if user.Role == "admin" {
		invoiceReservasi.Tipe = "offline"
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

	// Set is_available menjadi false
	for _, r := range inputInvoiceReservasi.Reservasi {
		if r.KavlingID != nil {
			if err := tx.Model(&model.Kavling{}).Where("id = ?", *r.KavlingID).Update("is_available", false).Error; err != nil {
				tx.Rollback()
				return nil, err
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

func (i *InvoiceReservasi) FindAll(userID string) ([]model.InvoiceReservasiDTO, error) {
	var invoices []model.InvoiceReservasi
	if err := i.db.Order("created_at DESC").Where("user_id = ?", userID).Preload("Reservasi.Kavling").Preload("Reservasi.Perlengkapan").Find(&invoices).Error; err != nil {
		return nil, err
	}
	var invoicesDTO []model.InvoiceReservasiDTO
	for _, i := range invoices {
		invoicesDTO = append(invoicesDTO, i.ToDTO())
	}
	return invoicesDTO, nil
}

func (i *InvoiceReservasi) FindByID(userID, id string) (*model.InvoiceReservasi, error) {
	var invoice model.InvoiceReservasi
	db := i.db.Preload("Reservasi.Kavling").Preload("Reservasi.Perlengkapan")
	if userID != "" {
		db = db.Where("user_id = ?", userID)
	}
	if err := db.First(&invoice, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (i *InvoiceReservasi) Update(userID, id string, inputInvoiceReservasi *model.InputInvoiceReservasi) (*model.InvoiceReservasi, error) {
	// Update hanya bisa mengubah tanggal kedatangan dan tanggal kepulangan dan status serta link pembayaran
	invoiceReservasi, err := i.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	timeParse := "2006-01-02"
	invoiceReservasi.TanggalKedatangan, err = time.Parse(timeParse, inputInvoiceReservasi.TanggalKedatangan)
	if err != nil {
		return nil, err
	}
	invoiceReservasi.TanggalKepulangan, err = time.Parse(timeParse, inputInvoiceReservasi.TanggalKepulangan)
	if err != nil {
		return nil, err
	}

	tx := i.db.Begin()

	// Cari dulu ditanggal kedatangan dan tanggal kepulangan apakah ada yang reservasi
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
		Where("invoice_reservasis.id != ?", id).
		Count(&reservasiCount).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if reservasiCount > 0 {
		tx.Rollback()
		return nil, utils.NewError(utils.ErrBadRequest, "Kavling sudah ada yang reservasi")
	}

	if err := tx.Save(invoiceReservasi).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return invoiceReservasi, nil
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

func (i *InvoiceReservasi) UpdateFile(userID, id string, fileInput *model.InvoiceReservasiFile) (*model.InvoiceReservasi, error) {
	invoiceReservasi, err := i.FindByID(userID, id)
	if err != nil {
		return nil, err
	}
	if fileInput.Pembayaran != nil {
		link, err := utils.SaveFile(fileInput.Pembayaran, "public/invoice")
		if err != nil {
			return nil, err
		}
		invoiceReservasi.LinkPembayaran = link
		invoiceReservasi.Status = "menunggu_verifikasi"
	}

	if fileInput.Perizinan != nil {
		link, err := utils.SaveFile(fileInput.Perizinan, "public/perizinan")
		if err != nil {
			return nil, err
		}
		invoiceReservasi.LinkPerizinan = link
	}

	if err := i.db.Save(invoiceReservasi).Error; err != nil {
		return nil, err
	}

	return invoiceReservasi, nil
}

func (i *InvoiceReservasi) VerifikasiInvoice(id string) (*model.InvoiceReservasi, error) {
	invoiceReservasi, err := i.FindByID("", id)
	if err != nil {
		return nil, err
	}
	invoiceReservasi.Status = "verifikasi"
	if err := i.db.Save(invoiceReservasi).Error; err != nil {
		return nil, err
	}
	return invoiceReservasi, nil
}

func (i *InvoiceReservasi) TolakInvoice(id string) (*model.InvoiceReservasi, error) {
	invoiceReservasi, err := i.FindByID("", id)
	if err != nil {
		return nil, err
	}
	invoiceReservasi.Status = "ditolak"
	if err := i.db.Save(invoiceReservasi).Error; err != nil {
		return nil, err
	}
	return invoiceReservasi, nil
}

func (i *InvoiceReservasi) AdminFindAll() (interface{}, error) {
	var invoices []model.InvoiceReservasi
	if err := i.db.Preload("Reservasi.Kavling").Preload("Reservasi.Perlengkapan").Find(&invoices).Error; err != nil {
		return nil, err
	}
	var invoicesDTO []model.InvoiceReservasiDTO
	for _, i := range invoices {
		invoicesDTO = append(invoicesDTO, i.ToDTO())
	}

	response := make(map[string]interface{})

	jumlahPembayaran := 0
	jumlahOnline := 0
	jumlahOffline := 0
	jumlahVerifikasi := 0

	for _, i := range invoicesDTO {
		if i.Status == "verifikasi" {
			jumlahPembayaran += i.Jumlah
			jumlahVerifikasi++
		}
		if i.Tipe == "offline" {
			jumlahOffline++
		} else {
			jumlahOnline++
		}
	}

	response["jumlah_pembayaran"] = jumlahPembayaran
	response["invoices"] = invoicesDTO
	response["jumlah_online"] = jumlahOnline
	response["jumlah_offline"] = jumlahOffline
	response["jumlah_verifikasi"] = jumlahVerifikasi
	response["perbandingan_sukses"] = (float64(jumlahVerifikasi) / float64(len(invoicesDTO))) * 100.0
	return response, nil
}

func (i *InvoiceReservasi) DailyCheck() error {
	if err := utils.DailyCheckKavlingRawSQL(i.db); err != nil {
		return err
	}
	return nil
}
