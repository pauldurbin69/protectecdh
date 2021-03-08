package protectecdh

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"os"

	"github.com/pauldurbin/ecdh"

	"github.com/denisbrodbeck/machineid"
)

const (
	appKey      = "h&ji(_8G$$hhukkwy56"
	keyFileName = ".ecdh-key"
)

// Protect
func protectKeyNonWindows(privateKey *ecdsa.PrivateKey) error {

	key, err := getKey()
	home, err := os.UserHomeDir()

	err = ecdh.SaveEcdhKeyToFile(privateKey, key, home+"/"+keyFileName)

	return err
}

// UnprotectKey
func unprotectKeyNonWindows() (*ecdsa.PrivateKey, error) {

	key, err := getKey()
	home, err := os.UserHomeDir()
	privateKey, err := ecdh.ReadEcdhKeyFromFile(key, home+"/"+keyFileName)

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
