package pwd

import (
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
	"golang.org/x/crypto/blake2b"
)

func Blake2b(password string) (encryptedPassword string, err error) {
	salt, err := pwd.GenerateSalt(16)
	if err != nil {
		return
	}
	h, err := blake2b.New256([]byte(salt))
	if err != nil {
		return
	}
	h.Write([]byte(password))
	encryptedPassword = fmt.Sprintf("%s$%s$%s", pwd.Blake2b, salt, pwd.Encode2string(h.Sum(nil)))
	return
}

func VerifyBlake2b(password, encryptedPassword string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 3 {
		err = fmt.Errorf("invalid encrypted password format")
		return
	}
	salt := parts[1]
	storedHash := parts[2]
	h, err := blake2b.New256([]byte(salt))
	if err != nil {
		return
	}
	h.Write([]byte(password))
	isValid = storedHash == pwd.Encode2string(h.Sum(nil))
	return
}
