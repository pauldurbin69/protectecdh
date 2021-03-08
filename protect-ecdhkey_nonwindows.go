// +build freebsd netbsd openbsd dragonfly solaris darmin linux

package protectecdh

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"os"

	"github.com/denisbrodbeck/machineid"
	"github.com/pauldurbin69/ecdh"
)

const (
	appKey      = "h&ji(_8G$$hhukkwy56"
	keyFileName = ".ecdh-key"
)

// Protect
func protectKey(privateKey *ecdsa.PrivateKey) error {

	return ecdh.SaveEcdhKeyToFile(privateKey, getKey(), getKeyFilePath())
}

// UnprotectKey
func unprotectKey() (*ecdsa.PrivateKey, error) {

	privateKey, err := ecdh.ReadEcdhKeyFromFile(getKey(), getKeyFilePath())

	return privateKey, err
}

func getKey() ([]byte, error) {

	id, err := machineid.ProtectedID(appKey)
	if err != nil {
		log.Fatal(err)
	}

	key, err := hex.DecodeString(id)

	return key, err
}

func getKeyFilePath() (string, error) {

	home, err := os.UserHomeDir()

	if err != nil {
		return "", nil
	}

	return home + "/" + keyFileName, nil
}
