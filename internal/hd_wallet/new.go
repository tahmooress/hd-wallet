package hd_wallet

import (
	"bitcoin-wallet/internal/entities/structs"
	"crypto/elliptic"
	"math/big"
)

type Curve struct {
	params *elliptic.CurveParams
}

func New() *Curve {
	secp256k1 := new(elliptic.CurveParams)
	secp256k1.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	secp256k1.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	secp256k1.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
	secp256k1.Gx, _ = new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	secp256k1.Gy, _ = new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)
	secp256k1.BitSize = 256

	return &Curve{
		params: secp256k1,
	}
}

func (c *Curve) GenPublicKey(privateKey []byte) *structs.PubKey {
	x, y := c.params.ScalarBaseMult(privateKey)

	return &structs.PubKey{
		X: x,
		Y: y,
	}
}
