package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"mime/multipart"
	"net/http"
)

func EncriptarArchivo(data []byte, key []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		panic(err)
	}
	cipherData := gcm.Seal(nonce, nonce, data, nil)
	return cipherData
}

func DesencriptarArchivo(data []byte, key []byte) []byte {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherData := data[:nonceSize], data[nonceSize:]
	plainData, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		panic(err)
	}
	return plainData
}

func GetFileContentType(out multipart.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	return http.DetectContentType(buffer), nil
}
