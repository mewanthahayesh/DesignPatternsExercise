package main

import (
	"crypto/rsa"
	"encoding/asn1"
	"encoding/base64"
	"fmt"
)

func main() {
	publicKeyBase64 := "MIGJAoGBAJJYXgBem1scLKPEjwKrW8+ci3B/YNN3aY2DJ3lc5e2wNc0SmFikDpow1TdYcKl2wdrXX7sMRsyjTk15IECMezyHzaJGQ9TinnkQixJ+YnlNdLC04TNWOg13plyahIXBforYAjYl2wVIA8Yma2bEQFhmAFkEX1A/Q1dIKy6EfQ+xAgMBAAE="

	// Base64 decode.
	publicKeyBinary, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		panic(err)
	}

	// rsa.PublicKey is a big.Int (N: modulus) and an integer (E: exponent).
	var pubKey rsa.PublicKey
	if rest, err := asn1.Unmarshal(publicKeyBinary, &pubKey); err != nil {
		panic(err)
	} else if len(rest) != 0 {
		panic("rest is not nil")
	}

	fmt.Printf("key: %+v\n", pubKey)
}
