package mnemonic

import (
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"golang.org/x/crypto/pbkdf2"
	"strconv"
	"strings"
)

var ErrCheckSumSize = errors.New("something goes wrong during creating checksum")

const defaultSize = 128

// GenerateMnemonic takes the size parameter and return a mnemonic sentence.
// size should be between 128 - 256 bits
func GenerateMnemonic(size int) (string, error) {
	if size < defaultSize || size > defaultSize*2 || size%32 != 0 {
		size = defaultSize
	}

	entropy := make([]byte, size/8)

	_, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}

	withCheckSum, err := addCheckSum(byte2bit(entropy))
	if err != nil {
		return "", err
	}

	var mnemonic string

	for i := 0; i < len(withCheckSum)/11; i++ {
		n := withCheckSum[i*11 : (i+1)*11]
		number, err := strconv.ParseInt(n, 2, 64)
		if err != nil {
			return "", err
		}

		mnemonic += wordList[number] + " "
	}

	return strings.TrimRight(mnemonic, " "), nil
}

func SeedFromMnemonic(mnemonicWords []byte) [64]byte {
	q := pbkdf2.Key(mnemonicWords, nil, 2048, 64, sha512.New)

	var re [64]byte

	copy(re[:], q)

	return re
}
