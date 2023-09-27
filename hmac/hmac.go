package pwd

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
)

func GenHMAC(password string, saltLength int) (encryptedPassword string, err error) {
	if saltLength <= 0 {
		saltLength = pwd.DefaultSaltLength
	}
	salt, err := pwd.GenerateSalt(saltLength)
	if err != nil {
		return
	}
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(password))
	encryptedPassword = fmt.Sprintf("%s$%s$%s", pwd.HMAC, salt, pwd.Encode2string(h.Sum(nil)))
	return
}

func VerifyHMAC(password, encryptedPassword string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 3 {
		err = fmt.Errorf("invalid encrypted password format")
		return
	}
	salt := parts[1]
	storedHash := parts[2]
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(password))
	isValid = storedHash == pwd.Encode2string(h.Sum(nil))
	return
}
