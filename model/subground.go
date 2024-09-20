package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubGround struct {
	ID        string         `json:"id" gorm:"type:varchar(36);primary_key"`
	Nama      string         `json:"nama" gorm:"type:varchar(255);not null"`
	GroundID  string         `json:"ground_id" gorm:"type:varchar(36);not null"`
	Kavlings  []Kavling      `json:"kavlings"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (k *SubGround) TableName() string {
	return "sub_grounds"
}

// BeforeCreate will set a UUID rather than numeric ID.
func (k *SubGround) BeforeCreate() error {
	k.ID = uuid.New().String()
	k.CreatedAt = time.Now()
	k.UpdatedAt = time.Now()
	return nil
}

// BeforeSave will set the updated_at timestamp to current time.
func (k *SubGround) BeforeSave() error {
	k.UpdatedAt = time.Now()
	return nil
}

type SubGroundInput struct {
	Nama     string `json:"nama" binding:"required"`
	GroundID string `json:"ground_id" binding:"required"`
}

func (k *SubGroundInput) ToSubGround() *SubGround {
	return &SubGround{
		Nama:     k.Nama,
		GroundID: k.GroundID,
	}
}
