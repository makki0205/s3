package s3

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	svc        *s3.S3
	BucketName string
}

func NewS3(AccessKeyId, SecretAccessKey, Region, BucketName string) (*S3, error) {
	creds := credentials.NewStaticCredentials(AccessKeyId, SecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
		return nil, err
	}
	cfg := aws.NewConfig().WithRegion(Region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)
	return &S3{svc: svc, BucketName: BucketName}, nil
}
func (f *S3) UpLoadPublic(LocalPath string, S3Path string) error {
	file, err := os.Open(LocalPath)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	params := &s3.PutObjectInput{
		Bucket:        aws.String(f.BucketName),
		Key:           aws.String(S3Path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
		ACL:           aws.String("public-read"),
	}
	_, err = f.svc.PutObject(params)
	return err
}
func (f *S3) UpLoad(LocalPath string, S3Path string) error {
	file, err := os.Open(LocalPath)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	params := &s3.PutObjectInput{
		Bucket:        aws.String(f.BucketName),
		Key:           aws.String(S3Path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	_, err = f.svc.PutObject(params)
	return err
}

func (f *S3) GetPreSignedUrl(S3Path string, expire time.Duration) (string, error) {
	req, _ := f.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(f.BucketName),
		Key:    aws.String(S3Path),
	})
	url, err := req.Presign(expire)
	if err != nil {
		return "", err
	}

	return url, nil
}
