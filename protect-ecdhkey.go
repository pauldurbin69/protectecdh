package protectecdh

import (
	"crypto/ecdsa"
)

// Protect write ecdh private key to OS secured file
func Protect(privateKey *ecdsa.PrivateKey) error {

	return protectKey(privateKey)
}

// Unprotect read ecdh private key from OS secured file
func Unprotect() (*ecdsa.PrivateKey, error) {

	return unprotectKey()
}
