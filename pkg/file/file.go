package file

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"payment-gateway-backend/pkg/str"
	"strings"

	"github.com/rs/xid"
)

var (
	imageContentTypeWhitelist = []string{"image/gif", "image/jpeg", "image/png"}
)

// Download ...
func Download(url, uploadPath string) (filename string, err error) {
	filename = uploadPath + "/" + xid.New().String() + filepath.Ext(url)
	fmt.Println("Downloading ", url, " to ", filename)

	resp, err := http.Get(url)
	if err != nil {
		return filename, err
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return filename, err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return filename, err
}

// DownloadImage ...
func DownloadImage(url, uploadPath, name string) (filename, contentType string, err error) {
	filename = uploadPath + "/" + name

	// Get file from url
	resp, err := http.Get(url)
	if err != nil {
		return filename, contentType, err
	}
	if resp.StatusCode != 200 {
		return filename, contentType, errors.New(resp.Status)
	}
	defer resp.Body.Close()

	// Create file in local directory
	f, err := os.Create(filename)
	if err != nil {
		return filename, contentType, err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)

	// Get the content type
	f, err = os.Open(filename)
	if err != nil {
		return filename, contentType, err
	}
	contentType, err = GetFileContentType(f)
	if err != nil {
		return filename, contentType, err
	}
	defer f.Close()

	return filename, contentType, err
}

// GetFileContentType ...
func GetFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}

// OpenFile ...
func OpenFile(url, path, name string) (f *os.File, fileURL string) {
	fileURL, _, err := DownloadImage(url, path, name)
	if err != nil {
		return nil, fileURL
	}

	f, err = os.Open(fileURL)
	if err != nil {
		return nil, fileURL
	}

	return f, fileURL
}

// ToBase64 ...
func ToBase64(file *os.File) (res string) {
	// Read entire JPG into byte slice.
	reader := bufio.NewReader(file)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return res
	}

	// Encode as base64.
	res = base64.StdEncoding.EncodeToString(content)

	return res
}

// DecodeBase64 ...
func DecodeBase64(data, filePath, fileName string) (localPath string, contentType string, extention string, err error) {
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return localPath, contentType, extention, err
	}

	// Get content type and file extention
	contentType, extention, err = GetFileContentTypeAndExtention(dec)
	if err != nil {
		return localPath, contentType, extention, err
	}
	extention = "." + extention

	// Generate full local path
	localPath = filePath + "/" + fileName + extention

	f, err := os.Create(localPath)
	if err != nil {
		return localPath, contentType, extention, err
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		return localPath, contentType, extention, err
	}
	if err := f.Sync(); err != nil {
		return localPath, contentType, extention, err
	}

	return localPath, contentType, extention, err
}

// GetFileContentTypeAndExtention ...
func GetFileContentTypeAndExtention(file []byte) (contentType string, extention string, err error) {
	contentType = http.DetectContentType(file)

	if !str.Contains(imageContentTypeWhitelist, contentType) {
		return contentType, contentType, errors.New("invalid_content_type")
	}

	contentTypeArr := strings.Split(contentType, "/")
	if len(contentTypeArr) != 2 {
		return contentType, contentType, errors.New("invalid_content_type_length")
	}

	extention = contentTypeArr[1]

	return contentType, extention, nil
}

// PathToBase64 ...
func PathToBase64(path string) (res, contentType string) {
	// Open file from path
	file, err := os.Open(path)
	if err != nil {
		return res, contentType
	}
	defer file.Close()

	// Get the content
	contentType, err = GetFileContentType(file)
	if err != nil {
		return res, contentType
	}

	// Read file as byte[]
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return res, contentType
	}

	// Encode as base64.
	res = base64.StdEncoding.EncodeToString(data)

	return res, contentType
}
