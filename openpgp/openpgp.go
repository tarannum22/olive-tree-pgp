package openpgp

import (
	"os"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

func DecryptFile(fileName string, privateKeyFile string, outputFilename string) {

	// Read a encrypted file
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// Read a private key file
	key, err := os.ReadFile(privateKeyFile)
	if err != nil {
		panic(err)
	}

	privateKeyObj, _ := crypto.NewKeyFromArmored(string(key))
	unlockedKeyObj, _ := privateKeyObj.Unlock([]byte(""))
	privateKeyRing, _ := crypto.NewKeyRing(unlockedKeyObj)

	pgpData := crypto.NewPGPMessage(fileData)

	message, _ := privateKeyRing.Decrypt(pgpData, nil, 0)
	os.WriteFile(outputFilename, message.Data, 0644)
}

func EncryptFile(fileName string, publicKeyFile string, outputFilename string) {

	// Read a plain file
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// Read a public key file
	key, err := os.ReadFile(publicKeyFile)
	if err != nil {
		panic(err)
	}

	publicKeyObj, _ := crypto.NewKeyFromArmored(string(key))
	publicKeyRing, _ := crypto.NewKeyRing(publicKeyObj)
	binMessage := crypto.NewPlainMessage(fileData)
	message, _ := publicKeyRing.Encrypt(binMessage, nil)
	os.WriteFile(outputFilename, message.Data, 0644)
}