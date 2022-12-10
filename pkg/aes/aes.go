package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

// Credential ...
type Credential struct {
	Key string
}

// shaKey ...
func (cred *Credential) shaKey() (res []byte) {
	h := sha1.New()
	h.Write([]byte(cred.Key))
	res = h.Sum(nil)
	res = res[:16]

	return res
}

// Encrypt ...
func (cred *Credential) Encrypt(textString string) (string, error) {
	key := cred.shaKey()
	text := []byte(textString)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	res := hex.EncodeToString(ciphertext)
	return res, nil
}

// EncryptNoErr ...
func (cred *Credential) EncryptNoErr(textString string) string {
	key := cred.shaKey()
	text := []byte(textString)

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return ""
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	res := hex.EncodeToString(ciphertext)
	return res
}

// Decrypt ...
func (cred *Credential) Decrypt(text string) (string, error) {
	key := cred.shaKey()
	ciphertext, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	data, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return "", err
	}

	res := string(data[:])
	return res, nil
}

// DecryptNoErr ...
func (cred *Credential) DecryptNoErr(text string) string {
	key := cred.shaKey()
	ciphertext, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	if len(ciphertext) < aes.BlockSize {
		return ""
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)
	data, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return ""
	}

	res := string(data[:])
	return res
}
