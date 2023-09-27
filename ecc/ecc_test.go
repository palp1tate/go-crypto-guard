package pwd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

func TestECC(t *testing.T) {
	pwd := "12345678"
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	encryptedPassword, err := ECC(pwd, privateKey)
	if err != nil {
		t.Errorf("ECC error: %v", err)
	}
	t.Log(encryptedPassword)
	publicKey := privateKey.PublicKey
	isValid, err := VerifyECC(pwd, encryptedPassword, publicKey)
	if !isValid {
		t.Errorf("VerifyECC error: %v", err)
	}
}
