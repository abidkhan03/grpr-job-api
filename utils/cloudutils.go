package utils

import (
	"context"
	"encoding/csv"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"github.com/grpr-job-api/utils/constants"
	"google.golang.org/api/option"
)

func uploadToGoogleCloud(filePath string, uploadFilePath string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		return err
	}

	defer client.Close()

	bucket := client.Bucket(constants.CloudConfig.GCloudBucketName)
	object := bucket.Object(uploadFilePath)
	writer := object.NewWriter(ctx)

	defer writer.Close()

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	if err != nil {
		return err
	}

	return nil
}

func readCsvFromGoogleCloud(filePath string) (*csv.Reader, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		return nil, err
	}

	defer client.Close()

	bucket := client.Bucket(constants.CloudConfig.GCloudBucketName)
	object := bucket.Object(filePath)
	reader, err := object.NewReader(ctx)
	if err != nil {
		return nil, err
	}

	return csv.NewReader(reader), nil
}
