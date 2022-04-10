package minio

import (
	"crypto/tls"
	"github.com/minio/minio-go/v6"
	"net/http"
)

// Connection ...
type Connection struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BaseURL         string
	DefaultBucket   string
}

// Connect ...
func (m Connection) Connect() (*minio.Client, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(m.Endpoint, m.AccessKeyID, m.SecretAccessKey, m.UseSSL)

	if !m.UseSSL {
		minioClient.SetCustomTransport(&http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}})
	}

	minioModel := NewMinioModel(minioClient)
	minioModel.CreateBucket(m.DefaultBucket, "")

	return minioClient, err
}
