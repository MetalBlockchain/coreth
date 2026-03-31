// Copyright (C) 2019-2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sync

import (
	"context"
	"errors"

	"github.com/MetalBlockchain/metalgo/codec"
	"github.com/MetalBlockchain/metalgo/ids"
	"github.com/MetalBlockchain/libevm/metrics"
	"github.com/MetalBlockchain/libevm/triedb"

	"github.com/MetalBlockchain/coreth/plugin/evm/message"
	"github.com/MetalBlockchain/coreth/sync/handlers"
	"github.com/MetalBlockchain/coreth/sync/handlers/stats"
)

var (
	_ handlers.LeafRequestHandler = (*uninitializedHandler)(nil)

	errUninitialized = errors.New("uninitialized handler")
)

type uninitializedHandler struct{}

func (*uninitializedHandler) OnLeafsRequest(_ context.Context, _ ids.NodeID, _ uint32, _ message.LeafsRequest) ([]byte, error) {
	return nil, errUninitialized
}

// atomicLeafHandler is a wrapper around handlers.LeafRequestHandler that allows for initialization after creation
type leafHandler struct {
	handlers.LeafRequestHandler
}

// NewAtomicLeafHandler returns a new uninitialized leafHandler that can be later initialized
func NewLeafHandler() *leafHandler {
	return &leafHandler{
		LeafRequestHandler: &uninitializedHandler{},
	}
}

// Initialize initializes the atomicLeafHandler with the provided atomicTrieDB, trieKeyLength, and networkCodec
func (a *leafHandler) Initialize(atomicTrieDB *triedb.Database, trieKeyLength int, networkCodec codec.Manager) {
	handlerStats := stats.GetOrRegisterHandlerStats(metrics.Enabled)
	a.LeafRequestHandler = handlers.NewLeafsRequestHandler(atomicTrieDB, trieKeyLength, nil, networkCodec, handlerStats)
}
