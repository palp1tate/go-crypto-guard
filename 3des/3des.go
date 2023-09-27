package pwd

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
)

func GenThreeDES(password, threeDesKey string) (encryptedPassword string, err error) {
	if len(threeDesKey) != 24 {
		err = fmt.Errorf("the length of 3des key must be 24")
		return
	}
	block, err := des.NewTripleDESCipher([]byte(threeDesKey))
	if err != nil {
		return
	}
	src := pwd.PKCS7Padding([]byte(password), block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(threeDesKey)[:block.BlockSize()])
	blockMode.CryptBlocks(src, src)
	encryptedPassword = fmt.Sprintf("%s$%s", pwd.ThreeDES, pwd.EncodeToString(src))
	return
}

func VerifyThreeDES(password, encryptedPassword, threeDesKey string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid encrypted password")
	}
	decodedPassword, err := pwd.DecodeString(parts[1])
	if err != nil {
		return
	}
	block, err := des.NewTripleDESCipher([]byte(threeDesKey))
	if err != nil {
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(threeDesKey)[:block.BlockSize()])
	blockMode.CryptBlocks(decodedPassword, decodedPassword)
	encodedPassword, err := pwd.PKCS7UnPadding(decodedPassword)
	if err != nil {
		return
	}
	isValid = password == string(encodedPassword)
	return
}
