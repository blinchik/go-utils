package main

import (
	"fmt"
	"os"
)

func main() {

	if *rsaKey.local {
		keyName := os.Args[1]

		rsaKey.savePEMKey(fmt.Sprintf("%s/.ssh/%s.pem", home, keyName), key)
		rsaKey.savePublicPEMKey(fmt.Sprintf("%s/.ssh/%s.pub", home, keyName), publicKey)

	}

}
