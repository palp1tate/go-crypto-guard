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

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
)

// Constants defining various password hashing algorithms
const (
	SHA512            = "pbkdf2_sha512"
	SHA384            = "pbkdf2_sha384"
	SHA256            = "pbkdf2_sha256"
	SHA1              = "pbkdf2_sha1"
	Md5               = "pbkdf2_md5"
	Bcrypt            = "bcrypt"
	Scrypt            = "scrypt"
	Argon2            = "argon2"
	HMAC              = "hmac"
	Blake2b           = "blake2b"
	Blake2s           = "blake2s"
	defaultSaltLen    = 16
	defaultIterations = 50
	defaultKeyLen     = 32
	defaultAlgorithm  = SHA512
)

// Options struct to customize hashing parameters
type Options struct {
	SaltLen    int    // Length of the salt to be used in hashing
	Iterations int    // Number of iterations for the hash function
	KeyLen     int    // Length of the derived key
	Algorithm  string // Hashing algorithm to be used
}

// generateSalt generates a random salt of the specified length
func generateSalt(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// Generate hashes a password using the specified algorithm and options
func Generate(password string, options *Options) (hashPwd string, err error) {
	saltLen := defaultSaltLen
	iter := defaultIterations
	keyLen := defaultKeyLen
	algorithm := defaultAlgorithm
	var dk []byte

	// Customize hashing options if provided
	if options != nil {
		if options.SaltLen != 0 {
			saltLen = options.SaltLen
		}
		if options.Iterations != 0 {
			iter = options.Iterations
		}
		if options.KeyLen != 0 {
			keyLen = options.KeyLen
		}
		if options.Algorithm != "" {
			algorithm = options.Algorithm
		}
	}

	// Generate a random salt
	salt := generateSalt(saltLen)

	// Hash the password based on the selected algorithm
	switch algorithm {

	case SHA512:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, sha512.New)

	case SHA256:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, sha256.New)

	case SHA1:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, sha1.New)

	case SHA384:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, sha512.New384)

	case Md5:
		dk = pbkdf2.Key([]byte(password), []byte(salt), iter, keyLen, md5.New)

	case Bcrypt:
		// Generate a Bcrypt hash
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		dk = hashedPassword

	case Scrypt:
		// Generate a Scrypt hash
		dk, err = scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
		if err != nil {
			return "", err
		}

	case Argon2:
		// Generate an Argon2 hash
		dk = argon2.IDKey([]byte(password), []byte(salt), uint32(iter), uint32(keyLen), uint8(1), uint32(keyLen))

	case HMAC:
		// Generate an HMAC hash
		h := hmac.New(sha256.New, []byte(salt))
		h.Write([]byte(password))
		dk = h.Sum(nil)

	case Blake2b:
		// Generate a Blake2b hash
		salt = generateSalt(16)
		h, err := blake2b.New256([]byte(salt))
		if err != nil {
			return "", err
		}
		h.Write([]byte(password))
		dk = h.Sum(nil)

	case Blake2s:
		// Generate a Blake2s hash
		salt = generateSalt(8)
		h, err := blake2s.New256([]byte(salt))
		if err != nil {
			return "", err
		}
		h.Write([]byte(password))
		dk = h.Sum(nil)

	default:
		return "", fmt.Errorf("algorithm %s not supported", algorithm)
	}

	// Format the hashed password for storage
	hashPwd = fmt.Sprintf("%s$%d$%s$%s", algorithm, iter, salt, hex.EncodeToString(dk))
	return
}
