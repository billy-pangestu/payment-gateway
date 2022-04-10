package handler

import (
	"net/http"
	"payment-gateway-backend/helper"
	"payment-gateway-backend/model"
	"payment-gateway-backend/pkg/str"
	"payment-gateway-backend/usecase"
	"strings"
)

// FileHandler ...
type FileHandler struct {
	Handler
}

// UploadHandler ...
func (h *FileHandler) UploadHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	// Check file size
	maxUploadSize := str.StringToInt(h.Handler.EnvConfig["FILE_MAX_UPLOAD_SIZE"])
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxUploadSize))
	err := r.ParseMultipartForm(int64(maxUploadSize))
	if err != nil {
		SendBadRequest(w, helper.FileTooBig)
		return
	}

	// Read file type
	fileType := r.PostFormValue("type")
	if !str.Contains(model.FileWhitelist, fileType) {
		SendBadRequest(w, helper.InvalidFileType)
		return
	}

	// Read file
	file, header, err := r.FormFile("file")
	if err != nil {
		SendBadRequest(w, helper.FileError)
		return
	}
	defer file.Close()

	// Upload file to local temporary
	fileUc := usecase.FileUC{ContractUC: h.ContractUC}
	res, err := fileUc.Upload(fileType, userID, strings.Replace(header.Filename, " ", "", -1), header)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}
