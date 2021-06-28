package sw

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	lu "github.com/captaincrazybro/literalutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetIgnoredFile() (*bufio.Reader, error) {
	key := os.Getenv(S3ObjectKeyEnv)
	bucket := os.Getenv(S3BucketEnv)
	if key == "" || bucket == "" {
		return nil, fmt.Errorf("%s environment variable has not been set", lu.STernary(key == "", S3ObjectKeyEnv, S3BucketEnv))
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us"),
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		return nil, err
	}

	// creates s3 service client
	svc := s3.New(sess)

	// fetches the file
	rawObject, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return bufio.NewReader(rawObject.Body), nil

}
