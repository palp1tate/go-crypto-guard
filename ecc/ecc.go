package pwd

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"

	"github.com/palp1tate/go-crypto-guard"
)

func ECC(password string, privateKey *ecdsa.PrivateKey) (encryptedPassword string, err error) {
	digest := sha256.Sum256([]byte(password))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, digest[:])
	if err != nil {
		return
	}
	param := privateKey.Curve.Params()
	curveOrderByteSize := param.P.BitLen() / 8
	rByte, sByte := r.Bytes(), s.Bytes()
	signature := make([]byte, curveOrderByteSize*2)
	copy(signature[curveOrderByteSize-len(rByte):], rByte)
	copy(signature[curveOrderByteSize*2-len(sByte):], sByte)
	encryptedPassword = fmt.Sprintf("%s$%s", pwd.ECC, pwd.EncodeToString(signature))
	return
}

func VerifyECC(password, encryptedPassword string, publicKey ecdsa.PublicKey) (isValid bool, err error) {
	parts := strings.Split(encryptedPassword, "$")
	if len(parts) != 2 {
		err = fmt.Errorf("invalid encrypted password")
		return
	}
	signature, err := pwd.DecodeString(parts[1])
	if err != nil {
		return
	}
	digest := sha256.Sum256([]byte(password))
	curveOrderByteSize := publicKey.Curve.Params().P.BitLen() / 8
	r, s := new(big.Int), new(big.Int)
	r.SetBytes(signature[:curveOrderByteSize])
	s.SetBytes(signature[curveOrderByteSize:])
	isValid = ecdsa.Verify(&publicKey, digest[:], r, s)
	return
}
