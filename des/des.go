package pwd

import (
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
)

func GenDES(password, desKey string) (encryptedPassword string, err error) {
	if len(desKey) != 8 {
		err = fmt.Errorf("the length of des key must be 8")
	}
	block, err := des.NewCipher([]byte(desKey))
	if err != nil {
		return
	}
	src1 := pwd.PKCS7Padding([]byte(password), block.BlockSize())
	iv := make([]byte, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	desc := make([]byte, len(src1))
	blockMode.CryptBlocks(desc, src1)
	encryptedPassword = fmt.Sprintf("%s$%s", pwd.DES, pwd.EncodeToString(desc))
	return
}

func VerifyDES(password, encryptedPassword, desKey string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 2 {
		return false, fmt.Errorf("invalid encrypted password")
	}
	decodedPassword, err := pwd.DecodeString(parts[1])
	if err != nil {
		return
	}
	block, err := des.NewCipher([]byte(desKey))
	if err != nil {
		return
	}
	iv := make([]byte, block.BlockSize())
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(decodedPassword, decodedPassword)
	encodedPassword, err := pwd.PKCS7UnPadding(decodedPassword)
	if err != nil {
		return
	}
	isValid = password == string(encodedPassword)
	return
}
