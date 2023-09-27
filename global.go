package pwd

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	DefaultSaltLength = 16
	DefaultIterations = 50
	DefaultKeyLength  = 32
)

type Algorithm string

const (
	SHA512   Algorithm = "pbkdf2_sha512"
	SHA384   Algorithm = "pbkdf2_sha384"
	SHA256   Algorithm = "pbkdf2_sha256"
	SHA1     Algorithm = "pbkdf2_sha1"
	Md5      Algorithm = "pbkdf2_md5"
	Bcrypt   Algorithm = "bcrypt"
	Scrypt   Algorithm = "scrypt"
	Argon2   Algorithm = "argon2"
	HMAC     Algorithm = "hmac"
	Blake2b  Algorithm = "blake2b"
	Blake2s  Algorithm = "blake2s"
	AES      Algorithm = "aes"
	Blowfish Algorithm = "blowfish"
	DES      Algorithm = "des"
	ThreeDES Algorithm = "3des"
	ECC      Algorithm = "ecc"
	RC4      Algorithm = "rc4"
	RSA      Algorithm = "rsa"
)

var HashMap = map[Algorithm]func() hash.Hash{
	SHA512: sha512.New,
	SHA384: sha512.New384,
	SHA256: sha256.New,
	SHA1:   sha1.New,
	Md5:    md5.New,
}

func GenerateSalt(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func GeneratePBKDF2(password string, salt string, iter int, keyLen int, hashFunc func() hash.Hash) []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, hashFunc)
}

func Encode2string(dk []byte) string {
	return hex.EncodeToString(dk)
}

func Decode2byte(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func EncodeToString(dk []byte) string {
	return base64.StdEncoding.EncodeToString(dk)
}

func DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCS7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, fmt.Errorf("empty data")
	}
	unPadding := int(origData[length-1])
	if unPadding < 1 || unPadding > length {
		return nil, fmt.Errorf("invalid padding length")
	}
	return origData[:(length - unPadding)], nil
}

func ParseParameters(password string, saltLength, keyLength, iterations int) (string, int, int, int, error) {
	if password == "" {
		err := fmt.Errorf("password cannot be empty")
		return "", 0, 0, 0, err
	}
	if saltLength <= 0 {
		saltLength = DefaultSaltLength
	}
	if keyLength <= 0 {
		keyLength = DefaultKeyLength
	}
	if iterations <= 0 {
		iterations = DefaultIterations
	}
	return password, saltLength, keyLength, iterations, nil
}

func ParsePassword(encoded string) (algorithm Algorithm, iter int, salt, storedHash string, err error) {
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
