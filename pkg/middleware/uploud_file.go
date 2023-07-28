package middleware

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		var dataFiles []string
		for _, fileHeaders := range form.File {
			for _, fileHeader := range fileHeaders {
				src, err := fileHeader.Open()
				if err != nil {
					return c.JSON(http.StatusBadRequest, err)
				}
				defer src.Close()

				tempFile, err := ioutil.TempFile("uploads", "image-*.png")
				if err != nil {
					return c.JSON(http.StatusBadRequest, err)
				}
				defer tempFile.Close()

				if _, err := io.Copy(tempFile, src); err != nil {
					return c.JSON(http.StatusBadRequest, err)
				}

				dataFiles = append(dataFiles, tempFile.Name())
			}
		}

		c.Set("dataFiles", dataFiles)
		return next(c)
	}
}
