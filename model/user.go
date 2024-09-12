package model

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user model.
// swagger:model User
type User struct {
	ID string `json:"id" gorm:"type:char(36);primary_key"`
	// Email bersifat unik dan tidak boleh kosong
	Email string `json:"email" gorm:"type:varchar(255);unique;not null"`
	// Name tidak boleh kosong
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	// Role merupakan enum yang berisi "admin" dan "user" dan not null
	Role string `json:"role" gorm:"type:varchar(255);not null"`
	// Password disimpan dalam bentuk hash
	Password string `json:"-" gorm:"type:text;not null"`
	// Link KTP merupakan link ke file KTP
	LinkKTP string `json:"link_ktp" gorm:"type:text"`
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

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewV6()
	if err != nil {
		return err
	}
	u.ID = id.String()
	u.Role = "user"
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return
}

type UserCreate struct {
	Email           string                `form:"email" binding:"required,email"`
	Name            string                `form:"name" binding:"required"`
	Password        string                `form:"password" binding:"required"`
	ConfirmPassword string                `form:"confirm_password" binding:"required,eqfield=Password"`
	FileKTP         *multipart.FileHeader `form:"file_ktp" binding:"required"`
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
