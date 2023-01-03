package file

import (
	"bytes"
	"context"
	"go-rriaudiobook-server/internal/utils/config"
	"go-rriaudiobook-server/internal/utils/logger"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

func UploadFile(file *multipart.FileHeader, path string, fallbacktype string) (fpath string, err error) {

	files, _ := file.Open()
	defer files.Close()

	fileByte, err := io.ReadAll(files)
	contentType := http.DetectContentType(fileByte)
	fpath = path + "/" + uuid.New().String() + filepath.Ext(file.Filename)

	if contentType != fallbacktype && fallbacktype != "" {
		contentType = fallbacktype
	}

	uploader := manager.NewUploader(s3Client)
	_, err = uploader.Upload(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(config.S3_BUCKET_NAME),
		Key:         aws.String(fpath),
		Body:        bytes.NewReader(fileByte),
		ContentType: aws.String(contentType),
		ACL:         "public-read",
	})

	if err != nil {
		logger.WriteLog(err)
	}

	fpath = config.S3_PUBLIC_ACCESS + "/" + fpath
	return
}
