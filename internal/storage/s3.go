package storage

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/omegion/db-backup/internal/backup"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3 provider backup storage.
type S3 struct {
	Bucket      string
	EndpointURL string
}

// Get returns backup with downloaded backup from S3.
func (s *S3) Get(backup backup.Backup) error {
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
		return fmt.Errorf("failed to get backup file: %w", err)
	}

	defer file.Close()

	_, err = downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(backup.Path),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to download backup file: %w", err)
	}

	return nil
}

// Save saves provider backup to S3.
func (s *S3) Save(backup backup.Backup) error {
	file, err := os.Open(backup.Path)
	if err != nil {
		return err
	}

	defer os.Remove(file.Name())

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	buffer := make([]byte, stat.Size())

	_, err = file.Read(buffer)
	if err != nil {
		return err
	}

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
		Bucket:  aws.String(s.Bucket),
		Body:    bytes.NewReader(buffer),
		Key:     aws.String(path),
		Tagging: aws.String("Backup=True"),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Dump is successful for %s\n", backup.Name)

	return nil
}

// Delete removes backup from S3.
func (s *S3) Delete(backup backup.Backup) error { return nil }

// List lists backups from S3.
func (s *S3) List(b backup.Backup) ([]backup.Backup, error) {
	config := aws.Config{}

	if s.EndpointURL != "" {
		config.Endpoint = aws.String(s.EndpointURL)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: config,
	}))

	svc := s3.New(sess)

	res, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:    aws.String(s.Bucket),
		Delimiter: nil,
		Prefix:    aws.String(fmt.Sprintf("%s/%s", b.Host, b.Name)),
	})
	if err != nil {
		return []backup.Backup{}, err
	}

	//nolint:prealloc // this is enough for now.
	var backups []backup.Backup

	for _, object := range res.Contents {
		backups = append(backups, backup.Backup{
			Name: b.Name,
			Path: *object.Key,
			Host: b.Host,
		})
	}

	return backups, nil
}
