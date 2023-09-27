package pwd

import "testing"

func TestAES(t *testing.T) {
	pwd := "123456"
	key := "palpitateabcdefghijklmn123456789"
	encryptedPwd, err := GenAES(pwd, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(encryptedPwd)
	isValid, err := VerifyAES(pwd, encryptedPwd, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(isValid)
}
