package pwd

import (
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
	"golang.org/x/crypto/blowfish"
)

func GenBlowfish(password, blowfishKey string) (encryptedPassword string, err error) {
	if len(password) != 8 {
		err = fmt.Errorf("password length must be 8")
		return
	}
	cipher, err := blowfish.NewCipher([]byte(blowfishKey))
	if err != nil {
		return
	}
	src := []byte(password)
	dst := make([]byte, len(src))
	cipher.Encrypt(dst, src)
	encryptedPassword = fmt.Sprintf("%s$%s", pwd.Blowfish, pwd.Encode2string(dst))
	return
}

func VerifyBlowfish(password, encryptedPassword, blowfishKey string) (isValid bool, err error) {
	if len(password) != 8 {
		err = fmt.Errorf("password length must be 8")
		return
	}
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid encrypted password format")
		return
	}
	cipherTextBytes, err := pwd.Decode2byte(parts[1])
	if err != nil {
		return
	}
	cipher, err := blowfish.NewCipher([]byte(blowfishKey))
	if err != nil {
		return
	}
	dst := make([]byte, len(cipherTextBytes))
	cipher.Decrypt(dst, cipherTextBytes)
	isValid = password == string(dst)
	return
}
