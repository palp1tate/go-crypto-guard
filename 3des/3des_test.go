package pwd

import "testing"

func TestThreeDES(t *testing.T) {
	pwd := "1285rt455s45b5er4b434"
	key := "123456789123485456789456"
	encryptedPwd, err := ThreeDES(pwd, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encryptedPwd)
	// pwd = "12345"
	isValid, err := VerifyThreeDES(pwd, encryptedPwd, key)
	if !isValid {
		t.Fatal(err)
	}
}
