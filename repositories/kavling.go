package repositories

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/model/dto"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"gorm.io/gorm"
)

type Kavling struct {
	db *gorm.DB
}

func NewKavlingRepository(db *gorm.DB) *Kavling {
	return &Kavling{db}
}

func (p *Kavling) FindAll(req *dto.FindAllKavlingRequest) (map[string]map[string][][]map[string]interface{}, error) {
	var grounds []model.Ground

	// Query GORM untuk mengambil data beserta relasi
	err := p.db.Preload("SubGrounds.Kavlings", func(db *gorm.DB) *gorm.DB {
		return db.Order("kolom ASC")
	}).Find(&grounds).Error
	if err != nil {
		return nil, err
	}

	response := make(map[string]map[string][][]map[string]interface{})

	for _, ground := range grounds {
		subGroundMap := make(map[string][][]map[string]interface{})

		for _, subGround := range ground.SubGrounds {
			kavlingList := make([][]map[string]interface{}, 0)

			// Group by baris
			kavlingByBaris := map[int][]map[string]interface{}{}
			for _, kavling := range subGround.Kavlings {
				kavlingData := map[string]interface{}{
					"kolom":        kavling.Kolom,
					"baris":        kavling.Baris,
					"id":           kavling.ID,
					"ground":       ground.Nama,
					"nomorGround":  subGround.Nama,
					"nomorKavling": kavling.Nama,
					"harga":        kavling.Harga,
					"isAvailable":  kavling.IsAvailable,
				}
				kavlingByBaris[kavling.Baris] = append(kavlingByBaris[kavling.Baris], kavlingData)
			}

			// Konversi baris yang dikelompokkan ke dalam bentuk array
			for _, kavlings := range kavlingByBaris {
				kavlingList = append(kavlingList, kavlings)
			}
			subGroundMap[subGround.Nama] = kavlingList
		}
		response[ground.Nama] = subGroundMap
	}
	return response, nil
}

func (p *Kavling) FindByID(id string) (*model.Kavling, error) {
	var Kavling model.Kavling
	if err := p.db.First(&Kavling, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &Kavling, nil
}

func (p *Kavling) Create(Kavling *model.KavlingInput) (*model.Kavling, error) {
	kavlingData := Kavling.ToKavling()
	kavlingData.BeforeCreate()
	if err := p.db.Create(kavlingData).Error; err != nil {
		return nil, utils.NewError(utils.ErrBadRequest, "gagal membuat Kavling")
	}
	return kavlingData, nil
}

func (p *Kavling) Update(id string, Kavling *model.KavlingInput) error {
	kavling := Kavling.ToKavling()
	kavling.BeforeSave()
	kavling.ID = id
	if err := p.db.Model(&model.Kavling{}).Where("id = ?", id).Updates(kavling).Error; err != nil {
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
