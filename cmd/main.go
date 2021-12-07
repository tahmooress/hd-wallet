package main

import (
	"bitcoin-wallet/internal/mnemonic"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("already have wallet? type 'yes' to go to next step to import it,\n" +
		"otherwise enter any key to generate your recovery phrase")
	fmt.Println("------------------------------------------------------------------------------------------")

	answer, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("readString from os.Stdin reader fail: %s", err)
	}

	if answer != "yes" {
		fmt.Println("choose your security length 128 or 256?")
		fmt.Println("*** longer length will bring more security ***")

		n, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("readString from os.Stdin reader fail: %s", err)
		}

		size, err := strconv.Atoi(strings.TrimSpace(n))
		if err != nil {
			log.Fatal(err)
		}

		for size != 128 && size != 256 {
			fmt.Println("security length should be either 128 or 256")

			n, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalf("readString from os.Stdin reader fail: %s", err)
			}

			size, err = strconv.Atoi(strings.TrimSpace(n))
			if err != nil {
				log.Fatal(err)
			}
		}

		recoveryPhrase, err := mnemonic.GenerateMnemonic(size)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("*** write the following phrase combination and keep it safe! ***")
		fmt.Println("*** this phrase is your recovery key and without it you will never be able to recover your wallet. ***")
		fmt.Println("----------------------------------------------------------------------------------------")
		fmt.Println("your recovery phrase: ")
		fmt.Println(recoveryPhrase)
		fmt.Println("*** KEEP IT SECRET AND NEVER SHARE IT WITH ANYONE ***")

		os.Exit(0)
	}

	//iW := hd_wallet.New()
	//
	//wallet, err := iW.GenerateKeyPair(strings.NewReader(strings.TrimSpace(seedWords)))
	//if err != nil {
	//	log.Fatalf("generating pub/priv key pair from seed words fail: %s", err)
	//}

	//fmt.Printf("your private key is : %X\n", wallet.PrivateKey)
	//fmt.Printf("x = %X \n y = %X \n", wallet.PublicKey.X, wallet.PublicKey.Y)
	//
	//fmt.Println(m)

	os.Exit(0)
}
