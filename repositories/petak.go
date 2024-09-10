package repositories

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"gorm.io/gorm"
)

type Petak struct {
	db *gorm.DB
}

func NewPetakRepository(db *gorm.DB) *Petak {
	return &Petak{db}
}

func (p *Petak) FindAll(idPetak string) ([]model.Petak, error) {
	var petaks []model.Petak
	if err := p.db.Find(&petaks, "id_petak = ?", idPetak).Error; err != nil {
		return nil, err
	}
	return petaks, nil
}

func (p *Petak) FindByID(id string) (*model.Petak, error) {
	var petak model.Petak
	if err := p.db.First(&petak, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &petak, nil
}

func (p *Petak) Create(petak *model.PetakInput) error {
	if err := p.db.Create(petak.ToPetak()).Error; err != nil {
		return err
	}
	return nil
}

func (p *Petak) Update(id string, petak *model.PetakInput) error {
	if err := p.db.Model(&model.Petak{}).Where("id = ?", id).Updates(petak.ToPetak()).Error; err != nil {
		return err
	}
	return nil
}

func (p *Petak) Delete(id string) error {
	if err := p.db.Delete(&model.Petak{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
