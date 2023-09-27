package pwd

import "testing"

func TestRSA(t *testing.T) {
	_ = GenRSAKey(2048)
	pwd := "123456wdWFezWAgf"
	encryptedPwd, err := RSA(pwd, "publicKey.pem")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encryptedPwd)
	isValid, err := VerifyRSA(pwd, encryptedPwd, "privateKey.pem")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(isValid)
}
