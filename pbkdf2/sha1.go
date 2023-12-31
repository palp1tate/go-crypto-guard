package pwd

import (
	"fmt"

	"github.com/palp1tate/go-crypto-guard"
)

func GenSHA1(password string, saltLength, keyLength, iterations int) (encryptedPassword string, err error) {
	password, saltLength, keyLength, iterations, err = pwd.ParseParameters(password, saltLength, keyLength, iterations)
	if err != nil {
		return
	}
	salt, err := pwd.GenerateSalt(saltLength)
	if err != nil {
		return
	}
	dk := pwd.GeneratePBKDF2(password, salt, iterations, keyLength, pwd.HashMap[pwd.SHA1])
	encryptedPassword = fmt.Sprintf("%s$%d$%s$%s", pwd.SHA1, iterations, salt, pwd.Encode2string(dk))
	return
}

func VerifySHA1(password, encryptedPassword string) (isValid bool, err error) {
	_, iter, salt, storedHash, err := pwd.ParsePassword(encryptedPassword)
	if err != nil {
		return
	}
	dk := pwd.GeneratePBKDF2(password, salt, iter, len(storedHash)/2, pwd.HashMap[pwd.SHA1])
	isValid = storedHash == pwd.Encode2string(dk)
	return
}
