package rsakey

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func SavePEMKey(fileName string, key *rsa.PrivateKey) {

	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)

	err = os.Chmod(fileName, 0400)
	if err != nil {
		fmt.Println(err)
	}

	checkError(err)
}

func SavePublicPEMKey(fileName string, pubkey rsa.PublicKey) {
	pubN, err := ssh.NewPublicKey(&pubkey)
	if err != nil {
		log.Fatal(pubN)
	}
	pubBytes := ssh.MarshalAuthorizedKey(pubN)

	// write the whole body at once
	err = ioutil.WriteFile(fileName, pubBytes, 0640)
	if err != nil {
		log.Fatal(err)
	}

}

func KeepPEMKey(key *rsa.PrivateKey) []byte {

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	pem := pem.EncodeToMemory(privateKey)

	return pem

}

func KeepPublicPEMKey(pubkey rsa.PublicKey) []byte {
	pubN, err := ssh.NewPublicKey(&pubkey)
	if err != nil {
		log.Fatal(pubN)
	}
	pubBytes := ssh.MarshalAuthorizedKey(pubN)

	return pubBytes
}
