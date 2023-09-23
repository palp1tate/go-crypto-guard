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

## Language

- [English](https://github.com/palp1tate/go-crypto-guard/blob/main/README.md)
- [中文](https://github.com/palp1tate/go-crypto-guard/blob/main/README_CN.md)

## Introduction

This repository contains a comprehensive password hashing library written in Go. The library supports multiple hashing algorithms including PBKDF2 (with SHA1, SHA256, SHA384, SHA512, and MD5), Bcrypt, Scrypt, Argon2, HMAC, Blake2b, and Blake2s. It allows for customizable salt length, iterations, key length, and algorithm selection. This open-source project aims to provide developers with a versatile tool for secure password storage and validation.

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

## Installation

```
go get github.com/palp1tate/go-crypto-guard 
```

## Usage

Some examples of usage are provided below:

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

For SHA512,SHA256,SHA1,SHA384,Md5,Argon2,you can fill in all the parameters or not.

But for other algorithms, they don't need so many parameters,you can even specify only what the algorithm is:

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

`Options` defines parameters for customizing the password hashing process.  Each field has a default value, even if you don't pass parameters.

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
	Algorithm  string //  Defaults to "SHA512".
}
```

## Contribute

Welcome contributions to the repository. Here are a few ways you can help:

1. **Report bugs**: If you encounter any issues or bugs, please open an issue on the GitHub repository.
2. **Suggest enhancements**: If you have ideas for new features or improvements, feel free to open an issue detailing your suggestion.
3. **Submit pull requests**: If you’ve fixed a bug or developed a new feature, we’d love to see it. Please submit a pull request with your changes.

Before contributing, please make sure to read and follow our code of conduct and contribution guidelines (if available).

## Future Plans

We plan to incorporate more hashing algorithms in future versions to cater to different scenarios and requirements. Here are some algorithms that we might consider:

- [RSA](https://en.wikipedia.org/wiki/RSA_(cryptosystem))
- [DES](https://en.wikipedia.org/wiki/Data_Encryption_Standard)
- [AES](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)
- …

Please note that this is just a plan and might change according to project needs and community feedback. We will keep our users updated with any changes or additions through our GitHub repository.

## License

This project is licensed under the Apache License 2.0. See the [Apache License 2.0](https://github.com/palp1tate/go-crypto-guard/blob/main/LICENSE) file for more details.
