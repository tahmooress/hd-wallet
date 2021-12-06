package structs

import (
	"crypto/ecdsa"
	"math/big"
)

type PubKey struct {
	X *big.Int
	Y *big.Int
}

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *PubKey
}
