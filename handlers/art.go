package handlers

import (
	"context"
	"fmt"
	artdto "juna/dto/art"
	dto "juna/dto/result"
	"juna/models"
	"juna/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type handlerArt struct {
	ArtRepository repositories.ArtRepository
}

func HandlerArt(ArtRepository repositories.ArtRepository) *handlerArt {
	return &handlerArt{ArtRepository}
}

func (h *handlerArt) CreateArt(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	dataFile := c.Get("dataFile").(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, dataFile, uploader.UploadParams{Folder: "WaysGallery_Art_Image"})

	if err != nil {
		fmt.Println(err.Error(), "ini error")
	}

	art := models.Art{
		ID:        resp.PublicID,
		Image:     resp.SecureURL,
		CreatedBy: int(userId),
	}

	_, err = h.ArtRepository.CreateArt(art)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.ArtRepository.GetArtByUserId(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: artdto.ArtsResponse{Arts: data}})
}

func (h *handlerArt) GetArtByUserId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	arts, err := h.ArtRepository.GetArtByUserId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: artdto.ArtsResponse{Arts: arts}})
}
