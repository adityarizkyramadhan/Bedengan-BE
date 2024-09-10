package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Petak struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Nama       string         `json:"nama" gorm:"type:varchar(255);not null"`
	KavlingID  uuid.UUID      `json:"kavling_id" gorm:"type:uuid;not null"`
	Harga      int            `json:"harga" gorm:"type:int;not null"`
	JenisTenda string         `json:"jenis_tenda" gorm:"type:text;not null"`
	Status     string         `json:"status" gorm:"type:enum('tersedia', 'terpesan', 'terjual');not null;default:'tersedia'"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

func (p *Petak) TableName() string {
	return "petaks"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (p *Petak) BeforeCreate() error {
	p.ID = uuid.Must(uuid.NewV6())
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}

// BeforeSave will set the updated_at timestamp to current time.
func (p *Petak) BeforeSave() error {
	p.UpdatedAt = time.Now()
	return nil
}

type PetakInput struct {
	Nama       string `json:"nama" binding:"required"`
	KavlingID  string `json:"kavling_id" binding:"required"`
	Harga      int    `json:"harga" binding:"required"`
	JenisTenda string `json:"jenis_tenda" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

func (p *PetakInput) ToPetak() *Petak {
	return &Petak{
		Nama:       p.Nama,
		KavlingID:  uuid.MustParse(p.KavlingID),
		Harga:      p.Harga,
		JenisTenda: p.JenisTenda,
		Status:     p.Status,
	}
}
