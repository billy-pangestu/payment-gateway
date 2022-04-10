package jwe

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func rsaConfigSetup(rsaPrivateKeyLocation, rsaPrivateKeyPassword string) (*rsa.PrivateKey, error) {
	if rsaPrivateKeyLocation == "" {
		fmt.Println("No RSA Key given, generating temp one")
		return GenRSA(4096)
	}

	priv, err := ioutil.ReadFile(rsaPrivateKeyLocation)
	if err != nil {
		fmt.Println("No RSA private key found, generating temp one")
		return GenRSA(4096)
	}

	privPem, _ := pem.Decode(priv)
	var privPemBytes []byte
	if privPem.Type != "RSA PRIVATE KEY" {
		fmt.Println("RSA private key is of the wrong type")
	}

	if rsaPrivateKeyPassword != "" {
		privPemBytes, err = x509.DecryptPEMBlock(privPem, []byte(rsaPrivateKeyPassword))
	} else {
		privPemBytes = privPem.Bytes
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil { // note this returns type `interface{}`
			fmt.Println("Unable to parse RSA private key, generating a temp one")
			return GenRSA(4096)
		}
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Println("Unable to parse RSA private key, generating a temp one (2)")
		return GenRSA(4096)
	}

	return privateKey, nil
}

// GenRSA returns a new RSA key of bits length
func GenRSA(bits int) (*rsa.PrivateKey, error) {
	key, err := rsa.GenerateKey(rand.Reader, bits)
	return key, err
}
