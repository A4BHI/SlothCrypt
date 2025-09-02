package login

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Login() {
	salt := make([]byte, 32)
	rand.Read(salt)
	fmt.Print(salt)

	strsalt := hex.EncodeToString(salt)
	fmt.Println(strsalt)
	orgsalt, _ := hex.DecodeString(strsalt)
	fmt.Println(orgsalt)

}
