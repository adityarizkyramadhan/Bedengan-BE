package dto

type FindAllKavlingRequest struct {
	GroundID          string `json:"ground_id" form:"ground_id" binding:"required"`
	TanggalKedatangan string `json:"tanggal_kedatangan" form:"tanggal_kedatangan" binding:"required"`
	TanggalKepulangan string `json:"tanggal_kepulangan" form:"tanggal_kepulangan" binding:"required"`
}
