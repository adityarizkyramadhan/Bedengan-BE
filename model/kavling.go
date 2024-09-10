package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Kavling struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
	Nama      string         `json:"nama" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (k *Kavling) TableName() string {
	return "kavlings"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (k *Kavling) BeforeCreate() error {
	k.ID = uuid.Must(uuid.NewV6())
	k.CreatedAt = time.Now()
	k.UpdatedAt = time.Now()
	return nil
}

// BeforeSave will set the updated_at timestamp to current time.
func (k *Kavling) BeforeSave() error {
	k.UpdatedAt = time.Now()
	return nil
}

type KavlingInput struct {
	Nama string `json:"nama" binding:"required"`
}

func (k *KavlingInput) ToKavling() *Kavling {
	return &Kavling{
		Nama: k.Nama,
	}
}
