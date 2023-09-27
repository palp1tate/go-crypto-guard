package pwd

import "testing"

func TestBlowfish(t *testing.T) {
	pwd := "12345678"
	key := "yourkey2523822,8mmsAFVDtesdhBzdvcvcvvsdxDV"
	encryptedPwd, err := Blowfish(pwd, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(encryptedPwd)
	isValid, err := VerifyBlowfish(pwd, encryptedPwd, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(isValid)
}
