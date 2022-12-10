package hmacsha

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Credential ...
type Credential struct {
	Key string
}

// Encrypt ...
func (cred *Credential) Encrypt(data string) (res string) {
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(cred.Key))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	res = hex.EncodeToString(h.Sum(nil))

	return res
}
