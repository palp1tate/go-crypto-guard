package pwd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
)

func GenRSAKey(bits int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	privateStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block1 := pem.Block{
		Type:  "private key",
		Bytes: privateStream,
	}
	fPrivate, err := os.Create("privateKey.pem")
	if err != nil {
		return err
	}
	defer fPrivate.Close()
	err = pem.Encode(fPrivate, &block1)
	if err != nil {
		return err
	}
	publicKey := privateKey.PublicKey
	publicStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	block2 := pem.Block{
		Type:  "public key",
		Bytes: publicStream,
	}
	fPublic, err := os.Create("publicKey.pem")
	if err != nil {
		return err
	}
	defer fPublic.Close()
	pem.Encode(fPublic, &block2)
	return nil
}

func GenRSA(password, publicKeyPath string) (encryptedPassword string, err error) {
	f, err := os.Open(publicKeyPath)
	if err != nil {
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	block, _ := pem.Decode(b)
	keyInit, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pubKey := keyInit.(*rsa.PublicKey)
	res, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey, []byte(password))
	encryptedPassword = fmt.Sprintf("%s$%s", pwd.RSA, pwd.EncodeToString(res))
	return
}

func VerifyRSA(password, encryptedPassword, privateKeyPath string) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid encrypted password format")
		return
	}
	decodedPassword, err := pwd.DecodeString(parts[1])
	f, _ := os.Open(privateKeyPath)
	defer f.Close()
	fileInfo, _ := f.Stat()
	b := make([]byte, fileInfo.Size())
	f.Read(b)
	block, _ := pem.Decode(b)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	encodedPassword, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, decodedPassword)
	isValid = password == string(encodedPassword)
	return
}
