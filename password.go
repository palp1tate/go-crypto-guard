package pwd

import (
	"fmt"
)

func Generate(password string, options *Options) (hashPwd string, err error) {
	opt := NewOptions(options)
	algorithm, iter, saltLen, keyLen := opt.Algorithm, opt.Iterations, opt.SaltLen, opt.KeyLen
	var dk []byte
	salt, err := generateSalt(saltLen)
	if err != nil {
		return
	}
	if hashFunc, ok := hashMap[algorithm]; ok {
		dk = generatePBKDF2(password, salt, iter, keyLen, hashFunc)
	} else {
		switch algorithm {
		case Bcrypt:
			dk, err = generateBcrypt(password)
			if err != nil {
				return
			}
		case Scrypt:
			dk, err = generateScrypt(password, salt, keyLen)
			if err != nil {
				return
			}
		case Argon2:
			dk = generateArgon2(password, salt, iter, keyLen)
		case HMAC:
			dk = generateHMAC(salt, password)
		case Blake2b:
			salt, err = generateSalt(16)
			if err != nil {
				return
			}
			dk, err = generateBlake2b(salt, password)
			if err != nil {
				return
			}
		case Blake2s:
			salt, err = generateSalt(8)
			if err != nil {
				return
			}
			dk, err = generateBlake2s(salt, password)
			if err != nil {
				return
			}
		default:
			return "", fmt.Errorf("algorithm %s not supported", algorithm)
		}
	}
	hashPwd = fmt.Sprintf("%s$%d$%s$%s", algorithm, iter, salt, encode2string(dk))
	return
}

func Verify(password, encoded string) (isValid bool, err error) {
	algorithm, iter, salt, storedHash, err := parsePassword(encoded)
	if err != nil {
		return
	}
	var dk []byte
	if hashFunc, ok := hashMap[algorithm]; ok {
		dk = generatePBKDF2(password, salt, iter, len(storedHash)/2, hashFunc)
	} else {
		switch algorithm {
		case Bcrypt:
			return checkBcrypt(password, storedHash)
		case Scrypt:
			dk, err = generateScrypt(password, salt, len(storedHash)/2)
			if err != nil {
				return
			}
		case Argon2:
			dk = generateArgon2(password, salt, iter, len(storedHash)/2)
		case HMAC:
			dk = generateHMAC(salt, password)
		case Blake2b:
			dk, err = generateBlake2b(salt, password)
			if err != nil {
				return
			}
		case Blake2s:
			dk, err = generateBlake2s(salt, password)
			if err != nil {
				return
			}
		default:
			return false, fmt.Errorf("unknown algorithm")
		}
	}
	if encode2string(dk) == storedHash {
		isValid = true
	}
	return
}
