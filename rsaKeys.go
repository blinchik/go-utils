package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
)

func main() {

	local := flag.Bool("local", false, "local")
	flag.Parse()

	reader := rand.Reader
	bitSize := 4096

	home, err := os.UserHomeDir()

	if err != nil {
		fmt.Println(err)
	}

	key, err := rsa.GenerateKey(reader, bitSize)
	checkError(err)

	publicKey := key.PublicKey

	if *local {
		keyName := os.Args[1]

		savePEMKey(fmt.Sprintf("%s/.ssh/%s.pem", home, keyName), key)
		savePublicPEMKey(fmt.Sprintf("%s/.ssh/%s.pub", home, keyName), publicKey)

	}

}

func savePEMKey(fileName string, key *rsa.PrivateKey) {

	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	err = os.Chmod(fileName, 0400)
	if err != nil {
		fmt.Println(err)
	}

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

func keepPEMKey(key *rsa.PrivateKey) []byte {

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	pem := pem.EncodeToMemory(privateKey)

	return pem

}

func keepPublicPEMKey(pubkey rsa.PublicKey) []byte {
	asn1Bytes, err := asn1.Marshal(pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pub := pem.EncodeToMemory(pemkey)
	return pub
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
