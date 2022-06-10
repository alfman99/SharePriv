package util

import (
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"sharepriv/entities"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

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

func DesencriptarArchivo(data []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic("error NewCipher")
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic("error NewGCM")
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherData := data[:nonceSize], data[nonceSize:]
	plainData, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return nil, err
	}
	return plainData, nil
}

func GenerateRandomString(length int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letter[seededRand.Intn(len(letter))]
	}
	return string(b)
}

func ContainsGroup(s []entities.Grupo, Id string) bool {
	for _, v := range s {
		if v.Id == Id {
			return true
		}
	}

	return false
}
