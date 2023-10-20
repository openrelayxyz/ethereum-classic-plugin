// Copyright 2019 The multi-geth Authors
// This file is part of the multi-geth library.
//
// The multi-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The multi-geth library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the multi-geth library. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"math/big"
	"errors"

	"github.com/openrelayxyz/plugeth-utils/core"
)

// Lengths of hashes and addresses in bytes.
const (
	// HashLength is the expected length of the hash
	HashLength = 32
	// AddressLength is the expected length of the address
	AddressLength = 20
)

// Uint64BigMapEncodesHex is a map that encodes and decodes w/ JSON hex format.
type Uint64BigMapEncodesHex map[uint64]*big.Int

type ConsensusEngineT int

const (
	ConsensusEngineT_Unknown = iota
	ConsensusEngineT_Ethash
	ConsensusEngineT_Clique
	ConsensusEngineT_Lyra2
)

func (c ConsensusEngineT) String() string {
	switch c {
	case ConsensusEngineT_Ethash:
		return "ethash"
	case ConsensusEngineT_Clique:
		return "clique"
	case ConsensusEngineT_Lyra2:
		return "lyra2"
	default:
		return "unknown"
	}
}

func (c ConsensusEngineT) IsEthash() bool {
	return c == ConsensusEngineT_Ethash
}

func (c ConsensusEngineT) IsClique() bool {
	return c == ConsensusEngineT_Clique
}

func (c ConsensusEngineT) IsLyra2() bool {
	return c == ConsensusEngineT_Lyra2
}

func (c ConsensusEngineT) IsUnknown() bool {
	return c == ConsensusEngineT_Unknown
}

type BlockSealingT int

const (
	BlockSealing_Unknown = iota
	BlockSealing_Ethereum
)

func (b BlockSealingT) String() string {
	switch b {
	case BlockSealing_Ethereum:
		return "ethereum"
	default:
		return "unknown"
	}
}

var big0 = big.NewInt(0)

var big1 = new(big.Int).SetInt64(1)

var big2 = new(big.Int).SetInt64(2)

var bigMinus99 = big.NewInt(-99)

var DisinflationRateQuotient = big.NewInt(4)

var DisinflationRateDivisor  = big.NewInt(5)

// DAOForkBlockExtra is the block header extra-data field to set for the DAO fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
// pick the side they want  ("dao-hard-fork").
var DAOForkBlockExtra = FromHex("0x64616f2d686172642d666f726b")

// DAOForkExtraRange is the number of consecutive blocks from the DAO fork point
// to override the extra-data in to prevent no-fork attacks.
var DAOForkExtraRange = big.NewInt(10)

// DAORefundContract is the address of the refund contract to send DAO balances to.
var DAORefundContract = core.HexToAddress("0xbf4ed7b27f1d666546e30d74d50d173d20bca754")

var ErrBadProDAOExtra = errors.New("bad DAO pro-fork extra-data")

var ExpDiffPeriod *big.Int = big.NewInt(100000)

const (
	datasetInitBytes    = 1 << 30 // Bytes in dataset at genesis
	datasetGrowthBytes  = 1 << 23 // Dataset growth per epoch
	cacheInitBytes      = 1 << 24 // Bytes in cache at genesis
	cacheGrowthBytes    = 1 << 17 // Cache growth per epoch
	epochLengthDefault  = 30000   // Default epoch length (blocks per epoch)
	epochLengthECIP1099 = 60000   // Blocks per epoch if ECIP-1099 is activated
	mixBytes            = 128     // Width of mix
	hashBytes           = 64      // Hash length in bytes
	hashWords           = 16      // Number of 32 bit ints in a hash
	datasetParents      = 256     // Number of parents of each dataset element
	cacheRounds         = 3       // Number of rounds in cache production
	loopAccesses        = 64      // Number of accesses in hashimoto loop
	maxEpoch            = 2048    // Max Epoch for included tables
)