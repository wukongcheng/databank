package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"errors"
)

func Encrypt(key []byte, plainText []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil,err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return cipherText,nil
}

func Decrypt(key,cipherText []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil,err
	}

	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return nil,err
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}