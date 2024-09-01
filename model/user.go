package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user model.
// swagger:model User
type User struct {
	ID uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	// Email bersifat unik dan tidak boleh kosong
	Email string `json:"email" gorm:"type:varchar(255);unique;not null"`
	// Name tidak boleh kosong
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	// Role merupakan enum yang berisi "admin" dan "user" dan not null
	Role string `json:"role" gorm:"type:varchar(255);not null"`
	// Password disimpan dalam bentuk hash
	Password string `json:"-" gorm:"type:text;not null"`
	// OTP digunakan untuk konfirmasi email
	OTP string `json:"-" gorm:"type:varchar(5)"`
	// IsVerified menandakan apakah email sudah diverifikasi
	IsVerified bool `json:"-" gorm:"type:tinyint(1);default:0"`
	// Provinsi tempat tinggal user
	Province string `json:"province" gorm:"type:varchar(255)"`
	// Kota tempat tinggal user
	City string `json:"city" gorm:"type:varchar(255)"`
	// CreatedAt menandakan waktu user dibuat
	CreatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	// UpdatedAt menandakan waktu user terakhir diupdate
	UpdatedAt time.Time `json:"-" gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	// DeletedAt menandakan waktu user dihapus
	DeletedAt gorm.DeletedAt `json:"-" gorm:"type:timestamp"`
}

func (u User) TableName() string {
	return "users"
}

type UserCreate struct {
	Email           string `json:"email" binding:"required,email"`
	Name            string `json:"name" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	Province        string `json:"province" binding:"required"`
	City            string `json:"city" binding:"required"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	Province string `json:"province"`
	City     string `json:"city"`
}
