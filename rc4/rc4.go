package pwd

import (
	"crypto/rc4"
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
)

func RC4(password, rc4Key string) (encryptedPassword string, err error) {
	cipher, err := rc4.NewCipher([]byte(rc4Key))
	if err != nil {
		return
	}
	ciphertext := make([]byte, len(password))
	cipher.XORKeyStream(ciphertext, []byte(password))
	encryptedPassword = fmt.Sprintf("%s$%s", pwd.RC4, pwd.Encode2string(ciphertext))
	return
}

func VerifyRC4(password, encryptedPassword, rc4Key string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid encrypted password format")
		return
	}
	cipherTextBytes, err := pwd.Decode2byte(parts[1])
	if err != nil {
		return
	}
	cipher, err := rc4.NewCipher([]byte(rc4Key))
	if err != nil {
		return
	}
	dst := make([]byte, len(cipherTextBytes))
	cipher.XORKeyStream(dst, cipherTextBytes)
	isValid = password == string(dst)
	return
}
