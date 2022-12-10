package aesfront

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// Credential ...
type Credential struct {
	Key string
	Iv  string
}

// Decrypt ...
func (cred *Credential) Decrypt(data string) (res string, err error) {
	if data == "" {
		return res, nil
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return res, err
	}

	block, err := aes.NewCipher([]byte(cred.Key))
	if err != nil {
		return res, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return res, err
	}

	mode := cipher.NewCBCDecrypter(block, []byte(cred.Iv))
	mode.CryptBlocks(ciphertext, ciphertext)

	return string(PKCS5Trimming(ciphertext)), err
}

// PKCS5Trimming ...
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
