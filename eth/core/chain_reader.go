// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// Use of this software is govered by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package core

import (
	"github.com/ethereum/go-ethereum/event"

	"pkg.berachain.dev/stargazer/eth/common"
	"pkg.berachain.dev/stargazer/eth/core/types"
)

// `CurrentHeader` returns the current header of the blockchain.
func (bc *blockchain) CurrentHeader() *types.StargazerHeader {
	return bc.processor.block.StargazerHeader
}

// `CurrentBlock` returns the current block of the blockchain.
func (bc *blockchain) CurrentBlock() *types.StargazerBlock {
	if bc.processor.block != nil {
		bc.blockCache.Add(bc.processor.block.Hash(), bc.processor.block)
	}
	return bc.processor.block
}

// `CurrentTransaction` returns the last finalized block of the blockchain.
func (bc *blockchain) FinalizedBlock() *types.StargazerBlock {
	if bc.processor.finalizedBlock != nil {
		bc.blockCache.Add(bc.processor.finalizedBlock.Hash(), bc.processor.finalizedBlock)
	}
	return bc.processor.finalizedBlock
}

// GetBlock retrieves a block from the database by hash and number,
// caching it if found.
func (bc *blockchain) GetStargazerBlockByNumber(number int64) *types.StargazerBlock {
	block := bc.Host().GetBlockPlugin().GetStargazerBlockByNumber(number)
	if block == nil {
		return nil
	}
	// Cache the found block for next time and return
	bc.blockCache.Add(block.Hash(), block)
	return block
}

// GetBlockByHash retrieves a block from the database by hash, caching it if found.
func (bc *blockchain) GetStargazerBlockByHash(hash common.Hash) *types.StargazerBlock {
	// Short circuit if the block's already in the cache, retrieve otherwise
	if block, ok := bc.blockCache.Get(hash); ok {
		return block
	}
	block := bc.Host().GetBlockPlugin().GetStargazerBlockByHash(hash)
	if block == nil {
		return nil
	}
	// Cache the found block for next time and return
	bc.blockCache.Add(block.Hash(), block)
	return block
}

// // SubscribeRemovedLogsEvent registers a subscription of RemovedLogsEvent.
// func (bc *blockchain) SubscribeRemovedLogsEvent(ch chan<- RemovedLogsEvent) event.Subscription {
// 	return bc.scope.Track(bc.rmLogsFeed.Subscribe(ch))
// }

// // SubscribeChainEvent registers a subscription of ChainEvent.
// func (bc *blockchain) SubscribeChainEvent(ch chan<- ChainEvent) event.Subscription {
// 	return bc.scope.Track(bc.chainFeed.Subscribe(ch))
// }

// SubscribeChainHeadEvent registers a subscription of ChainHeadEvent.
func (bc *blockchain) SubscribeChainHeadEvent(ch chan<- ChainHeadEvent) event.Subscription {
	return bc.scope.Track(bc.chainHeadFeed.Subscribe(ch))
}

// // SubscribeChainSideEvent registers a subscription of ChainSideEvent.
// func (bc *blockchain) SubscribeChainSideEvent(ch chan<- ChainSideEvent) event.Subscription {
// 	return bc.scope.Track(bc.chainSideFeed.Subscribe(ch))
// }

// // SubscribeLogsEvent registers a subscription of []*types.Log.
// func (bc *blockchain) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
// 	return bc.scope.Track(bc.logsFeed.Subscribe(ch))
// }

// // SubscribeBlockProcessingEvent registers a subscription of bool where true means
// // block processing has started while false means it has stopped.
// func (bc *blockchain) SubscribeBlockProcessingEvent(ch chan<- bool) event.Subscription {
// 	return bc.scope.Track(bc.blockProcFeed.Subscribe(ch))
// }
