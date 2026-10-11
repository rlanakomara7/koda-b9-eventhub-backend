package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveUploadedFile(
	c *gin.Context,
	fieldName string,
) (string, error) {

	file, err := c.FormFile(fieldName)

	// tidak upload file
	if errors.Is(err, http.ErrMissingFile) {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	// maksimal 2 MB
	if file.Size > 2*1024*1024 {
		return "", errors.New(
			"file size maximum is 2 MB",
		)
	}

	extension := strings.ToLower(
		filepath.Ext(file.Filename),
	)

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	if !allowedExtensions[extension] {
		return "", errors.New(
			"file must be jpg, jpeg, png, or webp",
		)
	}

	err = os.MkdirAll(
		"uploads",
		0755,
	)

	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf(
		"%d%s",
		time.Now().UnixNano(),
		extension,
	)

	filePath := filepath.Join(
		"uploads",
		fileName,
	)

	err = c.SaveUploadedFile(
		file,
		filePath,
	)

	if err != nil {
		return "", err
	}

	return "/uploads/" + fileName, nil
}
