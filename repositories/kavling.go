package repositories

import (
	"fmt"

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
	var reservedGrounds []model.Ground
	var unreservedGrounds []model.Ground
	var grounds []model.Ground

	// Query pertama: untuk kavlings yang sudah direservasi
	err1 := p.db.Preload("SubGrounds.Kavlings", func(db *gorm.DB) *gorm.DB {
		db = db.Joins("JOIN reservasis ON reservasis.kavling_id = kavlings.id").
			Joins("JOIN invoice_reservasis ON invoice_reservasis.id = reservasis.invoice_reservasi_id")

		if req.TanggalKedatangan != "" && req.TanggalKepulangan != "" {
			db = db.Where(
				"invoice_reservasis.tanggal_kedatangan <= ? AND invoice_reservasis.tanggal_kepulangan >= ?",
				req.TanggalKepulangan, req.TanggalKedatangan,
			)
		}

		return db.Order("kavlings.kolom ASC")
	}).Find(&reservedGrounds).Error

	// Query kedua: untuk kavlings yang belum direservasi
	err2 := p.db.Preload("SubGrounds.Kavlings", func(db *gorm.DB) *gorm.DB {
		db = db.Joins("LEFT JOIN reservasis ON reservasis.kavling_id = kavlings.id").
			Where("reservasis.id IS NULL")

		return db.Order("kavlings.kolom ASC")
	}).Find(&unreservedGrounds).Error

	response := make(map[string]map[string][][]map[string]interface{})
	// Cek jika ada error pada query
	if err1 != nil || err2 != nil {
		return response, fmt.Errorf("error loading grounds: %v %v", err1, err2)
	}

	grounds = append(reservedGrounds, unreservedGrounds...)

	for _, ground := range grounds {
		subGroundMap := make(map[string][][]map[string]interface{})

		for _, subGround := range ground.SubGrounds {
			kavlingList := make([][]map[string]interface{}, 0)

			// Group by baris
			kavlingByBaris := map[int][]map[string]interface{}{}
			for _, kavling := range subGround.Kavlings {
				// Tentukan apakah kavling aktif berdasarkan invoice reservasi
				isAktif := len(kavling.Reservasi) == 0

				kavlingData := map[string]interface{}{
					"kolom":         kavling.Kolom,
					"baris":         kavling.Baris,
					"id":            kavling.ID,
					"ground":        ground.Nama,
					"nomorGround":   subGround.Nama,
					"nomorKavling":  kavling.Nama,
					"harga":         kavling.Harga,
					"isAvailable":   isAktif, // Ubah sesuai status aktif
					"sub_ground_id": kavling.SubGroundID,
					"groud_id":      ground.ID,
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
