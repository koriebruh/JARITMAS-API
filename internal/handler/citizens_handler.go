package handler

import (
	"fmt"
	"github.com/firstudio-lab/JARITMAS-API/internal/dto"
	"github.com/firstudio-lab/JARITMAS-API/internal/usecase"
	"github.com/firstudio-lab/JARITMAS-API/pkg/helper"
	"github.com/firstudio-lab/JARITMAS-API/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type CitizenHandler interface {
	FindCitizenByNIK(ctx *fiber.Ctx) error
	FindCitizenPage(ctx *fiber.Ctx) error
	CreateCitizen(ctx *fiber.Ctx) error
	UpdateCitizenByNIK(ctx *fiber.Ctx) error
	DeleteCitizenByNIK(ctx *fiber.Ctx) error
}

type CitizensHandlerImpl struct {
	usecase.CitizensUsecase
}

func NewCitizensHandler(citizensUsecase usecase.CitizensUsecase) *CitizensHandlerImpl {
	return &CitizensHandlerImpl{CitizensUsecase: citizensUsecase}
}

func (h CitizensHandlerImpl) FindCitizenByNIK(ctx *fiber.Ctx) error {
	params := ctx.Params("nik")
	atoi, err := strconv.Atoi(params)
	if err != nil {
		err := fmt.Errorf("%d:%v", http.StatusBadRequest, "Nik is not suitable")
		return helper.WResponses(ctx, err, "", nil)
	}

	Citizen, err := h.CitizensUsecase.FindCitizenByNIK(ctx.Context(), int64(atoi))
	if err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusOK).JSON(helper.UseData{
		Status:  "OK",
		Message: "successfully get",
		Data:    Citizen,
	})
}

func (h CitizensHandlerImpl) FindCitizenPage(ctx *fiber.Ctx) error {
	query := ctx.Query("page")
	if query == "" {
		err := fmt.Errorf("%d:%v", http.StatusBadRequest, "bad value cant get data")
		return helper.WResponses(ctx, err, "", nil)
	}

	atoi, err := strconv.Atoi(query)
	if err != nil {
		err := fmt.Errorf("%d:%v", http.StatusBadRequest, "bad value cant get data")
		return helper.WResponses(ctx, err, "", nil)
	}

	Citizens, err := h.CitizensUsecase.FindCitizenPage(ctx.Context(), atoi)
	if err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusOK).JSON(helper.UseData{
		Status:  "OK",
		Message: "successfully get 10/page",
		Data:    Citizens,
	})
}

func (h CitizensHandlerImpl) CreateCitizen(ctx *fiber.Ctx) error {
	var body dto.CitizenReqCreate
	if err := ctx.BodyParser(&body); err != nil {
		logger.Log.Errorf("Fail to parse body %e", err)
		err := fmt.Errorf("%d:%v", http.StatusInternalServerError, " failed to parse json")
		return helper.WResponses(ctx, err, "", nil)
	}

	if err := h.CitizensUsecase.CreateCitizen(ctx.Context(), body); err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusCreated).JSON(helper.NoData{
		Status:  "CREATED",
		Message: "successfully created new Citizen",
	})
}

func (h CitizensHandlerImpl) UpdateCitizenByNIK(ctx *fiber.Ctx) error {
	var body dto.CitizenReqUpdate
	params := ctx.Params("nik")
	atoi, err := strconv.Atoi(params)
	if err != nil {
		err := fmt.Errorf("%d:%v", http.StatusBadRequest, "Nik is not suitable")
		return helper.WResponses(ctx, err, "", nil)
	}
	if err := ctx.BodyParser(&body); err != nil {
		logger.Log.Errorf("Fail to parse body %e", err)
		err := fmt.Errorf("%d:%v", http.StatusInternalServerError, " failed to parse json")
		return helper.WResponses(ctx, err, "", nil)
	}

	if err := h.CitizensUsecase.UpdateCitizenByNIK(ctx.Context(), int64(atoi), body); err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusCreated).JSON(helper.NoData{
		Status:  "OK",
		Message: "successfully updated Citizen",
	})
}

func (h CitizensHandlerImpl) DeleteCitizenByNIK(ctx *fiber.Ctx) error {
	params := ctx.Params("nik")
	atoi, err := strconv.Atoi(params)
	if err != nil {
		err := fmt.Errorf("%d:%v", http.StatusBadRequest, "Nik is not suitable")
		return helper.WResponses(ctx, err, "", nil)
	}

	if err := h.CitizensUsecase.DeleteCitizenByNIK(ctx.Context(), int64(atoi)); err != nil {
		return helper.WResponses(ctx, err, "", nil)
	}

	return ctx.Status(http.StatusCreated).JSON(helper.NoData{
		Status:  "OK",
		Message: "DELETE Citizen successfully",
	})
}
