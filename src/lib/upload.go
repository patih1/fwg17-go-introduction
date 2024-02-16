package lib

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context, dest string) (string, error) {
	file, err := c.FormFile("picture")
	if err != nil {
		return "no file", errors.New("no file")
	}

	ext := map[string]string{
		"image/png":  ".png",
		"image/jpeg": ".jpeg",
		"image/jpg":  ".jpg",
	}

	fileType := file.Header["Content-Type"][0]
	extention := ext[fileType]
	fmt.Println(extention)

	if extention == "" {
		return "wrong ext", errors.New("wrong ext")
	}

	fileName := fmt.Sprintf("uploads/%v/%v%v", dest, uuid.NewString(), extention)
	c.SaveUploadedFile(file, fileName)
	return fileName, nil
}
