package repositories

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"gorm.io/gorm"
)

type Kavling struct {
	db *gorm.DB
}

func NewKavlingRepository(db *gorm.DB) *Kavling {
	return &Kavling{db}
}

func (p *Kavling) FindAll(idKavling string) ([]model.Kavling, error) {
	var Kavlings []model.Kavling
	if err := p.db.Find(&Kavlings, "id_Kavling = ?", idKavling).Error; err != nil {
		return nil, err
	}
	return Kavlings, nil
}

func (p *Kavling) FindByID(id string) (*model.Kavling, error) {
	var Kavling model.Kavling
	if err := p.db.First(&Kavling, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &Kavling, nil
}

func (p *Kavling) Create(Kavling *model.KavlingInput) error {
	if err := p.db.Create(Kavling.ToKavling()).Error; err != nil {
		return err
	}
	return nil
}

func (p *Kavling) Update(id string, Kavling *model.KavlingInput) error {
	if err := p.db.Model(&model.Kavling{}).Where("id = ?", id).Updates(Kavling.ToKavling()).Error; err != nil {
		return err
	}
	return nil
}

func (p *Kavling) Delete(id string) error {
	if err := p.db.Delete(&model.Kavling{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
