package pwd

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

const (
	DefaultSaltLen    = 16
	DefaultIterations = 50
	DefaultKeyLen     = 32
	DefaultAlgorithm  = SHA512
)

type Algorithm string

const (
	SHA512  Algorithm = "pbkdf2_sha512"
	SHA384  Algorithm = "pbkdf2_sha384"
	SHA256  Algorithm = "pbkdf2_sha256"
	SHA1    Algorithm = "pbkdf2_sha1"
	Md5     Algorithm = "pbkdf2_md5"
	Bcrypt  Algorithm = "bcrypt"
	Scrypt  Algorithm = "scrypt"
	Argon2  Algorithm = "argon2"
	HMAC    Algorithm = "hmac"
	Blake2b Algorithm = "blake2b"
	Blake2s Algorithm = "blake2s"
)

type Options struct {
	SaltLen    int
	Iterations int
	KeyLen     int
	Algorithm  Algorithm
}

var hashMap = map[Algorithm]func() hash.Hash{
	SHA512: sha512.New,
	SHA384: sha512.New384,
	SHA256: sha256.New,
	SHA1:   sha1.New,
	Md5:    md5.New,
}

func NewOptions(options *Options) *Options {
	opt := &Options{
		SaltLen:    DefaultSaltLen,
		Iterations: DefaultIterations,
		KeyLen:     DefaultKeyLen,
		Algorithm:  DefaultAlgorithm,
	}
	if options != nil {
		if options.SaltLen != 0 {
			opt.SaltLen = options.SaltLen
		}
		if options.Iterations != 0 {
			opt.Iterations = options.Iterations
		}
		if options.KeyLen != 0 {
			opt.KeyLen = options.KeyLen
		}
		if options.Algorithm != "" {
			opt.Algorithm = options.Algorithm
		}
	}
	return opt
}

func generateSalt(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func generatePBKDF2(password string, salt string, iter int, keyLen int, hashFunc func() hash.Hash) []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, hashFunc)
}

func generateBcrypt(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func generateScrypt(password string, salt string, keyLen int) ([]byte, error) {
	return scrypt.Key([]byte(password), []byte(salt), 512, 8, 1, keyLen)
}

func generateArgon2(password string, salt string, iter int, keyLen int) []byte {
	return argon2.IDKey([]byte(password), []byte(salt), uint32(iter), uint32(keyLen), uint8(1), uint32(keyLen))
}

func generateHMAC(salt, password string) []byte {
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(password))
	return h.Sum(nil)
}

func generateBlake2b(salt, password string) ([]byte, error) {
	h, err := blake2b.New256([]byte(salt))
	if err != nil {
		return nil, err
	}
	h.Write([]byte(password))
	return h.Sum(nil), nil
}

func generateBlake2s(salt, password string) ([]byte, error) {
	h, err := blake2s.New256([]byte(salt))
	if err != nil {
		return nil, err
	}
	h.Write([]byte(password))
	return h.Sum(nil), nil
}

func encode2string(dk []byte) string {
	return hex.EncodeToString(dk)
}

func parsePassword(encoded string) (algorithm Algorithm, iter int, salt, storedHash string, err error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 4 {
		err = fmt.Errorf("incorrect password format")
		return
	}
	algorithm = Algorithm(parts[0])
	iter, _ = strconv.Atoi(parts[1])
	salt = parts[2]
	storedHash = parts[3]
	return
}

func checkBcrypt(password, storedHash string) (bool, error) {
	var str2byte []byte
	str2byte, err := hex.DecodeString(storedHash)
	if err != nil {
		return false, err
	}
	return bcrypt.CompareHashAndPassword(str2byte, []byte(password)) == nil, nil
}
