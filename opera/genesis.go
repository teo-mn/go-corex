package opera

import (
	"math/big"

	"github.com/corex-mn/corex-base/hash"
	"github.com/corex-mn/corex-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"

	"github.com/corex-mn/go-corex/inter"
	"github.com/corex-mn/go-corex/opera/genesis"
	"github.com/corex-mn/go-corex/opera/genesis/gpos"
)

type Genesis struct {
	Accounts    genesis.Accounts
	Storage     genesis.Storage
	Delegations genesis.Delegations
	Blocks      genesis.Blocks
	RawEvmItems genesis.RawEvmItems
	Validators  gpos.Validators

	FirstEpoch    idx.Epoch
	PrevEpochTime inter.Timestamp
	Time          inter.Timestamp
	ExtraData     []byte

	TotalSupply *big.Int

	DriverOwner common.Address

	Rules Rules

	Hash func() hash.Hash
}
