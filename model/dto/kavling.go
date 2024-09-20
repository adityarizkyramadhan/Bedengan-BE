package dto

type FindAllKavlingRequest struct {
	GroundID          string `json:"ground_id" form:"ground_id"`
	TanggalKedatangan string `json:"tanggal_kedatangan" form:"tanggal_kedatangan"`
	TanggalKepulangan string `json:"tanggal_kepulangan" form:"tanggal_kepulangan"`
}
