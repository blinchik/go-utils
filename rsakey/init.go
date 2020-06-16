package rsakey

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
)

var bitSize int
var Home string
var err error

var Key *rsa.PrivateKey
var PublicKey rsa.PublicKey

func init() {

	Home, err = os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
	}

	reader := rand.Reader

	bitSize = 4096
	Key, err = rsa.GenerateKey(reader, bitSize)
	PublicKey = Key.PublicKey

	checkError(err)

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
