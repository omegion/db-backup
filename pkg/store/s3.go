package store

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/omegion/go-db-backup/pkg/exporter"
	"os"
	"path/filepath"
)

type S3 struct {
	Bucket      string
	EndpointURL string
}

func (s *S3) Store(backup *exporter.ExportResult) error {
	file, err := os.Open(backup.Path)
	if err != nil {
		return err
	}

	defer func() {
		os.Remove(file.Name())
	}()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	buffer := make([]byte, stat.Size())
	file.Read(buffer)

	config := aws.Config{}

	if s.EndpointURL != "" {
		config.Endpoint = aws.String(s.EndpointURL)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: config,
	}))

	svc := s3.New(sess)

	path := filepath.Join(backup.DatabaseName, backup.Path)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:       aws.String(s.Bucket),
		Body:         bytes.NewReader(buffer),
		Key:          aws.String(path),
		StorageClass: aws.String("GLACIER"),
	})
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Dump is successful for %s", backup.DatabaseName))

	return nil
}
