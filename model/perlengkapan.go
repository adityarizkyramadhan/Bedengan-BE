package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Perlengkapan struct {
	ID        string         `json:"id" gorm:"type:varchar(36);primaryKey;default:(UUID())"`
	Nama      string         `json:"nama" gorm:"type:varchar(255);not null"`
	Deskripsi string         `json:"deskripsi" gorm:"type:text;not null"`
	Jenis     string         `json:"jenis" gorm:"type:varchar(255)"`
	Harga     int            `json:"harga" gorm:"type:int;not null"`
	Stok      int            `json:"stok" gorm:"type:int;not null"`
	Image     string         `json:"image" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (p *Perlengkapan) TableName() string {
	return "perlengkapans"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (p *Perlengkapan) BeforeCreate() error {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeSave will set the updated_at timestamp to current time.
func (p *Perlengkapan) BeforeSave() error {
	p.UpdatedAt = time.Now()
	return nil
}

type PerlengkapanInput struct {
	Nama      string                `form:"nama" binding:"required"`
	Deskripsi string                `form:"deskripsi" binding:"required"`
	Jenis     string                `form:"jenis" binding:"required"`
	Harga     int                   `form:"harga" binding:"required"`
	Stok      int                   `form:"stok" binding:"required"`
	Image     *multipart.FileHeader `form:"image"`
}

func (p *PerlengkapanInput) ToPerlengkapan() *Perlengkapan {
	return &Perlengkapan{
		Nama:      p.Nama,
		Deskripsi: p.Deskripsi,
		Jenis:     p.Jenis,
		Harga:     p.Harga,
		Stok:      p.Stok,
	}
}
