package interfaces

import "bitcoin-wallet/internal/entities/structs"

type IWallet interface {
	Transfer(Amount uint64, address []byte) error
	Receive() structs.Address
	Balance() uint64
}
