package pwd

import "testing"

const password = "123456"

func TestVerifySha512(t *testing.T) {
	encoded, err := Generate(password, &Options{
		SaltLen:    20,
		KeyLen:     20,
		Iterations: 100,
		Algorithm:  SHA512,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifySHA512 success")
	}
}

func TestVerifySha384(t *testing.T) {
	encoded, err := Generate(password, &Options{
		SaltLen:    20,
		KeyLen:     20,
		Iterations: 100,
		Algorithm:  SHA384,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifySHA384 success")
	}
}

func TestVerifySha256(t *testing.T) {
	encoded, err := Generate(password, &Options{
		SaltLen:    20,
		KeyLen:     20,
		Iterations: 100,
		Algorithm:  SHA256,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifySHA256 success")
	}
}

func TestVerifySha1(t *testing.T) {
	encoded, err := Generate(password, &Options{
		SaltLen:    20,
		KeyLen:     20,
		Iterations: 100,
		Algorithm:  SHA1,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifySHA1 success")
	}
}

func TestVerifyMd5(t *testing.T) {
	encoded, err := Generate(password, &Options{
		SaltLen:    20,
		KeyLen:     20,
		Iterations: 100,
		Algorithm:  Md5,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifyMd5 success")
	}
}

func TestVerifyBcrypt(t *testing.T) {
	encoded, err := Generate(password, &Options{
		SaltLen:   874,
		Algorithm: Bcrypt,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifyBcrypt success")
	}
}

func TestVerifyArgon2(t *testing.T) {
	encoded, err := Generate(password, &Options{
		KeyLen:    58,
		Algorithm: Argon2,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifyArgon2 success")
	}
}

func TestVerifyHmac(t *testing.T) {
	encoded, err := Generate(password, &Options{
		Iterations: 96,
		Algorithm:  HMAC,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifyHmac success")
	}
}

func TestVerifyBlake2b(t *testing.T) {
	encoded, err := Generate(password, &Options{
		SaltLen:   96,
		Algorithm: Blake2b,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifyBlake2b success")
	}
}

func TestVerifyBlake2s(t *testing.T) {
	encoded, err := Generate(password, &Options{
		SaltLen:   85,
		Algorithm: Blake2s,
	})
	if err != nil {
		t.Error(err)
	}
	if ok, err := Verify(password, encoded); !ok {
		t.Error(err)
	} else {
		t.Log("VerifyBlake2s success")
	}
}
