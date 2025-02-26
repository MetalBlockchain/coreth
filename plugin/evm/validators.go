// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"

	"github.com/MetalBlockchain/metalgo/ids"
	"github.com/MetalBlockchain/metalgo/utils/set"
)

type validatorSet struct {
	set set.Set[ids.NodeID]
}

func (v *validatorSet) Has(ctx context.Context, nodeID ids.NodeID) bool {
	return v.set.Contains(nodeID)
}
