package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"hash"
)

const (
	//SHA256 algorithm
	SHA256 = "SHA256"
	//SHA512 algorithm
	SHA512 = "SHA512"
)

func Base64Encode(algorithm string, key string, data string) string {
	hmac := encode(algorithm, key, data)
	return base64.StdEncoding.EncodeToString(hmac)
}

func HexStringEncode(algorithm string, key string, data string) string {
	hmac := encode(algorithm, key, data)
	return hex.EncodeToString(hmac)
}

func HashMD5(params string) string {
	hasher := md5.New()
	hasher.Write([]byte(params))
	return hex.EncodeToString(hasher.Sum(nil))
}

func EncodeByteParamSHA256(params []byte) (string, error) {
	if len(params) == 0 {
		return "", nil
	}
	sha256RequestBody := sha256.Sum256(params)
	contentDigest := base64.StdEncoding.EncodeToString(sha256RequestBody[:])
	return contentDigest, nil
}

func encode(algorithm string, key string, data string) []byte {
	mac := hmac.New(getAlgorithm(algorithm), []byte(key))
	mac.Write([]byte(data))
	return mac.Sum(nil)
}

func getAlgorithm(algorithm string) func() hash.Hash {
	switch algorithm {
	case SHA256:
		return sha256.New
	case SHA512:
		return sha512.New
	}

	return nil
}
