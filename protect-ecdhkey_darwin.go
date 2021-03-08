// +build darwin

package protectecdh

import (
	"crypto/ecdsa"
)

// Protect
func protectKey(privateKey *ecdsa.PrivateKey) error {

	return protectKeyNonWindows(privateKey)
}

// UnprotectKey
func unprotectKey() (*ecdsa.PrivateKey, error) {

	return unprotectKeyNonWindows()
}
