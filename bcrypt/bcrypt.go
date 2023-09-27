package pwd

import (
	"fmt"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
	"golang.org/x/crypto/bcrypt"
)

func Bcrypt(password string) (encryptedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	encryptedPassword = fmt.Sprintf("%s$%s", pwd.Bcrypt, pwd.Encode2string(hash))
	return
}

func VerifyBcrypt(password, encryptedPassword string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid encrypted password format")
		return
	}
	var str2byte []byte
	str2byte, err = pwd.Decode2byte(parts[1])
	if err != nil {
		return false, err
	}
	return bcrypt.CompareHashAndPassword(str2byte, []byte(password)) == nil, nil
}
