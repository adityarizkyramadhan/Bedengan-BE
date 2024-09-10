package repositories

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"gorm.io/gorm"
)

type Kavling struct {
	db *gorm.DB
}

func NewKavlingRepository(db *gorm.DB) *Kavling {
	return &Kavling{db}
}

func (k *Kavling) FindAll() ([]model.Kavling, error) {
	var kavlings []model.Kavling
	if err := k.db.Find(&kavlings).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "kavling tidak ditemukan")
	}
	return kavlings, nil
}

func (k *Kavling) FindByID(id string) (*model.Kavling, error) {
	var kavling model.Kavling
	if err := k.db.First(&kavling, "id = ?", id).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "kavling tidak ditemukan")
	}
	return &kavling, nil
}

func (k *Kavling) Create(kavling *model.KavlingInput) error {
	if err := k.db.Create(kavling.ToKavling()).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal membuat kavling")
	}
	return nil
}

func (k *Kavling) Update(id string, kavling *model.KavlingInput) error {
	if err := k.db.Model(&model.Kavling{}).Where("id = ?", id).Updates(kavling.ToKavling()).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal memperbarui kavling")
	}
	return nil
}

func (k *Kavling) Delete(id string) error {
	if err := k.db.Delete(&model.Kavling{}, "id = ?", id).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal menghapus kavling")
	}
	return nil
}
