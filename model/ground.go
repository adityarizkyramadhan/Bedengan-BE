package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Ground struct {
	ID         string         `json:"id" gorm:"type:varchar(36);primary_key"`
	Nama       string         `json:"nama" gorm:"type:varchar(255);not null"`
	SubGrounds []SubGround    `json:"sub_grounds"`
	ImageLink  string         `json:"image_link" gorm:"type:text"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

func (k *Ground) TableName() string {
	return "grounds"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (k *Ground) BeforeCreate() error {
	k.ID = uuid.New().String()
	k.CreatedAt = time.Now()
	k.UpdatedAt = time.Now()
	return nil
}

// BeforeSave will set the updated_at timestamp to current time.
func (k *Ground) BeforeSave() error {
	k.UpdatedAt = time.Now()
	return nil
}

type GroundInput struct {
	Nama  string                `form:"nama" binding:"required"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

func (k *GroundInput) ToGround() *Ground {
	return &Ground{
		Nama: k.Nama,
	}
}
