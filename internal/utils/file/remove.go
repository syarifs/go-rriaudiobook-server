package file

import (
	"context"
	"go-rriaudiobook-server/internal/utils/config"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func RemoveFile(path string) (err error) {
	path = strings.TrimPrefix(path, config.S3_PUBLIC_ACCESS)
	path = strings.TrimLeft(path, "/")
	_, err = s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(config.S3_BUCKET_NAME),
		Key:    aws.String(path),
	})
	return
}
