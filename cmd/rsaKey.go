package main

import (
	"flag"
	"fmt"
	"os"

	rsaKey "github.com/blinchik/go-utils/rsakey"
)

func main() {

	local := flag.Bool("local", false, "local")
	flag.Parse()

	if *local {
		keyName := os.Args[2]

		rsaKey.SavePEMKey(fmt.Sprintf("%s/.ssh/%s.pem", rsaKey.Home, keyName), rsaKey.Key)
		rsaKey.SavePublicPEMKey(fmt.Sprintf("%s/.ssh/%s.pub", rsaKey.Home, keyName), rsaKey.PublicKey)

	}

}
