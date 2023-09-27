package pwd

import (
	"fmt"

	"github.com/palp1tate/go-crypto-guard"
)

func GenMd5(password string, saltLength, keyLength, iterations int) (encryptedPassword string, err error) {
	password, saltLength, keyLength, iterations, err = pwd.ParseParameters(password, saltLength, keyLength, iterations)
	if err != nil {
		return
	}
	salt, err := pwd.GenerateSalt(saltLength)
	if err != nil {
		return
	}
	dk := pwd.GeneratePBKDF2(password, salt, iterations, keyLength, pwd.HashMap[pwd.Md5])
	encryptedPassword = fmt.Sprintf("%s$%d$%s$%s", pwd.Md5, iterations, salt, pwd.Encode2string(dk))
	return
}

func VerifyMd5(password, encryptedPassword string) (isValid bool, err error) {
	_, iter, salt, storedHash, err := pwd.ParsePassword(encryptedPassword)
	if err != nil {
		return
	}
	dk := pwd.GeneratePBKDF2(password, salt, iter, len(storedHash)/2, pwd.HashMap[pwd.Md5])
	isValid = storedHash == pwd.Encode2string(dk)
	return
}
