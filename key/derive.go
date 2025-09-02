package key

import (
	"context"
	"crypto/pbkdf2"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sloth/db"
)

func DeriveMasterKey(username string, password string) ([]byte, error) {
	conn := db.Connect()
	row, _ := conn.Query(context.Background(), "Select salt from users where username =  $1 and password = $2", username, password)
	// var pass string
	var salt string
	for row.Next() {
		// row.Scan(&pass)
		row.Scan(&salt)
	}
	// fmt.Println(pass)
	fmt.Println(salt)

	rawsalt, _ := hex.DecodeString(salt)

	masterkey, err := pbkdf2.Key(sha256.New, password, rawsalt, 100000, 32)
	fmt.Println("MasterKey:", masterkey)
	return masterkey, err
}
