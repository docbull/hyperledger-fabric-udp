package des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

// this package provides DES cryptography for Blockchain Block data
// sending to other peers.

func DesEncryption(key, iv, plainText []byte) ([]byte, error) {
	cipherBlock, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := cipherBlock.BlockSize()
	originalData := PKCS5Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(cipherBlock, iv)
	cryptedData := make([]byte, len(originalData))
	blockMode.CryptBlocks(cryptedData, originalData)
	return cryptedData, nil
}

func DesDecryption(key, iv, cipherText []byte) ([]byte, error) {
	cipherBlock, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(cipherBlock, iv)
	originalData := make([]byte, len(cipherText))
	blockMode.CryptBlocks(originalData, cipherText)
	originalData = PKCS5UnPadding(originalData)
	return originalData, nil
}

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
