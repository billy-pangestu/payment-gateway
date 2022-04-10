package usecase

import (
	"errors"
	"mime/multipart"
	"payment-gateway-backend/model"
	"payment-gateway-backend/pkg/logruslogger"
	"payment-gateway-backend/usecase/viewmodel"
	"time"

	uuid "github.com/satori/go.uuid"
)

// FileUC ...
type FileUC struct {
	*ContractUC
}

// FindOne ...
func (uc FileUC) FindOne(id string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.FindOne"

	fileModel := model.NewFileModel(uc.DB)
	data, err := fileModel.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_id", uc.ReqID)
		return res, err
	}

	mininoUc := MinioUC{ContractUC: uc.ContractUC}
	tempURL, err := mininoUc.GetFileURL(data.URL.String)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_url", uc.ReqID)
		return res, err
	}

	res = viewmodel.FileVM{
		ID:         data.ID,
		Type:       data.Type.String,
		URL:        data.URL.String,
		TempURL:    tempURL,
		UserUpload: data.UserUpload.String,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt.String,
	}

	return res, err
}

// FindOneUnassigned check if image is unassigned
func (uc FileUC) FindOneUnassigned(id, types, userUpload string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.FindOneUnassigned"

	fileModel := model.NewFileModel(uc.DB)
	data, err := fileModel.FindUnassignedByID(id, types, userUpload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_id", uc.ReqID)
		return res, err
	}

	mininoUc := MinioUC{ContractUC: uc.ContractUC}
	tempURL, err := mininoUc.GetFileURL(data.URL.String)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_url", uc.ReqID)
		return res, err
	}

	res = viewmodel.FileVM{
		ID:         data.ID,
		Type:       data.Type.String,
		URL:        data.URL.String,
		TempURL:    tempURL,
		UserUpload: data.UserUpload.String,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt.String,
	}

	return res, err
}

// FindOneAssigned check if image is assigned
func (uc FileUC) FindOneAssigned(id, types string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.FindOneAssigned"

	fileModel := model.NewFileModel(uc.DB)
	data, err := fileModel.FindAssignedByID(id, types)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_id", uc.ReqID)
		return res, errors.New("Invalid " + types + " file")
	}

	mininoUc := MinioUC{ContractUC: uc.ContractUC}
	tempURL, err := mininoUc.GetFileURL(data.URL.String)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_url", uc.ReqID)
		return res, err
	}

	res = viewmodel.FileVM{
		ID:         data.ID,
		Type:       data.Type.String,
		URL:        data.URL.String,
		TempURL:    tempURL,
		UserUpload: data.UserUpload.String,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
		DeletedAt:  data.DeletedAt.String,
	}

	return res, err
}

// Create ...
func (uc FileUC) Create(types, url, userUpload string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.Create"

	fileModel := model.NewFileModel(uc.DB)

	// Delete all unused files first
	err = uc.DeleteAllUnused(userUpload, types)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete_unused", uc.ReqID)
		return res, err
	}

	now := time.Now().UTC()
	res = viewmodel.FileVM{
		Type:       types,
		URL:        url,
		UserUpload: userUpload,
		CreatedAt:  now.Format(time.RFC3339),
		UpdatedAt:  now.Format(time.RFC3339),
	}
	res.ID, err = fileModel.Store(res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "create", uc.ReqID)
		return res, err
	}

	// Get temp url
	mininoUc := MinioUC{ContractUC: uc.ContractUC}
	res.TempURL, err = mininoUc.GetFileURL(res.URL)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "get_url", uc.ReqID)
		return res, err
	}

	return res, err
}

// Delete ...
func (uc FileUC) Delete(id string) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.Delete"

	now := time.Now().UTC()
	fileModel := model.NewFileModel(uc.DB)
	res.ID, err = fileModel.Destroy(id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "delete", uc.ReqID)
		return res, err
	}

	return res, err
}

// DeleteAllUnused ...
func (uc FileUC) DeleteAllUnused(userUpload, types string) (err error) {
	ctx := "FileUC.DeleteAllUnused"
	fileModel := model.NewFileModel(uc.DB)
	unusedFile, err := fileModel.FindAllUnassignedByUserID(userUpload, types)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_all_unused", uc.ReqID)
		return err
	}

	minioUc := MinioUC{ContractUC: uc.ContractUC}
	now := time.Now().UTC()
	for _, r := range unusedFile {
		err = minioUc.Delete(r.URL.String)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "minio", uc.ReqID)
		}

		_, err = fileModel.Destroy(r.ID, now)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_file", uc.ReqID)
		}
	}
	err = nil

	return err
}

// Upload ...
func (uc FileUC) Upload(types, userUpload, fileName string, file *multipart.FileHeader) (res viewmodel.FileVM, err error) {
	ctx := "FileUC.Upload"

	minioUc := MinioUC{ContractUC: uc.ContractUC}
	fileName = types + "/" + userUpload + "/" + uuid.NewV4().String() + fileName
	fileUpload, err := minioUc.Upload(fileName, file)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "upload_file", uc.ReqID)
		return res, err
	}

	res, err = uc.Create(types, fileUpload.Name, userUpload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "create", uc.ReqID)
		return res, err
	}

	return res, err
}
