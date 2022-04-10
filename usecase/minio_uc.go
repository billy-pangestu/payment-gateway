package usecase

import (
	"mime/multipart"
	"payment-gateway-backend/pkg/logruslogger"
	"payment-gateway-backend/pkg/minio"
	"payment-gateway-backend/usecase/viewmodel"
	"strings"
)

// MinioUC ...
type MinioUC struct {
	*ContractUC
}

// CreateBucket create bucket function
func (uc MinioUC) CreateBucket(bucketName, bucketLocation string) (res string, err error) {
	ctx := "MinioUC.CreateBucket"

	minioModel := minio.NewMinioModel(uc.Minio)
	res, err = minioModel.CreateBucket(bucketName, bucketLocation)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "create_bucket", uc.ReqID)
		return res, err
	}
	return res, err
}

// Upload upload file to min.io server
func (uc MinioUC) Upload(path string, file *multipart.FileHeader) (res viewmodel.MinioVM, err error) {
	ctx := "MinioUC.Upload"

	defaultBucket := uc.ContractUC.EnvConfig["MINIO_DEFAULT_BUCKET"]
	minioModel := minio.NewMinioModel(uc.Minio)
	res, err = minioModel.Upload(defaultBucket, path, file)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "upload", uc.ReqID)
		return res, err
	}

	return res, err
}

// GetFileURL get file url by object name
func (uc MinioUC) GetFileURL(objectName string) (res string, err error) {
	ctx := "MinioUC.GetFileURL"

	defaultBucket := uc.ContractUC.EnvConfig["MINIO_DEFAULT_BUCKET"]
	minioModel := minio.NewMinioModel(uc.Minio)
	if objectName == "" {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "empty_parameter", uc.ReqID)
		return res, err
	}

	res, err = minioModel.GetFileURL(defaultBucket, objectName)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_file_url", uc.ReqID)
		return res, err
	}

	res = strings.Replace(res, "http://"+uc.ContractUC.EnvConfig["MINIO_ENDPOINT"], uc.ContractUC.EnvConfig["MINIO_BASE_URL"], 1)

	return res, err
}

// GetFileURLNoErr get file url by object name wo err response
func (uc MinioUC) GetFileURLNoErr(objectName string) (res string) {
	defaultBucket := uc.ContractUC.EnvConfig["MINIO_DEFAULT_BUCKET"]
	minioModel := minio.NewMinioModel(uc.Minio)
	if objectName != "" {
		res, _ = minioModel.GetFileURL(defaultBucket, objectName)
	}

	return res
}

// Delete delete object
func (uc MinioUC) Delete(objectName string) (err error) {
	ctx := "MinioUC.Delete"

	defaultBucket := uc.ContractUC.EnvConfig["MINIO_DEFAULT_BUCKET"]
	minioModel := minio.NewMinioModel(uc.Minio)
	err = minioModel.Delete(defaultBucket, objectName)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete", uc.ReqID)
		return err
	}

	return err
}
