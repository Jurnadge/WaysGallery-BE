package handlers

import (
	hireddto "juna/dto/hired"
	dto "juna/dto/result"
	"juna/models"
	"juna/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type handlerHired struct {
	HiredRepository repositories.HiredRepository
}

func HandlerHired(HiredRepository repositories.HiredRepository) *handlerHired {
	return &handlerHired{HiredRepository}
}

func (h *handlerHired) CreateHired(c echo.Context) error {
	request := new(hireddto.CreateHiredRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request.Status = "pending"
	price, _ := strconv.Atoi(request.Price)
	orderTo, _ := strconv.Atoi(request.OrderTo)
	startProject, _ := time.Parse("2006-01-02", request.StartProject)
	endProject, _ := time.Parse("2006-01-02", request.EndProject)

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	hired := models.Hired{
		Title:        request.Title,
		Description:  request.Description,
		StartProject: startProject,
		EndProject:   endProject,
		Price:        price,
		OrderBy:      int(userId),
		OrderTo:      orderTo,
		Status:       request.Status,
	}

	dataTransaction, err := h.HiredRepository.CreateHired(hired)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: dataTransaction})
}
