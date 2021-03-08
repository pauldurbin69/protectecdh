// +build windows

package protectecdh

import (
	"crypto/ecdsa"
	"io/ioutil"

	"github.com/billgraziano/dpapi"
	"github.com/pauldurbin69/ecdh"
)

// func protectKey(privateKey *ecdsa.PrivateKey) error {

func protectKey(privateKey *ecdsa.PrivateKey) error {

	bytes, err := ecdh.EncodeEcPrivateKey(privateKey, getKey())

	cipherBytes, err := dpapi.EncryptBytesMachineLocal(bytes)

	err = ioutil.WriteFile(getKeyFilePath(), cipherBytes, 0600)

	return err
}

// unprotectKey
func unprotectKey() (*ecdsa.PrivateKey, error) {

	cipherBytes, err := ioutil.ReadFile(getKeyFilePath())

	encPriv, err := dpapi.DecryptBytes(cipherBytes)

	privateKey, err := ecdh.DecodeEcPrivateKey(encPriv, getKey())

	return privateKey, err
}
