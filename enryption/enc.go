package enryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"sloth/key"
)

func EncryptFileKey(filekey []byte, masterkey []byte) []byte {
	block, err := aes.NewCipher(masterkey)
	if err != nil {
		fmt.Println("Error creating block for filekey:", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error Creating GCM for Filekey:", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	rand.Read(nonce)

	encfilekey := gcm.Seal(nil, nonce, filekey, nil)

	return encfilekey

}
func Encryption(username string, password string) {
	filekey := make([]byte, 32)
	rand.Read(filekey)
	fmt.Println("Orginal Filekey := ", filekey)
	masterkey, _ := key.DeriveMasterKey(username, password)

	encFilekey := EncryptFileKey(filekey, masterkey)
	fmt.Println("Encrypted Filekey:=", encFilekey)

}
