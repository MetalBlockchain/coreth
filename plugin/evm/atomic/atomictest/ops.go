// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomictest

import (
	"github.com/MetalBlockchain/metalgo/ids"

	"github.com/MetalBlockchain/coreth/plugin/evm/atomic"

	avalancheatomic "github.com/MetalBlockchain/metalgo/chains/atomic"
)

func ConvertToAtomicOps(tx *atomic.Tx) (map[ids.ID]*avalancheatomic.Requests, error) {
	id, reqs, err := tx.AtomicOps()
	if err != nil {
		return nil, err
	}
	return map[ids.ID]*avalancheatomic.Requests{id: reqs}, nil
}
