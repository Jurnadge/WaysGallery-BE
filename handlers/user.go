package handlers

import (
	"context"
	"io"
	"io/ioutil"
	dto "juna/dto/result"
	usersdto "juna/dto/users"
	"juna/repositories"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

func (h *handlerUser) GetUser(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: user})
}

func (h *handlerUser) UpdateProfile(c echo.Context) error {
	var (
		avatarFile, artFile *multipart.FileHeader
		avatarPath, artPath string
	)

	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// Get files from form
	files := form.File["image"]
	if len(files) > 0 {
		avatarFile = files[0]
	}
	files = form.File["art"]
	if len(files) > 0 {
		artFile = files[0]
	}

	// Handle avatar file upload
	if avatarFile != nil {
		avatarSrc, err := avatarFile.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		defer avatarSrc.Close()

		avatarTempFile, err := ioutil.TempFile("uploads", "avatar-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		defer avatarTempFile.Close()

		if _, err := io.Copy(avatarTempFile, avatarSrc); err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		avatarPath = avatarTempFile.Name()
	}

	// Handle art file upload
	if artFile != nil {
		artSrc, err := artFile.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		defer artSrc.Close()

		artTempFile, err := ioutil.TempFile("uploads", "art-*.png")
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		defer artTempFile.Close()

		if _, err := io.Copy(artTempFile, artSrc); err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}

		artPath = artTempFile.Name()
	}

	// Upload files to Cloudinary
	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	var avatarURL string
	var artURL string
	if avatarPath != "" {
		avatarResp, err := cld.Upload.Upload(ctx, avatarPath, uploader.UploadParams{Folder: "WaysGallery_Profile_picture"})
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		avatarURL = avatarResp.SecureURL
	}

	if artPath != "" {
		artResp, err := cld.Upload.Upload(ctx, artPath, uploader.UploadParams{Folder: "WaysGallery_Art"})
		if err != nil {
			return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		}
		artURL = artResp.SecureURL
	}

	request := usersdto.UpdateUserRequest{
		Fullname: c.FormValue("fullname"),
		Greeting: c.FormValue("greeting"),
		Avatar:   avatarURL,
		Art:      artURL,
	}

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.UserRepository.GetUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Fullname != "" {
		user.Fullname = request.Fullname
	}
	if request.Greeting != "" {
		user.Greeting = request.Greeting
	}
	if request.Avatar != "" {
		user.Avatar = request.Avatar
	}
	if request.Art != "" {
		user.Art = request.Art
	}

	data, err := h.UserRepository.UpdateProfile(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: data})
}

func (h *handlerUser) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: users})
}

func (h *handlerUser) GetUserDetailByLogin(c echo.Context) error {
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	user, err := h.UserRepository.GetUserDetailByLogin(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	posts, err := h.UserRepository.FindPostByUserId(int(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: usersdto.UserDetailResponse{
		ID:       int(user.ID),
		Fullname: user.Fullname,
		Email:    user.Email,
		Greeting: user.Greeting,
		Avatar:   user.Avatar,
		Post:     posts,
		Arts:     user.Art,
	}})
}

func (h *handlerUser) GetUserDetailById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUserDetailById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	posts, err := h.UserRepository.FindPostByUserId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccesResult{Code: http.StatusOK, Data: usersdto.UserDetailResponse{
		ID:       int(user.ID),
		Fullname: user.Fullname,
		Email:    user.Email,
		Greeting: user.Greeting,
		Avatar:   user.Avatar,
		Post:     posts,
		Arts:     user.Art,
	}})
}
