package file

import (
	"fmt"
	"go-rriaudiobook-server/internal/utils/errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

func ValidateFileType(file *multipart.FileHeader, types string) (err error) {
	files, _ := file.Open()
	defer files.Close()

	fileByte, err := io.ReadAll(files)
	contentType := http.DetectContentType(fileByte)

	if strings.Contains(types, "/*") {
		types = strings.TrimSuffix(types, "/*")
	}

	if strings.Contains(contentType, types) {
		return
	}

	err = errors.New(203, fmt.Sprintf("file type must %s", types))
	return
}
