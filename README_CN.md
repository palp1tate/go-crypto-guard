<h1 align="center" style="border-bottom: none;">go-crypto-guard </h1>

<div class="labels" align="center">
    <a href="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
      <img src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg" alt="semantic-release">
    </a>
    <a href="https://pkg.go.dev/github.com/palp1tate/go-crypto-guard/v2">
      <img src="https://godoc.org/github.com/palp1tate/go-crypto-guard?status.svg" alt="Godoc">
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

该存储库包含一个用 Go 编写的综合密码哈希库。该库支持多种哈希算法，包括 PBKDF2（使用 SHA1、SHA256、SHA384、SHA512 和 MD5）、Bcrypt、Scrypt、Argon2、HMAC、Blake2b 和 Blake2s。它允许自定义盐长度、迭代、密钥长度和算法选择。该开源项目旨在为开发人员提供用于安全密码存储和验证的多功能工具。

支持的算法:

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

password的格式与[Django](https://www.djangoproject.com/)内置的加密算法格式相同:

```go
<algorithm>$<iterations>$<salt>$<hash>
```

## 安装

```
go get github.com/palp1tate/go-crypto-guard 
```

## 用法

下面提供了一些用法示例：

```go
package main

import (
	"fmt"
	"github.com/palp1tate/go-crypto-guard"
)

func main() {
	originPwd := "123456"
	options := pwd.Options{
		SaltLen:    16,
		KeyLen:     32,
		Iterations: 100,
		Algorithm:  pwd.SHA512,
	}
	encodedPwd, err := pwd.Generate(originPwd, &options)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Encoded password:", encodedPwd)

	if ok, err := pwd.Verify(originPwd, encodedPwd); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Verify result:", ok)
	}
}
```

对于SHA512、SHA256、SHA1、SHA384、Md5、Argon2，可以填写全部参数，也可以不完全填写。但对于其他算法，它们不需要那么多参数，你甚至可以只用指定具体的算法：

```go
//Bcrypt
options := pwd.Options{
		Algorithm: pwd.Bcrypt,
	}

//HMAC
options := pwd.Options{
		Algorithm: pwd.HMAC,
	}

//...
```

`Options`定义用于自定义密码散列过程的参数。每个字段都有一个默认值，即使您不传递参数也是如此。

```go
// Fields:
//   - SaltLen: Length of the salt to be generated for password hashing.
//   - Iterations: Number of iterations to apply during the hashing process.
//   - KeyLen: Length of the derived key produced by the hashing algorithm.
//   - Algorithm: The specific hashing algorithm to be used for password hashing.
type Options struct {
	SaltLen    int    //  Defaults to 16.
	Iterations int    //  Defaults to 50.
	KeyLen     int    //  Defaults to 32.
	Algorithm  Algorithm //  Defaults to "SHA512".
}
```

## 贡献

欢迎贡献。您可以通过以下几种方式提供帮助：

1. **报告错误**：如果您遇到任何问题或错误，请在 GitHub 存储库上提出问题。
2. **建议增强功能**：如果您对新功能或改进有任何想法，请随时提出一个问题，详细说明您的建议。
3. **提交拉取请求**：如果您修复了错误或开发了新功能，我们很乐意看到它。请提交包含您的更改的拉取请求。

在贡献之前，请务必阅读并遵守我们的行为准则和贡献指南（如果有）。

## 未来的计划

我们计划在未来的版本中加入更多的哈希算法，以满足不同的场景和需求。以下是我们可能考虑的一些算法：

- [RSA](https://en.wikipedia.org/wiki/RSA_(cryptosystem))
- [DES](https://en.wikipedia.org/wiki/Data_Encryption_Standard)
- [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)
- ……

请注意，这只是一个计划，可能会根据项目需求和社区反馈进行更改。我们将通过 GitHub 存储库向用户通报任何更改或添加的最新情况。

## 开源协议

该项目根据 Apache License 2.0 获得许可。有关更多详细信息，请参阅[Apache 许可证 2.0](https://github.com/palp1tate/go-crypto-guard/blob/main/LICENSE)文件。
