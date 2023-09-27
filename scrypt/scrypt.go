package pwd

import (
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
	"golang.org/x/crypto/scrypt"
)

func Scrypt(password string, saltLength int, keyLength int) (encryptedPassword string, err error) {
	if saltLength <= 0 {
		saltLength = pwd.DefaultSaltLength
	}
	if keyLength <= 0 {
		keyLength = pwd.DefaultKeyLength
	}
	salt, err := pwd.GenerateSalt(saltLength)
	if err != nil {
		return
	}
	hash, err := scrypt.Key([]byte(password), []byte(salt), 512, 8, 1, keyLength)
	if err != nil {
		return
	}
	encryptedPassword = fmt.Sprintf("%s$%s$%s", pwd.Scrypt, salt, pwd.Encode2string(hash))
	return
}

func VerifyScrypt(password, encryptedPassword string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 3 {
		err = fmt.Errorf("invalid encrypted password format")
		return
	}
	salt := parts[1]
	storedHash := parts[2]
	hash, err := scrypt.Key([]byte(password), []byte(salt), 512, 8, 1, len(storedHash)/2)
	if err != nil {
		return
	}
	isValid = storedHash == pwd.Encode2string(hash)
	return
}
