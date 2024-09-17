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

func (p *Kavling) FindAll(req *dto.FindAllKavlingRequest) ([]model.Kavling, error) {
	var Kavlings []model.Kavling
	if err := p.db.Find(&Kavlings, "ground_id = ?", req.GroundID).Error; err != nil {
		return nil, err
	}
	if len(Kavlings) == 0 {
		return nil, utils.NewError(utils.ErrNotFound, "Kavling tidak ditemukan")
	}

	for i := range Kavlings {
		var reservasiCount int64
		// Tambahkan Tanggal Kedatangan dan Tanggal Kepulangan
		if err := p.db.Model(&model.Reservasi{}).
			Joins("JOIN invoice_reservasis ON invoice_reservasis.id = reservasis.invoice_reservasi_id").
			Where("reservasis.kavling_id = ?", Kavlings[i].ID).
			Where("invoice_reservasis.tanggal_kedatangan <= ?", req.TanggalKepulangan).
			Where("invoice_reservasis.tanggal_kepulangan >= ?", req.TanggalKedatangan).
			Count(&reservasiCount).Error; err != nil {
			return nil, err
		}
		if reservasiCount > 0 {
			Kavlings[i].Status = "terisi"
		}
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
