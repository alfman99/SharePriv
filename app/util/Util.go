package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
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
