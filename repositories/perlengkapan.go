package repositories

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"gorm.io/gorm"
)

type Perlengkapan struct {
	db *gorm.DB
}

func NewPerlengkapanRepository(db *gorm.DB) *Perlengkapan {
	return &Perlengkapan{db}
}

func (p *Perlengkapan) FindAll() ([]model.Perlengkapan, error) {
	var perlengkapans []model.Perlengkapan
	if err := p.db.Find(&perlengkapans).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "perlengkapan tidak ditemukan")
	}
	return perlengkapans, nil
}

func (p *Perlengkapan) FindByID(id string) (*model.Perlengkapan, error) {
	var perlengkapan model.Perlengkapan
	if err := p.db.First(&perlengkapan, "id = ?", id).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "perlengkapan tidak ditemukan")
	}
	return &perlengkapan, nil
}

func (p *Perlengkapan) Create(perlengkapan *model.PerlengkapanInput) error {
	if err := p.db.Create(perlengkapan.ToPerlengkapan()).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal membuat perlengkapan")
	}
	return nil
}

func (p *Perlengkapan) Update(id string, perlengkapan *model.PerlengkapanInput) error {
	if err := p.db.Model(&model.Perlengkapan{}).Where("id = ?", id).Updates(perlengkapan.ToPerlengkapan()).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal memperbarui perlengkapan")
	}
	return nil
}

func (p *Perlengkapan) Delete(id string) error {
	if err := p.db.Delete(&model.Perlengkapan{}, "id = ?", id).Error; err != nil {
		return utils.NewError(utils.ErrInternalServer, "gagal menghapus perlengkapan")
	}
	return nil
}
