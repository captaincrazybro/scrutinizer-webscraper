package sw

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	lu "github.com/captaincrazybro/literalutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	bucketName string = "sw-ignored-list"
)

func DownloadIgnoredFile() error {
	key := os.Getenv(S3ObjectKeyEnv)
	bucket := os.Getenv(S3BucketEnv)
	if key == "" || bucket == "" {
		return fmt.Errorf("%s environment variable has not been set", lu.STernary(key == "", S3ObjectKeyEnv, S3BucketEnv))
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us"),
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		return err
	}

	// creates s3 service client
	svc := s3.New(sess)

	// creates the downloader
	downloader := s3manager.NewDownloaderWithClient(svc)

	// downloads the file
	file, err := os.Open(IgnoredReposFileName)
	if err != nil {
		file, err = os.Create(IgnoredReposFileName)
		if err != nil {
			return err
		}
	}

	_, err = downloader.Download(file, &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &key,
	})
	if err != nil {
		return err
	}

	return nil

}
