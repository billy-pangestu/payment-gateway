package jwe

import (
	"encoding/json"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwe"
)

// Credential ...
type Credential struct {
	KeyLocation string
	Passphrase  string
}

// Generate ...
func (cred *Credential) Generate(payload map[string]interface{}) (res string, err error) {
	privkey, err := rsaConfigSetup(cred.KeyLocation, cred.Passphrase)
	if err != nil {
		return res, err
	}

	// Convert payload to string
	payloadString, err := json.Marshal(payload)
	if err != nil {
		return res, err
	}

	// Generate JWE
	jweRes, err := jwe.Encrypt([]byte(payloadString), jwe.WithKey(jwa.RSA1_5, &privkey.PublicKey), jwe.WithContentEncryption(jwa.A128CBC_HS256), jwe.WithCompress(jwa.Deflate))
	res = string(jweRes)

	return res, err
}

// Rollback ...
func (cred *Credential) Rollback(userID string) (res map[string]interface{}, err error) {
	privkey, err := rsaConfigSetup(cred.KeyLocation, cred.Passphrase)
	if err != nil {
		return res, err
	}

	decrypted, err := jwe.Decrypt([]byte(userID), jwe.WithKey(jwa.RSA1_5, privkey))
	if err != nil {
		return res, err
	}

	res = map[string]interface{}{}
	err = json.Unmarshal(decrypted, &res)
	if err != nil {
		return res, err
	}

	return res, err
}
