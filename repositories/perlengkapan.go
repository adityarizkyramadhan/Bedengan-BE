package repositories

import (
	"strings"

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

func (p *Perlengkapan) FindAll(query map[string]string) ([]model.Perlengkapan, error) {
	var perlengkapans []model.Perlengkapan
	db := p.db

	if query["jenis"] != "" {
		// isi dari query antara tenda_paket,tenda_non_paket
		// Jadi harus displit by comma
		jenis := strings.Split(query["jenis"], ",")
		db = db.Where("jenis IN (?)", jenis)
	}

	if err := db.Find(&perlengkapans).Error; err != nil {
		return nil, utils.NewError(utils.ErrNotFound, "perlengkapan tidak ditemukan")
	}
	if len(perlengkapans) == 0 {
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

func (p *Perlengkapan) Create(perlengkapan *model.PerlengkapanInput) (*model.Perlengkapan, error) {
	perlengkapanData := perlengkapan.ToPerlengkapan()
	perlengkapanData.BeforeCreate()

	link, err := utils.SaveFile(perlengkapan.Image, "public/perlengkapan")
	if err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, err.Error())
	}

	perlengkapanData.Image = link

	if err := p.db.Create(perlengkapanData).Error; err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal membuat perlengkapan")
	}
	return perlengkapanData, nil
}

func (p *Perlengkapan) Update(id string, perlengkapan *model.PerlengkapanInput) (*model.Perlengkapan, error) {
	perlengkapanData := perlengkapan.ToPerlengkapan()
	perlengkapanData.BeforeSave()
	perlengkapanData.ID = id
	if perlengkapan.Image != nil {
		link, err := utils.SaveFile(perlengkapan.Image, "public/perlengkapan")
		if err != nil {
			return nil, utils.NewError(utils.ErrBadRequest, err.Error())
		}
		perlengkapanData.Image = link
	}
	if err := p.db.Model(&model.Perlengkapan{}).Where("id = ?", id).Updates(perlengkapanData).Error; err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal memperbarui perlengkapan")
	}
	return perlengkapanData, nil
}

func (p *Perlengkapan) Delete(id string) error {
	if err := p.db.Delete(&model.Perlengkapan{}, "id = ?", id).Error; err != nil {
		return utils.NewError(utils.ErrBadRequest, "gagal menghapus perlengkapan")
	}
	return nil
}
