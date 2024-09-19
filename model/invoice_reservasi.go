package model

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceReservasi struct {
	ID                string         `json:"id" gorm:"type:varchar(36);primaryKey"`
	UserID            string         `json:"user_id" gorm:"type:varchar(36)"`
	User              User           `json:"-" gorm:"foreignKey:UserID"`
	NomorInvoice      string         `json:"nomor_invoice" gorm:"type:varchar(50)"`
	JenisPengunjung   string         `json:"jenis_pengunjung" binding:"required"`
	Total             int            `json:"total"`
	LinkPembayaran    string         `json:"link_pembayaran"`
	LinkPerizinan     string         `json:"link_perizinan"`
	Jumlah            int            `json:"jumlah" gorm:"default:0"`
	Status            string         `json:"status" gorm:"type:text;default:'menunggu_pembayaran'"`
	TanggalKedatangan time.Time      `json:"tanggal_kedatangan"`
	TanggalKepulangan time.Time      `json:"tanggal_kepulangan" binding:"required"`
	Reservasi         []Reservasi    `json:"reservasi" gorm:"foreignKey:InvoiceReservasiID"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `json:"-" gorm:"index"`
}

type InvoiceReservasiDTO struct {
	ID                string         `json:"id" gorm:"type:varchar(36);primaryKey"`
	UserID            string         `json:"user_id" gorm:"type:varchar(36)"`
	User              User           `json:"-" gorm:"foreignKey:UserID"`
	NomorInvoice      string         `json:"nomor_invoice" gorm:"type:varchar(50)"`
	JenisPengunjung   string         `json:"jenis_pengunjung" binding:"required"`
	Total             int            `json:"total"`
	LinkPembayaran    string         `json:"link_pembayaran"`
	LinkPerizinan     string         `json:"link_perizinan"`
	Jumlah            int            `json:"jumlah" gorm:"default:0"`
	Status            string         `json:"status" gorm:"type:text;default:'menunggu_pembayaran'"`
	TanggalKedatangan time.Time      `json:"tanggal_kedatangan"`
	TanggalKepulangan time.Time      `json:"tanggal_kepulangan" binding:"required"`
	Reservasi         []ReservasiDTO `json:"reservasi" gorm:"foreignKey:InvoiceReservasiID"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
}

func (i *InvoiceReservasi) ToDTO() InvoiceReservasiDTO {
	var reservasiDTO []ReservasiDTO
	for _, r := range i.Reservasi {
		reservasiDTO = append(reservasiDTO, r.ToDTO())
	}

	return InvoiceReservasiDTO{
		ID:                i.ID,
		UserID:            i.UserID,
		User:              i.User,
		NomorInvoice:      i.NomorInvoice,
		JenisPengunjung:   i.JenisPengunjung,
		Total:             i.Total,
		LinkPembayaran:    i.LinkPembayaran,
		LinkPerizinan:     i.LinkPerizinan,
		Jumlah:            i.Jumlah,
		Status:            i.Status,
		TanggalKedatangan: i.TanggalKedatangan,
		TanggalKepulangan: i.TanggalKepulangan,
		Reservasi:         reservasiDTO,
		CreatedAt:         i.CreatedAt,
		UpdatedAt:         i.UpdatedAt,
	}
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
	InvoiceReservasi   InvoiceReservasi `json:"-" gorm:"foreignKey:InvoiceReservasiID"`
	PerlengkapanID     *string          `json:"perlengkapan_id" gorm:"type:varchar(36);default:null"`
	Perlengkapan       Perlengkapan     `json:"-" gorm:"foreignKey:PerlengkapanID"`
	KavlingID          *string          `json:"kavling_id" gorm:"type:varchar(36);default:null"`
	Kavling            Kavling          `json:"-" gorm:"foreignKey:KavlingID"`
	UserID             string           `json:"user_id" gorm:"type:varchar(36)"`
	User               User             `json:"-" gorm:"foreignKey:UserID"`
	Jumlah             int              `json:"jumlah"`
	Harga              int              `json:"harga"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
	DeletedAt          gorm.DeletedAt   `json:"-" gorm:"index"`
}
type ReservasiDTO struct {
	ID                 string       `json:"id" gorm:"type:varchar(36);primaryKey"`
	InvoiceReservasiID string       `json:"invoice_reservasi_id" gorm:"type:varchar(36)"`
	PerlengkapanID     *string      `json:"perlengkapan_id" gorm:"type:varchar(36);default:null"`
	Perlengkapan       Perlengkapan `json:"perlengkapan" gorm:"foreignKey:PerlengkapanID"`
	KavlingID          *string      `json:"kavling_id" gorm:"type:varchar(36);default:null"`
	Kavling            Kavling      `json:"kavling" gorm:"foreignKey:KavlingID"`
	UserID             string       `json:"user_id" gorm:"type:varchar(36)"`
	Jumlah             int          `json:"jumlah"`
	Harga              int          `json:"harga"`
	CreatedAt          time.Time    `json:"created_at"`
	UpdatedAt          time.Time    `json:"updated_at"`
}

func (r *Reservasi) ToDTO() ReservasiDTO {
	return ReservasiDTO{
		ID:                 r.ID,
		InvoiceReservasiID: r.InvoiceReservasiID,
		PerlengkapanID:     r.PerlengkapanID,
		Perlengkapan:       r.Perlengkapan,
		KavlingID:          r.KavlingID,
		Kavling:            r.Kavling,
		UserID:             r.UserID,
		Jumlah:             r.Jumlah,
		Harga:              r.Harga,
		CreatedAt:          r.CreatedAt,
		UpdatedAt:          r.UpdatedAt,
	}
}

func (r *Reservasi) BeforeSave(tx *gorm.DB) error {
	if r.KavlingID == nil && r.PerlengkapanID == nil {
		return errors.New("either KavlingID or PerlengkapanID must be provided")
	}
	return nil
}

func (r *Reservasi) BeforeCreate() {
	r.ID = uuid.New().String()
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
}

type InputInvoiceReservasi struct {
	JenisPengunjung   string           `json:"jenis_pengunjung" binding:"required"`
	TanggalKedatangan string           `json:"tanggal_kedatangan" binding:"required"`
	TanggalKepulangan string           `json:"tanggal_kepulangan" binding:"required"`
	Reservasi         []InputReservasi `json:"reservasi" binding:"required"`
}

type InputReservasi struct {
	PerlengkapanID *string `json:"perlengkapan_id,omitempty"`
	KavlingID      *string `json:"kavling_id,omitempty"`
	Jumlah         int     `json:"jumlah" binding:"required"`
}

func (input *InputInvoiceReservasi) CalculateLamaHari() (int, error) {
	const layout = "2006-01-02" // Adjust the layout according to your date format

	tanggalKedatangan, err := time.Parse(layout, input.TanggalKedatangan)
	if err != nil {
		return 0, fmt.Errorf("invalid TanggalKedatangan: %v", err)
	}

	tanggalKepulangan, err := time.Parse(layout, input.TanggalKepulangan)
	if err != nil {
		return 0, fmt.Errorf("invalid TanggalKepulangan: %v", err)
	}

	if tanggalKepulangan.Before(tanggalKedatangan) {
		return 0, fmt.Errorf("TanggalKepulangan cannot be before TanggalKedatangan")
	}

	// Add 1 to include the departure date as a full day
	lamaHari := int(tanggalKepulangan.Sub(tanggalKedatangan).Hours()/24) + 1
	return lamaHari, nil
}

func (i *InputInvoiceReservasi) ToInvoiceReservasi() *InvoiceReservasi {
	// i.TanggalKedatangan string jadikan time.Time
	tanggalKedatangan, _ := time.Parse("2006-01-02", i.TanggalKedatangan)
	// i.TanggalKepulangan string jadikan time.Time
	tanggalKepulangan, _ := time.Parse("2006-01-02", i.TanggalKepulangan)

	invoiceReservasi := &InvoiceReservasi{
		JenisPengunjung:   i.JenisPengunjung,
		TanggalKedatangan: tanggalKedatangan,
		TanggalKepulangan: tanggalKepulangan,
	}
	return invoiceReservasi
}

func (i *InputReservasi) ToReservasi(invReservasi *InvoiceReservasi) *Reservasi {
	reservasi := &Reservasi{
		InvoiceReservasiID: invReservasi.ID,
		PerlengkapanID:     i.PerlengkapanID,
		KavlingID:          i.KavlingID,
		Jumlah:             i.Jumlah,
	}
	return reservasi
}

type InvoiceReservasiFile struct {
	Perizinan  *multipart.FileHeader `form:"perizinan" binding:"required"`
	Pembayaran *multipart.FileHeader `form:"pembayaran" binding:"required"`
}
