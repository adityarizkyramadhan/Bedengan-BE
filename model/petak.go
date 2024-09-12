package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kavling struct {
	ID         string         `json:"id" gorm:"type:varchar(36);primary_key"`
	Nama       string         `json:"nama" gorm:"type:varchar(255);not null"`
	KavlingID  string         `json:"kavling_id" gorm:"type:varchar(36);not null"` // Ganti dari uuid ke varchar(36)
	Harga      int            `json:"harga" gorm:"type:bigint;not null"`           // Sesuaikan tipe harga ke bigint
	JenisTenda string         `json:"jenis_tenda" gorm:"type:text;not null"`
	Status     string         `json:"status" gorm:"type:text;not null;default:'tersedia'"` // GORM default
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

func (p *Kavling) TableName() string {
	return "kavlings"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (p *Kavling) BeforeCreate() error {
	id, err := uuid.NewV6()
	if err != nil {
		return err
	}
	p.ID = id.String()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeSave will set the updated_at timestamp to current time.
func (p *Kavling) BeforeSave() error {
	p.UpdatedAt = time.Now()
	return nil
}

type KavlingInput struct {
	Nama       string `json:"nama" binding:"required"`
	KavlingID  string `json:"kavling_id" binding:"required"`
	Harga      int    `json:"harga" binding:"required"`
	JenisTenda string `json:"jenis_tenda" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

func (p *KavlingInput) ToKavling() *Kavling {
	return &Kavling{
		Nama:       p.Nama,
		KavlingID:  p.KavlingID,
		Harga:      p.Harga,
		JenisTenda: p.JenisTenda,
		Status:     p.Status,
	}
}
