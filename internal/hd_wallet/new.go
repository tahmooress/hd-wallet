package hd_wallet

import (
	"bitcoin-wallet/internal/entities/interfaces"
	"bitcoin-wallet/internal/entities/structs"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"io"
	"math/big"
)

type singleWallet struct {
	curve *elliptic.CurveParams
}

func New() interfaces.IWallet {
	secp256k1 := new(elliptic.CurveParams)
	secp256k1.P, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)
	secp256k1.N, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	secp256k1.B, _ = new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000007", 16)
	secp256k1.Gx, _ = new(big.Int).SetString("79BE667EF9DCBBAC55A06295CE870B07029BFCDB2DCE28D959F2815B16F81798", 16)
	secp256k1.Gy, _ = new(big.Int).SetString("483ADA7726A3C4655DA4FBFC0E1108A8FD17B448A68554199C47D08FFB10D4B8", 16)
	secp256k1.BitSize = 256

	return &singleWallet{
		curve: secp256k1,
	}
}

func (s *singleWallet) GenerateKeyPair(reader io.Reader) (*structs.Wallet, error) {
	priv, err := ecdsa.GenerateKey(s.curve, reader)
	if err != nil {
		return nil, err
	}

	fmt.Println("is on curve: ", s.curve.IsOnCurve(priv.X, priv.Y))

	//return &structs.Wallet{
	//	PrivateKey: priv,
	//	PublicKey: &structs.PubKey{
	//		X: pubX,
	//		Y: pubY,
	//	},
	//}, nil

	return nil, nil
}
