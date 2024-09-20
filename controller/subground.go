package controller

import (
	"github.com/adityarizkyramadhan/template-go-mvc/model"
	"github.com/adityarizkyramadhan/template-go-mvc/repositories"
	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
)

type (
	// SubGroundController is a struct to represent the sub ground controller
	SubGroundController struct {
		subGroundService *repositories.SubGround
	}
)

// NewSubGroundController will create a new sub ground controller
func NewSubGroundController(subGroundService *repositories.SubGround) *SubGroundController {
	return &SubGroundController{subGroundService}
}

// Create will create a new sub ground
// @Summary Create a new sub ground
// @Description Create a new sub ground
// @Tags SubGround
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param sub_ground body model.SubGroundInput true "Sub Ground"
// @Success 200 {object} model.SubGround
// @Router /sub-ground [post]
func (s *SubGroundController) Create(ctx *gin.Context) {
	var subGround model.SubGroundInput
	if err := ctx.ShouldBindJSON(&subGround); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	subGroundData, err := s.subGroundService.Create(&subGround)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, subGroundData)
}

// FindAll will return all sub ground
// @Summary Find all sub ground
// @Description Find all sub ground
// @Tags SubGround
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param ground_id query string true "Ground ID"
// @Success 200 {object} []model.SubGround
// @Router /sub-ground [get]
func (s *SubGroundController) FindAll(ctx *gin.Context) {
	groundID := ctx.Query("ground_id")
	subGrounds, err := s.subGroundService.FindAll(groundID)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, subGrounds)
}

// FindByID will return a sub ground by id
// @Summary Find a sub ground by id
// @Description Find a sub ground by id
// @Tags SubGround
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Sub Ground ID"
// @Success 200 {object} model.SubGround
// @Router /sub-ground/{id} [get]
func (s *SubGroundController) FindByID(ctx *gin.Context) {
	id := ctx.Param("id")
	subGround, err := s.subGroundService.FindByID(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, subGround)
}

// Update will update a sub ground by id
// @Summary Update a sub ground by id
// @Description Update a sub ground by id
// @Tags SubGround
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Sub Ground ID"
// @Param sub_ground body model.SubGroundInput true "Sub Ground"
// @Success 200 {object} model.SubGround
// @Router /sub-ground/{id} [put]
func (s *SubGroundController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var subGround model.SubGroundInput
	if err := ctx.ShouldBindJSON(&subGround); err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}
	subGroundData, err := s.subGroundService.Update(id, &subGround)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, subGroundData)
}

// Delete will delete a sub ground by id
// @Summary Delete a sub ground by id
// @Description Delete a sub ground by id
// @Tags SubGround
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Sub Ground ID"
// @Success 200
// @Router /sub-ground/{id} [delete]
func (s *SubGroundController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.subGroundService.Delete(id)
	if err != nil {
		_ = ctx.Error(err)
		ctx.Next()
		return
	}

	utils.SuccessResponse(ctx, 200, nil)
}
