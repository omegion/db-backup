package storage

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/omegion/go-db-backup/pkg/database"
	"os"
	"path/filepath"
)

type S3 struct {
	Bucket      string
	EndpointURL string
}

func (s *S3) Get(backup database.Backup) error {
	config := aws.Config{}

	if s.EndpointURL != "" {
		config.Endpoint = aws.String(s.EndpointURL)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: config,
	}))

	downloader := s3manager.NewDownloader(sess)

	file, err := os.Create(filepath.Join(backup.Filename()))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(backup.Path),
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3) Save(backup database.Backup) error {
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

	path := filepath.Join(backup.Host, backup.Name, backup.Path)

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Body:   bytes.NewReader(buffer),
		Key:    aws.String(path),
	})
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Dump is successful for %s", backup.Name))

	return nil
}
func (s *S3) Delete(backup database.Backup) error { return nil }
func (s *S3) List()                               {}
