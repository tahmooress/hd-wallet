package interfaces

import (
	"bitcoin-wallet/internal/entities/structs"
	"io"
)

type IWallet interface {
	GenerateKeyPair(reader io.Reader) (*structs.Wallet, error)
	//Balance() (float64, error)
	//Transfer() error
}
