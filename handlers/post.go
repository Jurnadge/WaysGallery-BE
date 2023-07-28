package handlers

import (
	postdto "juna/dto/post"
	dto "juna/dto/result"
	"juna/models"
	"juna/pkg/middleware"
	"juna/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type handlerPost struct {
	PostRepository repositories.PostRepository
}

func HandlerPost(PostRepository repositories.PostRepository) *handlerPost {
	return &handlerPost{PostRepository}
}

func (h *handlerPost) FindPosts(c echo.Context) error {
	posts, err := h.PostRepository.FindPosts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: postdto.PostResponse{Post: posts}})
}

func (h *handlerPost) GetPost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var post models.Post
	post, err := h.PostRepository.GetPost(int(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: postdto.PostResponse{Post: post}})
}

func (h *handlerPost) CreatePost(c echo.Context) error {
	request := postdto.CreatePostRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	var postImage []models.PostImage
	dataContex := c.Get("dataFiles")
	filename := dataContex.([]middleware.ImageResult)
	postImage = make([]models.PostImage, len(filename))
	for i, value := range filename {
		postImage[i] = models.PostImage{ID: value.PublicID, Image: value.SecureURL}
	}

	post := models.Post{
		Title:       request.Title,
		Description: request.Description,
		CreatedBy:   int(userId),
		PostImage:   postImage,
	}

	data, err := h.PostRepository.CreatePost(post)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	dataInserted, err := h.PostRepository.GetPost(int(data.ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: postdto.PostResponse{Post: dataInserted}})
}
