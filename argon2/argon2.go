package pwd

import (
	"fmt"

	"github.com/palp1tate/go-crypto-guard"
	"golang.org/x/crypto/argon2"
)

func GenArgon2(password string, saltLength, keyLength, iterations int) (encryptedPassword string, err error) {
	password, saltLength, keyLength, iterations, err = pwd.ParseParameters(password, saltLength, keyLength, iterations)
	if err != nil {
		return
	}
	salt, err := pwd.GenerateSalt(saltLength)
	if err != nil {
		return
	}
	dk := argon2.IDKey([]byte(password), []byte(salt), uint32(iterations), 1024, uint8(1), uint32(keyLength))
	encryptedPassword = fmt.Sprintf("%s$%d$%s$%s", pwd.Argon2, iterations, salt, pwd.Encode2string(dk))
	return
}

func VerifyArgon2(password, encryptedPassword string) (isValid bool, err error) {
	_, iter, salt, storedHash, err := pwd.ParsePassword(encryptedPassword)
	if err != nil {
		return
	}
	dk := argon2.IDKey([]byte(password), []byte(salt), uint32(iter), 1024, uint8(1), uint32(len(storedHash)/2))
	isValid = storedHash == pwd.Encode2string(dk)
	return
}
