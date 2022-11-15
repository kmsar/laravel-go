package Encription

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/kmsar/laravel-go/Framework/Contracts/Support"
	"github.com/pkg/errors"

	"github.com/kmsar/laravel-go/Framework/Contracts/IEncryption"
)

type aesEncryptor struct {
	key   []byte
	block cipher.Block
}

func AES(key string) IEncryption.Encryptor {
	keyBytes := []byte(key)
	block, err := aes.NewCipher(keyBytes)

	if err != nil {
		panic(EncryptException{
			err, Support.Fields{"key": key},
		})
	}

	return &aesEncryptor{key: keyBytes, block: block}
}

func (this *aesEncryptor) Encode(value string) string {
	// 转成字节数组
	origData := []byte(value)

	// 获取秘钥块的长度
	blockSize := this.block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(this.block, this.key[:blockSize])
	// 创建数组
	encrypted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(encrypted, origData)

	return base64.StdEncoding.EncodeToString(encrypted)
}

func (this *aesEncryptor) Decode(encrypted string) (result string, err error) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			switch value := panicValue.(type) {
			case error:
				err = value
			default:
				err = errors.Errorf("%v", value)
			}
		}

	}()
	// 转成字节数组
	encryptedByte, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	// 获取秘钥块的长度
	blockSize := this.block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(this.block, this.key[:blockSize])
	// 创建数组
	orig := make([]byte, len(encryptedByte))
	// 解密
	blockMode.CryptBlocks(orig, encryptedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig), nil
}

// 补码
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
