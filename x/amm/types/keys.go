package types

import (
	"github.com/cosmos-builders/chaos/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "amm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the message route for this module
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_amm"
)

var (
	LastPairIDKey      = []byte{0x01}
	PairKeyPrefix      = []byte{0x02}
	PairIndexKeyPrefix = []byte{0x03}
)

func GetPairKey(pairID uint64) []byte {
	return append(PairKeyPrefix, sdk.Uint64ToBigEndian(pairID)...)
}

func GetPairIndexKey(denom0, denom1 string) []byte {
	return append(append(PairIndexKeyPrefix, util.LengthPrefix([]byte(denom0))...), denom1...)
}
