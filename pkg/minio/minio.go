package minio

import (
	"errors"
	"mime/multipart"
	"net/url"
	"payment-gateway-backend/usecase/viewmodel"
	"time"

	"github.com/minio/minio-go/v6"
)

var (
	urlDuration = time.Second * 30 * 60
)

// IMinio ...
type IMinio interface {
	CreateBucket(bucketName, bucketLocation string) (string, error)
	Upload(bucketName, path string, fileHeader *multipart.FileHeader) (viewmodel.MinioVM, error)
	GetFileURL(bucketName, objectName string) (string, error)
	Delete(bucketName, objectName string) (err error)
}

// Entity ....
type Entity struct {
	Name string
	Size int64
	Type string
}

// minioModel ...
type minioModel struct {
	Client *minio.Client
}

// NewMinioModel ...
func NewMinioModel(client *minio.Client) IMinio {
	return &minioModel{Client: client}
}

// CreateBucket ...
func (model minioModel) CreateBucket(bucketName, bucketLocation string) (data string, err error) {
	err = model.Client.MakeBucket(bucketName, bucketLocation)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := model.Client.BucketExists(bucketName)
		if errBucketExists == nil && exists {
			return data, errors.New("We already own " + bucketName)
		}

		return data, err
	}

	return "Successfully created " + bucketName, err
}

// Upload ...
func (model minioModel) Upload(bucketName, path string, fileHeader *multipart.FileHeader) (res viewmodel.MinioVM, err error) {
	src, err := fileHeader.Open()
	if err != nil {
		return res, err
	}
	defer src.Close()

	res.Name = path
	res.Type = fileHeader.Header.Get("Content-Type")
	res.Size = fileHeader.Size
	_, err = model.Client.PutObject(bucketName, res.Name, src, res.Size, minio.PutObjectOptions{ContentType: res.Type})
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetFileURL ...
func (model minioModel) GetFileURL(bucketName, objectName string) (data string, err error) {
	reqParams := make(url.Values)
	reqParams.Set(`response-content-disposition`, `attachment; filename=\"`+objectName+`\"`)

	// Generates a presigned url which expires in a day.
	url, err := model.Client.PresignedGetObject(bucketName, objectName, urlDuration, reqParams)
	if err != nil {
		return data, err
	}
	data = url.String()
	return data, err
}

// Delete ...
func (model minioModel) Delete(bucketName, objectName string) (err error) {
	err = model.Client.RemoveObject(bucketName, objectName)

	return err
}
