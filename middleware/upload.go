package middleware

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SavedFile struct {
	FileName string
	FilePath string
	FileUrl  string
	Type     string
}

func UploadFile(c *gin.Context, param string) (*SavedFile, error) {
	paramFile := param
	if param == "" {
		paramFile = "images"
	}
	file, errFile := c.FormFile(paramFile)

	if file == nil {
		return &SavedFile{}, fmt.Errorf("empty File: %s is empty", param)
	}

	if errFile != nil {
		return &SavedFile{}, errFile
	}

	hType := file.Header.Get("Content-Type")
	fmt.Println(hType)
	if !strings.HasPrefix(hType, "image/") {
		return &SavedFile{}, errors.New("only image file allowed")
	}

	folderPath := "./images/"
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.MkdirAll(folderPath, 0755)
	}

	extFile := strings.Split(file.Filename, ".")
	randFileName := uuid.New()
	file.Filename = randFileName.String() + "image" + "." + extFile[len(extFile)-1]

	path := folderPath + file.Filename

	host := "http://" + c.Request.Host

	url := host + "/api/v1/images/" + file.Filename

	if err := c.SaveUploadedFile(file, path); err != nil {
		return &SavedFile{}, err
	}

	return &SavedFile{
		FileName: file.Filename,
		FilePath: path,
		FileUrl:  url,
		Type:     file.Header.Get("Content-Type"),
	}, nil
}
