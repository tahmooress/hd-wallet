package mnemonic

import (
	"crypto/sha256"
	"strconv"
	"strings"
)

func byte2bit(b []byte) string {
	bits := make([]string, 0)

	for i := range b {
		bb := strconv.FormatInt(int64(b[i]), 2)

		for len(bb) < 8 {
			bb = "0" + bb
		}

		bits = append(bits, bb)
	}

	return strings.Join(bits, "")
}

func addCheckSum(entropy string) (string, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(entropy)); err != nil {
		return "", err
	}

	digest := byte2bit(h.Sum(nil))

	checkSum := entropy + digest[:(len(entropy)/32)]

	if len(checkSum)%11 != 0 {
		return "", ErrCheckSumSize
	}

	return checkSum, nil
}
