package model

import (
	"time"

	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceReservasi struct {
	ID                string         `json:"id" gorm:"type:varchar(36);primaryKey"`
	UserID            string         `json:"user_id" gorm:"type:varchar(36)"`
	User              User           `json:"user" gorm:"foreignKey:UserID"`
	NomorInvoice      string         `json:"nomor_invoice" gorm:"type:varchar(50)"`
	JenisPengunjung   string         `json:"jenis_pengunjung" binding:"required"`
	Total             int            `json:"total"`
	LinkPembayaran    string         `json:"link_pembayaran"`
	LinkPerizinan     string         `json:"link_perizinan"`
	Status            string         `json:"status"`
	Jumlah            int            `json:"jumlah"`
	TanggalKedatangan time.Time      `json:"tanggal_kedatangan"`
	TanggalKepulangan time.Time      `json:"tanggal_kepulangan" binding:"required"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (i *InvoiceReservasi) BeforeCreate() {
	i.ID = uuid.New().String()
	i.NomorInvoice = utils.GenerateNomorInvoice()
	i.CreatedAt = time.Now()
	i.UpdatedAt = time.Now()
}

type Reservasi struct {
	ID                 string           `json:"id" gorm:"type:varchar(36);primaryKey"`
	InvoiceReservasiID string           `json:"invoice_reservasi_id" gorm:"type:varchar(36)"`
	InvoiceReservasi   InvoiceReservasi `json:"invoice_reservasi" gorm:"foreignKey:InvoiceReservasiID"`
	PerlengkapanID     string           `json:"perlengkapan_id" gorm:"type:varchar(36)"`
	Perlengkapan       Perlengkapan     `json:"perlengkapan" gorm:"foreignKey:PerlengkapanID"`
	KavlingID          string           `json:"kavling_id" gorm:"type:varchar(36)"`
	Kavling            Kavling          `json:"kavling" gorm:"foreignKey:KavlingID"`
	UserID             string           `json:"user_id" gorm:"type:varchar(36)"`
	User               User             `json:"user" gorm:"foreignKey:UserID"`
	Jumlah             int              `json:"jumlah"`
	Harga              int              `json:"harga"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
	DeletedAt          gorm.DeletedAt   `json:"deleted_at" gorm:"index"`
}

func (r *Reservasi) BeforeCreate() {
	r.ID = uuid.New().String()
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

type InputInvoiceReservasi struct {
	JenisPengunjung   string           `json:"jenis_pengunjung" binding:"required"`
	Jumlah            int              `json:"jumlah" binding:"required"`
	TanggalKedatangan string           `json:"tanggal_kedatangan" binding:"required"`
	TanggalKepulangan string           `json:"tanggal_kepulangan" binding:"required"`
	Reservasi         []InputReservasi `json:"reservasi" binding:"required"`
}

type InputReservasi struct {
	PerlengkapanID string `json:"perlengkapan_id" binding:"required"`
	KavlingID      string `json:"kavling_id" binding:"required"`
	Jumlah         int    `json:"jumlah" binding:"required"`
}

func (i *InputInvoiceReservasi) ToInvoiceReservasi() *InvoiceReservasi {
	// i.TanggalKedatangan string jadikan time.Time
	tanggalKedatangan, _ := time.Parse("2006-01-02", i.TanggalKedatangan)
	// i.TanggalKepulangan string jadikan time.Time
	tanggalKepulangan, _ := time.Parse("2006-01-02", i.TanggalKepulangan)

	invoiceReservasi := &InvoiceReservasi{
		JenisPengunjung:   i.JenisPengunjung,
		Jumlah:            i.Jumlah,
		TanggalKedatangan: tanggalKedatangan,
		TanggalKepulangan: tanggalKepulangan,
	}
	return invoiceReservasi
}

func (i *InputReservasi) ToReservasi() *Reservasi {
	reservasi := &Reservasi{
		PerlengkapanID: i.PerlengkapanID,
		KavlingID:      i.KavlingID,
		Jumlah:         i.Jumlah,
	}
	return reservasi
}
