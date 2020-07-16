package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/Bezunca/API/internal/config"
	"io/ioutil"
)

const CEIPassword = "cei_password"

func loadPublicKey(fileName string) (*rsa.PublicKey, error) {

	pub, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	pubPem, _ := pem.Decode(pub)

	parsedKey, err := x509.ParsePKCS1PublicKey(pubPem.Bytes)
	if err != nil {
		return nil, err
	}

	return parsedKey, nil
}

func RSAEncript(message, label string) (string, error) {

	configs := config.Get()

	publicKey, err := loadPublicKey(configs.RSAPublicKeyPath)
	if err != nil {
		return "", err
	}

	secretMessage := []byte(message)
	secretLabel := []byte(label)

	encryptedMessage, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, secretMessage, secretLabel)
	if err != nil {
		return "", err
	}

	b64EncryptedMessage := base64.StdEncoding.EncodeToString(encryptedMessage)
	return b64EncryptedMessage, nil
}
