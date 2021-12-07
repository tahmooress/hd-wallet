package structs

import (
	"math/big"
)

type PubKey struct {
	X *big.Int
	Y *big.Int
}

type ExtendedPrivateKey struct {
	PrivateKey []byte
	ChainCode  [32]byte
}

type ExtendedPubKey struct {
	*PubKey
	ChainCode [32]byte
}

type Address []byte

type MasterKey struct {
	RootPrivateKey []byte
}
