<h1 align="center" style="border-bottom: none;">go-crypto-guard </h1>

<div class="labels" align="center">
    <a href="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
      <img src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg" alt="semantic-release">
    </a>
    <a href="https://pkg.go.dev/github.com/palp1tate/go-crypto-guard">
      <img src="https://godoc.org/github.com/palp1tate/go-crypto-guard?status.svg" alt="Godoc">
    </a>
    <a href="https://pkg.go.dev/github.com/palp1tate/go-crypto-guard?tab=doc">
      <img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square" alt="go.dev reference">
    </a>
    <a href="https://github.com/palp1tate/go-crypto-guard/blob/master/LICENSE">
      <img src="https://img.shields.io/github/license/palp1tate/go-crypto-guard?style=flat-square" alt="license">
    </a>
    <a href="https://github.com/palp1tate/go-crypto-guard/issues">
      <img src="https://img.shields.io/github/issues/palp1tate/go-crypto-guard?style=flat-square" alt="GitHub issues">
    </a>
    <a href="#">
      <img src="https://img.shields.io/github/stars/palp1tate/go-crypto-guard?style=flat-square" alt="GitHub stars">
    </a>
    <a href="https://github.com/palp1tate/go-crypto-guard/network">
      <img src="https://img.shields.io/github/forks/palp1tate/go-crypto-guard?style=flat-square" alt="GitHub forks">
    </a>
    <a href="https://github.com/palp1tate/go-crypto-guard/releases/latest">
      <img src="https://img.shields.io/github/release/palp1tate/go-crypto-guard.svg" alt="Release">
    </a>
    <a href=https://goreportcard.com/report/github.com/palp1tate/go-crypto-guard>
        <img src="https://goreportcard.com/badge/github.com/palp1tate/go-crypto-guard" alt="go report">
    </a>
    <a href="#">
      <img src="https://img.shields.io/github/languages/top/palp1tate/go-crypto-guard" alt="language">
    </a>
    <a href="#">
      <img src="https://img.shields.io/github/last-commit/palp1tate/go-crypto-guard" alt="last commit">
    </a>
   <a href="#">
      <img src="https://komarev.com/ghpvc/?username=go-crypto-guard&label=Views&color=0e75b6&style=flat" alt="访问量统计" />
    </a>
</div>

## Language

- [English](https://github.com/palp1tate/go-crypto-guard/blob/main/README.md)
- [中文](https://github.com/palp1tate/go-crypto-guard/blob/main/README_CN.md)

## Introduction

This repository contains a comprehensive password hashing library written in Go. The library supports multiple hashing algorithms,it allows for customizable salt length, iterations, key length, and algorithm selection. This open-source
project aims to provide developers with a versatile tool for secure password storage and validation.

Algorithms supported:

- [SHA512](https://medium.com/@zaid960928/cryptography-explaining-sha-512-ad896365a0c1)
- [SHA384](https://medium.com/@zaid960928/cryptography-explaining-sha-512-ad896365a0c1)
- [SHA256](https://golden.com/wiki/SHA-256-XKEJ8AB)
- [SHA1](https://bing.com/search?q=SHA1+algorithm+Wikipedia)
- [Md5](https://en.wikipedia.org/wiki/MD5)
- [HMAC](https://en.wikipedia.org/wiki/HMAC)
- [Argon2](https://bing.com/search?q=Argon2+algorithm+Wikipedia)
- [Bcrypt](https://en.wikipedia.org/wiki/Bcrypt)
- [Scrypt](https://en.wikipedia.org/wiki/Scrypt)
- [Blake2b](https://en.wikipedia.org/wiki/Comparison_of_cryptographic_hash_functions)
- [Blake2s](https://en.wikipedia.org/wiki/Comparison_of_cryptographic_hash_functions)
- AES
- DES
- 3DES
- RSA
- RC4
- Blowfish
- ECC

The format of the some `passwords` is same as the encryption algorithm format that comes with [Django](https://www.djangoproject.com/):

```go
<algorithm>$<iterations>$<salt>$<hash>
```

others may be like:

```go
<algorithm>$<hash>
```

## Installation

```bash
go get -u github.com/palp1tate/go-crypto-guard 
```

## Usage

Some examples of usage are provided below:

### SHA512

```go
// SHA512 encrypts a password using PBKDF2 and SHA-512.
// It takes a password, salt length, key length, and iterations as input.If you pass in an invalid value, the function takes the default value.
// It generates a salt, derives a key using PBKDF2 and SHA-512, and returns the encrypted password.
//The format of password:<algorithm>$<iterations>$<salt>$<hash>
//pbkdf2_sha512$100$40fde046f66c1d9e55b4435d$1fdd34c50a98e576b612d66be507f019

password := "12345678"
encodedPassword, _ := SHA512(password, 12, 16, 100)
ok, _ := VerifySHA512(password, encodedPassword)
```

**The use of SHA384、SHA256、SHA1、Md5 and Argon2 are the same as for SHA512**.

### HMAC

```go
// HMAC encrypts a password using HMAC and SHA-256.
// It takes a password and salt length as input.
// It generates a salt, computes the HMAC of the password using the salt and SHA-256, and returns the encrypted password.
//The format of password:<algorithm>$<salt>$<hash>
//hmac$3bf4e2c1a9ed54575d0d1f937eb363ab$a6ed73f8fe48867db2bd58c69ebe6c0fb91ecdd8147c4352fecf018d07cb4f43

password := "12345678"
encodedPassword, _ := HMAC(password, 16)
ok, _ := VerifyHMAC(password, encodedPassword)
```

### Bcrypt

```go
// Bcrypt encrypts a password using the Bcrypt hashing function.
// It takes a password as input, generates a hash from the password using Bcrypt's default cost, and returns the encrypted password.
//The format of password:<algorithm>$<hash>
//bcrypt$243261243130246769545174546869684f565835616a694a4e3578432e6e387a4c426451526932692e443067756758334a436d3532717365784e5661

password := "12345678"
encodedPassword, _ := Bcrypt(password)
ok, _ := VerifyBcrypt(password, encodedPassword)
```

**The use of Blake2b、Blake2s、 are the same as for Bcrypt**.

### Scrypt

```go
// Scrypt encrypts a password using the Scrypt key derivation function.
// It takes a password, salt length, and key length as input.
// It generates a salt, derives a key using Scrypt and the provided parameters, and returns the encrypted password.
//The format of password:<algorithm>$<salt>$<hash>
//scrypt$679a0a3c8336a9ff36b809862e7d494c$c4cec5ca742fa984045457f76d217acf245f032251c6a3952c4d68e1cba4a488

password := "12345678"
encodedPassword, _ := Scrypt(password, 16, 32)
ok, _ := VerifyScrypt(password, encodedPassword)
```

### AES

```go
// AES encrypts a password using the AES encryption algorithm.
// It takes a password and an AES key as input.
// It creates a new cipher block from the AES key, applies PKCS7 padding to the password, and encrypts the password using CBC mode.
// It returns the encrypted password.
//The format of password:<algorithm>$<hash>
//aes$BhV9oJiePwpsEwDWizJoCA==

password := "12345678"
//the length of aes key must be 32
aesKey := "palpitateabcdefghijklmn123456789"
encodedPassword, _ := AES(password, aesKey)
ok, _ := VerifyAES(password, encodedPassword, aesKey)
```

**The use of DES 、ThreeDES、RC4 and Blowfish are the same as for Bcrypt.For DES,the length of des key must be 8.For ThreeDES,the length of threedes key must be 24.There is no limit to the length of the rc4Key and blowfishKey ,but  for Blowfish, the length of password must be 8.**

### RSA

```go
// GenRSAKey generates a pair of RSA keys and saves them to files.
// It takes the number of bits for the key as input.2048 or 4096 is recommended.
// It generates a private key and a public key, and writes them to "privateKey.pem" and "publicKey.pem" respectively.


// RSA encrypts a password using the RSA encryption algorithm.
// It takes a password and the path to a public key file as input.
// It reads the public key from the file, encrypts the password using RSA and PKCS1v15 padding, and returns the encrypted password.
//The format of password:<algorithm>$<hash>
//rsa$3p1+X80iFIDtwtKOQFjXm+deyv+cxkEIbpXuwXcqbcCvean6zyWvcrogQtDj2MkYOE2ScHpARR93RYxs3y+RXetKAHhrDqWURYcyJwuTwShBmR4hz+3WkFzhqm44IgPdlgdt70uO7TXx6fj1WmUTsZpNDTF/WNdEUO7Rzc8wahYBcnMOnPgUXrnUCYRSX7OBjuLwThnd9FTgh8CdaqESHWh6UPgkj9xz3G2uRplx2Tae0Pbsk8vQTuJXsqT//Q8yoC+ELo+5S6wTE6H8AMBdgvJgNHzFDldQD8UsZ7Ta/u2uF/joHwBA6V6IS4+1ithspE9ceJZCBWo2Cj6fMIbvjg==

//Before you can encrypt a password, you must first generate a pair of keys.This function can be called only once, remembering that the same key pair is required when verifying the password.
_ = GenRSAKey(2048)	//It only needs to be called once
password := "12345678"
encodedPassword, _ := RSA(password, "publicKey.pem")
ok, _ := VerifyRSA(password, encodedPassword, "privateKey.pem")
```

### ECC

```go
// ECC encrypts a password using the ECC encryption algorithm.
// It takes a password and a private key as input.
// It computes the SHA-256 digest of the password, signs the digest using the private key, and returns the encrypted password.
//The format of password:<algorithm>$<hash>
//ecc$BQOoQvBhRHKi9GsV0qpPiyMJ5hRwdiXlQL7CcMsPCo1GvIomtb8xzjNnmq7RNRWmS9AKXo+i0Cg4fmAdLeCN8w==


password := "12345678"
privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
encodedPassword, _ := ECC(password, privateKey)
publicKey := privateKey.PublicKey
ok, _ := VerifyECC(password, encodedPassword publicKey)
```



## Contribute

Welcome contributions to the repository. Here are a few ways you can help:

1. **Report bugs**: If you encounter any issues or bugs, please open an issue on the GitHub repository.
2. **Suggest enhancements**: If you have ideas for new features or improvements, feel free to open an issue detailing
   your suggestion.
3. **Submit pull requests**: If you’ve fixed a bug or developed a new feature, we’d love to see it. Please submit a pull
   request with your changes.

Before contributing, please make sure to read and follow our code of conduct and contribution guidelines (if available).

## License

This project is licensed under the Apache License 2.0. See
the [Apache License 2.0](https://github.com/palp1tate/go-crypto-guard/blob/main/LICENSE) file for more details.
