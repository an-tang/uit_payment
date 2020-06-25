package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"os"
)

func CreateRSA(params interface{}, publicKeyStr string) (string, error) {
	var hashStr string
	key := fmt.Sprintf("`\n-----BEGIN PUBLIC KEY-----\n%v\n-----END PUBLIC KEY-----", publicKeyStr)

	publicKeyByte := []byte(key)
	block, _ := pem.Decode([]byte(publicKeyByte))
	if block == nil {
		err := errors.New("failed to parse PEM block containing the public key")
		return hashStr, err
	}
	pkixPub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		err := errors.New("failed to parse DER encoded public key: " + err.Error())
		return hashStr, err
	}

	var publicKey *rsa.PublicKey
	publicKey = pkixPub.(*rsa.PublicKey)
	randomReader := rand.Reader
	rawJashJSON, err := json.Marshal(params)
	if err != nil {
		log.Println(err)
		return hashStr, err
	}
	ciphertext, err := rsa.EncryptPKCS1v15(randomReader, publicKey, rawJashJSON)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return hashStr, err
	}
	hashStr = base64.StdEncoding.EncodeToString(ciphertext)
	return hashStr, err
}

func CreateSHA256Signature(privateKeyStr string, data string) (string, error) {
	hashKey := ""
	key := fmt.Sprintf("\n-----BEGIN RSA PRIVATE KEY-----\n%s\n-----END RSA PRIVATE KEY-----", privateKeyStr)
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return hashKey, errors.New("ssh: no key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return hashKey, err
		}
		rawkey = rsa
	default:
		return hashKey, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}

	var sshKey *rsa.PrivateKey
	switch t := rawkey.(type) {
	case *rsa.PrivateKey:
		sshKey = t
	default:
		return hashKey, fmt.Errorf("ssh: unsupported key type %T", rawkey)
	}

	hash := sha256.New()
	hash.Write([]byte(data))
	hashedData := hash.Sum(nil)
	signed, err := rsa.SignPKCS1v15(rand.Reader, sshKey, crypto.SHA256, hashedData)
	if err != nil {

	}

	hashKey = base64.StdEncoding.EncodeToString(signed)
	return hashKey, nil
}
