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
	if len(Grounds) == 0 {
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

func (k *Ground) Create(Ground *model.GroundInput) (*model.Ground, error) {
	groundData := Ground.ToGround()
	groundData.BeforeCreate()

	link, err := utils.SaveFile(Ground.Image, "public/ground")
	if err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal membuat Ground")
	}

	groundData.ImageLink = link

	if err := k.db.Create(groundData).Error; err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal membuat Ground")
	}
	return groundData, nil
}

func (k *Ground) Update(id string, Ground *model.GroundInput) (*model.Ground, error) {
	groundData := Ground.ToGround()
	groundData.BeforeSave()
	groundData.ID = id
	link, err := utils.SaveFile(Ground.Image, "public/ground")
	if err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal membuat Ground")
	}
	groundData.ImageLink = link
	if err := k.db.Model(&model.Ground{}).Where("id = ?", id).Updates(groundData).Error; err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal memperbarui Ground")
	}
	return groundData, nil
}

func (k *Ground) Delete(id string) error {
	if err := k.db.Delete(&model.Ground{}, "id = ?", id).Error; err != nil {
		return utils.NewError(utils.ErrBadRequest, "gagal menghapus Ground")
	}
	return nil
}
