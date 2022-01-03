package repository

import (
	"context"
	"fmt"
	"github.com/binishmaharjan/memrizr/account/model/apperrors"
	"io"
	"log"
	"mime/multipart"

	"cloud.google.com/go/storage"
	"github.com/binishmaharjan/memrizr/account/model"
)

type gcImageRepository struct {
	Storage    *storage.Client
	BucketName string
}

// NewImageRepository is a factory for initializing User Repositories
func NewImageRepository(gcClient *storage.Client, bucketName string) model.ImageRepository {
	return &gcImageRepository{
		Storage:    gcClient,
		BucketName: bucketName,
	}
}

func (r *gcImageRepository) UpdateProfile(ctx context.Context, objName string, imageFile multipart.File) (string, error) {
	bckt := r.Storage.Bucket(r.BucketName)

	object := bckt.Object(objName)
	wc := object.NewWriter(ctx)

	wc.ObjectAttrs.CacheControl = "Cache-Control:no-cache, max-age=0"

	if _, err := io.Copy(wc, imageFile); err != nil {
		log.Printf("Unable to write file to Google Cloud Storage: %v \n", err)
		return "", apperrors.NewInternal()
	}

	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	imageURL := fmt.Sprintf("https:storage.googleapis.com/%s/%s", r.BucketName, objName)

	return imageURL, nil
}
