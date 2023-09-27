package pwd

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
)

func GenAES(password, aesKey string) (encryptedPassword string, err error) {
	if len(aesKey) != 32 {
		err = fmt.Errorf("the length of aes key must be 32")
		return
	}
	passwordBytes, aesKeyBytes := []byte(password), []byte(aesKey)
	block, err := aes.NewCipher(aesKeyBytes)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	passwordBytes = pwd.PKCS7Padding(passwordBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, aesKeyBytes[:blockSize])
	ciphertext := make([]byte, len(passwordBytes))
	blockMode.CryptBlocks(ciphertext, passwordBytes)
	encryptedPassword = fmt.Sprintf("%s$%s", pwd.AES, pwd.EncodeToString(ciphertext))
	return
}

func VerifyAES(password, encryptedPassword, aesKey string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid encrypted password format")
		return
	}
	aesKeyBytes := []byte(aesKey)
	ciphertext, err := pwd.DecodeString(parts[1])
	if err != nil {
		return
	}
	block, err := aes.NewCipher(aesKeyBytes)
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, aesKeyBytes[:blockSize])
	origData := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(origData, ciphertext)
	decodedPassword, err := pwd.PKCS7UnPadding(origData)
	if err != nil {
		return
	}
	isValid = password == string(decodedPassword)
	return
}
