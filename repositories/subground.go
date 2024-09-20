package repositories

import (
	"gorm.io/gorm"

	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
)

func NewSubGroundRepository(db *gorm.DB) *SubGround {
	return &SubGround{db}
}

type SubGround struct {
	db *gorm.DB
}

// Create will create a new sub ground
func (p *SubGround) Create(subGround *model.SubGroundInput) (*model.SubGround, error) {
	subGroundData := subGround.ToSubGround()
	subGroundData.BeforeCreate()
	if err := p.db.Create(subGroundData).Error; err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal membuat sub ground")
	}
	return subGroundData, nil
}

// FindAll will return all sub ground
func (p *SubGround) FindAll(groundID string) ([]model.SubGround, error) {
	var subGrounds []model.SubGround
	if err := p.db.Where("ground_id = ?", groundID).Find(&subGrounds).Error; err != nil {
		return nil, err
	}
	if len(subGrounds) == 0 {
		return nil, utils.NewError(utils.ErrNotFound, "sub ground tidak ditemukan")
	}
	return subGrounds, nil
}

// FindByID will return a sub ground by id
func (p *SubGround) FindByID(id string) (*model.SubGround, error) {
	var subGround model.SubGround
	if err := p.db.First(&subGround, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &subGround, nil
}

// Update will update a sub ground by id
func (p *SubGround) Update(id string, subGround *model.SubGroundInput) (*model.SubGround, error) {
	subGroundData := subGround.ToSubGround()
	subGroundData.BeforeSave()
	subGroundData.ID = id

	if err := p.db.Save(subGroundData).Error; err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal update sub ground")
	}
	return subGroundData, nil
}

// Delete will delete a sub ground by id
func (p *SubGround) Delete(id string) error {
	if err := p.db.Delete(&model.SubGround{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
