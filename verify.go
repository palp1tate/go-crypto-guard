package pwd

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

func Verify(password, encoded string) (bool, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 4 {
		return false, fmt.Errorf("incorrect password format")
	}
	algorithm := parts[0]
	iter, _ := strconv.Atoi(parts[1])
	salt := parts[2]
	storedHash := parts[3]
	var dk []byte

	switch algorithm {

	case SHA512:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, len(storedHash)/2, sha512.New)

	case SHA384:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, len(storedHash)/2, sha512.New384)

	case SHA256:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, len(storedHash)/2, sha256.New)

	case SHA1:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, len(storedHash)/2, sha1.New)

	case Md5:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, len(storedHash)/2, md5.New)

	case Bcrypt:
		str2byte, err := hex.DecodeString(storedHash)
		if err != nil {
			return false, err
		}
		isValid := bcrypt.CompareHashAndPassword(str2byte, []byte(password))
		return isValid == nil, nil

	case Scrypt:
		var err error
		dk, err = scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
		if err != nil {
			return false, err
		}
	case Argon2:
		dk = argon2.IDKey([]byte(password), []byte(salt), uint32(iter), uint32(len(storedHash)/2), uint8(1), uint32(len(storedHash)/2))

	case HMAC:
		h := hmac.New(sha256.New, []byte(salt))
		h.Write([]byte(password))
		dk = h.Sum(nil)

	case Blake2b:
		h, err := blake2b.New256([]byte(salt))
		if err != nil {
			return false, err
		}
		h.Write([]byte(password))
		dk = h.Sum(nil)

	case Blake2s:
		h, err := blake2s.New256([]byte(salt))
		if err != nil {
			return false, err
		}
		h.Write([]byte(password))
		dk = h.Sum(nil)

	default:
		return false, fmt.Errorf("unknown algorithm")
	}

	return hex.EncodeToString(dk) == storedHash, nil
}
