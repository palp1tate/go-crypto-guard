package pwd

import "testing"

func TestBcrypt(t *testing.T) {
	password := "123456"
	encryptedPassword, err := GenBcrypt(password)
	if err != nil {
		t.Error(err)
	}
	t.Log(encryptedPassword)
	// password = "1234567"
	if ok, err := VerifyBcrypt(password, encryptedPassword); !ok {
		t.Error("VerifyBcrypt failed", err)
	} else {
		t.Log("VerifyBcrypt success")
	}
}
