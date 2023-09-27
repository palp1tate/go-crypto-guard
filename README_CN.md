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

## 介绍

此存储库包含用 Go 编写的全面的密码哈希库。该库支持多种哈希算法，它允许可定制的盐长度、迭代、键长度和算法选择。这个开源项目旨在为开发人员提供一个多功能的工具，用于安全的密码存储和验证。

支持的算法:

- [SHA512](https://github.com/palp1tate/go-crypto-guard/blob/main/pbkdf2/sha512.go)
- [SHA384](https://github.com/palp1tate/go-crypto-guard/blob/main/pbkdf2/sha384.go)
- [SHA256](https://github.com/palp1tate/go-crypto-guard/blob/main/pbkdf2/sha256.go)
- [SHA1](https://github.com/palp1tate/go-crypto-guard/blob/main/pbkdf2/sha1.go)
- [Md5](https://github.com/palp1tate/go-crypto-guard/blob/main/pbkdf2/md5.go)
- [HMAC](https://github.com/palp1tate/go-crypto-guard/blob/main/hmac/hmac.go)
- [Argon2](https://github.com/palp1tate/go-crypto-guard/blob/main/argon2/argon2.go)
- [Bcrypt](https://github.com/palp1tate/go-crypto-guard/blob/main/bcrypt/bcrypt.go)
- [Scrypt](https://github.com/palp1tate/go-crypto-guard/blob/main/scrypt/scrypt.go)
- [Blake2b](https://github.com/palp1tate/go-crypto-guard/blob/main/blake2b/blake2b.go)
- [Blake2s](https://github.com/palp1tate/go-crypto-guard/blob/main/blake2s/blake2s.go)
- [AES](https://github.com/palp1tate/go-crypto-guard/blob/main/aes/aes.go)
- [DES](https://github.com/palp1tate/go-crypto-guard/blob/main/des/des.go)
- [3DES](https://github.com/palp1tate/go-crypto-guard/blob/main/3des/3des.go)
- [RSA](https://github.com/palp1tate/go-crypto-guard/blob/main/rsa/rsa.go)
- [RC4](https://github.com/palp1tate/go-crypto-guard/blob/main/rc4/rc4.go)
- [Blowfish](https://github.com/palp1tate/go-crypto-guard/blob/main/blowfish/blowfish.go)
- [ECC](https://github.com/palp1tate/go-crypto-guard/blob/main/ecc/ecc.go)

一些加密过后的密码格式与[Django](https://www.djangoproject.com/)内置的加密算法格式相同:

```go
<algorithm>$<iterations>$<salt>$<hash>
```

其他可能的格式:

```go
<algorithm>$<hash>
```

## 安装

```go
go get -u github.com/palp1tate/go-crypto-guard 
```

## 用法

### SHA512

```go
// SHA512 使用 PBKDF2 和 SHA-512 对密码进行加密。
// 它接受密码、盐长度、密钥长度和迭代次数作为输入。如果你传入一个无效的值，函数将采取默认值。
// 它生成一个盐，使用 PBKDF2 和 SHA-512 派生一个密钥，并返回加密的密码。
// 密码的格式：<algorithm>$<iterations>$<salt>$<hash>
//pbkdf2_sha512$100$40fde046f66c1d9e55b4435d$1fdd34c50a98e576b612d66be507f019

password := "12345678"
encodedPassword, _ := pwd.GenSHA512(password, 12, 16, 100)
ok, _ := pwd.VerifySHA512(password, encodedPassword)
```

**SHA384、 SHA256、 SHA1、 Md5和 Argon2的用法与 SHA512相同**.

### HMAC

```go
// HMAC 使用 HMAC 和 SHA-256 对密码进行加密。
// 它接受密码和盐长度作为输入。
// 它生成一个盐，使用盐和 SHA-256 计算密码的 HMAC，并返回加密的密码。
// 密码的格式：<algorithm>$<salt>$<hash>
//hmac$3bf4e2c1a9ed54575d0d1f937eb363ab$a6ed73f8fe48867db2bd58c69ebe6c0fb91ecdd8147c4352fecf018d07cb4f43

password := "12345678"
encodedPassword, _ := pwd.GenHMAC(password, 16)
ok, _ := pwd.VerifyHMAC(password, encodedPassword)
```

### Bcrypt

```go
// Bcrypt 使用 Bcrypt 哈希函数对密码进行加密。
// 它接受一个密码作为输入，使用 Bcrypt 的默认成本从密码生成一个哈希，并返回加密的密码。
// 密码的格式：<algorithm>$<hash>
//bcrypt$243261243130246769545174546869684f565835616a694a4e3578432e6e387a4c426451526932692e443067756758334a436d3532717365784e5661

password := "12345678"
encodedPassword, _ := pwd.GenBcrypt(password)
ok, _ := pwd.VerifyBcrypt(password, encodedPassword)
```

**对 Blake2b、 Blake2s 的使用与对 Bcrypt 的使用相同**.

### Scrypt

```go
// Scrypt 使用 Scrypt 密钥派生函数对密码进行加密。
// 它接受一个密码、盐长度和密钥长度作为输入。
// 它生成一个盐，使用 Scrypt 和提供的参数派生一个密钥，并返回加密的密码。
// 密码的格式：<algorithm>$<salt>$<hash>
//scrypt$679a0a3c8336a9ff36b809862e7d494c$c4cec5ca742fa984045457f76d217acf245f032251c6a3952c4d68e1cba4a488

password := "12345678"
encodedPassword, _ := pwd.GenScrypt(password, 16, 32)
ok, _ := pwd.VerifyScrypt(password, encodedPassword)
```

### AES

```go
// AES 使用 AES 加密算法对密码进行加密。
// 它接受一个密码和一个 AES 密钥作为输入。
// 它从 AES 密钥创建一个新的密码块，对密码应用 PKCS7 填充，并使用 CBC 模式加密密码。它返回加密的密码。
// 密码的格式：<algorithm>$<hash>
// aes$BhV9oJiePwpsEwDWizJoCA==

password := "12345678"
//aes key的长度必须为32
aesKey := "palpitateabcdefghijklmn123456789"
encodedPassword, _ := pwd.GenAES(password, aesKey)
ok, _ := pwd.VerifyAES(password, encodedPassword, aesKey)
```

**DES、ThreeDES、RC4和Blowfish的使用与Bcrypt相同，对于DES，desKey的长度必须为8。对于ThreeDES，threeDesKey的长度必须为24。rc4Key和BlowfishKey的长度没有限制，但对于Blowfish，密码的长度必须为8。**

### RSA

```go
// GenRSAKey 生成一对 RSA 密钥并将它们保存到文件中。 
// 它接受密钥的位数作为输入。推荐使用 2048 或 4096。 
// 它生成一个私钥和一个公钥，并分别将它们写入 “privateKey.pem” 和 “publicKey.pem”。

// RSA 使用 RSA 加密算法对密码进行加密。 
// 它接受一个密码和公钥文件的路径作为输入。 
// 它从文件中读取公钥，使用 RSA 和 PKCS1v15 填充对密码进行加密，并返回加密的密码。 
// 密码的格式：<algorithm>$<hash> 
//rsa$3p1+X80iFIDtwtKOQFjXm+deyv+cxkEIbpXuwXcqbcCvean6zyWvcrogQtDj2MkYOE2ScHpARR93RYxs3y+RXetKAHhrDqWURYcyJwuTwShBmR4hz+3WkFzhqm44IgPdlgdt70uO7TXx6fj1WmUTsZpNDTF/WNdEUO7Rzc8wahYBcnMOnPgUXrnUCYRSX7OBjuLwThnd9FTgh8CdaqESHWh6UPgkj9xz3G2uRplx2Tae0Pbsk8vQTuJXsqT//Q8yoC+ELo+5S6wTE6H8AMBdgvJgNHzFDldQD8UsZ7Ta/u2uF/joHwBA6V6IS4+1ithspE9ceJZCBWo2Cj6fMIbvjg==

// 在你可以加密密码之前，你必须先生成一对密钥。这个函数只能被调用一次，记住在验证密码时需要相同的密钥对。

_ = pwd.GenRSAKey(2048)	//只需要执行一次就可以注释掉
password := "12345678"
encodedPassword, _ := pwd.GenRSA(password, "publicKey.pem")
ok, _ := pwd.VerifyRSA(password, encodedPassword, "privateKey.pem")
```

### ECC

```go
// ECC 使用 ECC 加密算法对密码进行加密。
// 它接受一个密码和一个私钥作为输入。
// 它计算密码的 SHA-256 摘要，使用私钥对摘要进行签名，并返回加密的密码。
// 密码的格式：<algorithm>$<hash>
//ecc$BQOoQvBhRHKi9GsV0qpPiyMJ5hRwdiXlQL7CcMsPCo1GvIomtb8xzjNnmq7RNRWmS9AKXo+i0Cg4fmAdLeCN8w==


password := "12345678"
privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
encodedPassword, _ := pwd.GenECC(password, privateKey)
publicKey := privateKey.PublicKey
ok, _ := pwd.VerifyECC(password, encodedPassword publicKey)
```

## 贡献

欢迎贡献。您可以通过以下几种方式提供帮助：

1. **报告错误**：如果您遇到任何问题或错误，请在 GitHub 存储库上提出问题。
2. **建议增强功能**：如果您对新功能或改进有任何想法，请随时提出一个问题，详细说明您的建议。
3. **提交拉取请求**：如果您修复了错误或开发了新功能，我们很乐意看到它。请提交包含您的更改的拉取请求。

在贡献之前，请务必阅读并遵守我们的行为准则和贡献指南（如果有）。

## 开源协议

该项目根据 Apache License 2.0
获得许可。有关更多详细信息，请参阅[Apache 许可证 2.0](https://github.com/palp1tate/go-crypto-guard/blob/main/LICENSE)文件。
