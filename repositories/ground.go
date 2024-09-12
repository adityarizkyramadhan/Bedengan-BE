package repositories

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"gorm.io/gorm"
)

type Ground struct {
	db *gorm.DB
}

func NewGroundRepository(db *gorm.DB) *Ground {
	return &Ground{db}
}

func (k *Ground) FindAll() ([]model.Ground, error) {
	var Grounds []model.Ground
	if err := k.db.Find(&Grounds).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "Ground tidak ditemukan")
	}
	return Grounds, nil
}

func (k *Ground) FindByID(id string) (*model.Ground, error) {
	var Ground model.Ground
	if err := k.db.First(&Ground, "id = ?", id).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "Ground tidak ditemukan")
	}
	return &Ground, nil
}

func (k *Ground) Create(Ground *model.GroundInput) error {
	if err := k.db.Create(Ground.ToGround()).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal membuat Ground")
	}
	return nil
}

func (k *Ground) Update(id string, Ground *model.GroundInput) error {
	if err := k.db.Model(&model.Ground{}).Where("id = ?", id).Updates(Ground.ToGround()).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal memperbarui Ground")
	}
	return nil
}

func (k *Ground) Delete(id string) error {
	if err := k.db.Delete(&model.Ground{}, "id = ?", id).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal menghapus Ground")
	}
	return nil
}
