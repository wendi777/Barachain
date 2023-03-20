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

package jsonrpc

import (
	"context"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"pkg.berachain.dev/polaris/cosmos/testing/network"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cosmos/testing/jsonrpc:integration")
}

var _ = Describe("Network", func() {
	var net *network.Network
	var client *ethclient.Client
	var wsClient *ethclient.Client
	_ = wsClient
	var rpcClient *gethrpc.Client
	var ctx context.Context
	var err error
	BeforeEach(func() {
		net = network.New(GinkgoT(), network.DefaultConfig())
		time.Sleep(1 * time.Second)
		_, err := net.WaitForHeightWithTimeout(1, defaultTimeout)
		Expect(err).ToNot(HaveOccurred())
	})

	AfterEach(func() {
		// TODO: FIX THE OFFCHAIN DB
		os.RemoveAll("data")
	})

	Context("checking rpc endopints", func() {

		BeforeEach(func() {
			// Dial an Ethereum RPC Endpoint
			rpcClient, err = gethrpc.DialContext(ctx, net.Validators[0].APIAddress+"/eth/rpc")
			Expect(err).ToNot(HaveOccurred())
			client = ethclient.NewClient(rpcClient)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should connect", func() {
			// Dial an Ethereum RPC Endpoint
			rpcClient, err := gethrpc.DialContext(ctx, net.Validators[0].APIAddress+"/eth/rpc")
			Expect(err).ToNot(HaveOccurred())
			client = ethclient.NewClient(rpcClient)
			Expect(err).ToNot(HaveOccurred())
		})

		It("eth_chainId, eth_gasPrice, eth_blockNumber, eth_getBalance", func() {
			chainID, err := client.ChainID(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(chainID.String()).To(Equal("69420"))
			gasPrice, err := client.SuggestGasPrice(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(gasPrice).ToNot(BeNil())
			blockNumber, err := client.BlockNumber(context.Background())
			Expect(err).ToNot(HaveOccurred())
			Expect(blockNumber).To(BeNumerically(">", 0))
			balance, err := client.BalanceAt(context.Background(), network.TestAddress, nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(balance).To(Equal(big.NewInt(1000000000000000000)))
		})

		It("should deploy, mint tokens, and check balance", func() {
			// Deploy the contract
			erc20Contract := DeployERC20(BuildTransactor(client), client)

			// Mint tokens
			tx, err := erc20Contract.Mint(BuildTransactor(client),
				network.TestAddress, big.NewInt(100000000))
			Expect(err).ToNot(HaveOccurred())
			ExpectMined(client, tx)
			ExpectSuccessReceipt(client, tx)

			// Check the erc20 balance
			erc20Balance, err := erc20Contract.BalanceOf(&bind.CallOpts{}, network.TestAddress)
			Expect(err).ToNot(HaveOccurred())
			Expect(erc20Balance).To(Equal(big.NewInt(100000000)))
		})
	})

	Context("checking websockets", func() {
		BeforeEach(func() {
			// Dial an Ethereum RPC Endpoint
			wsaddr := "ws:" + strings.TrimPrefix(net.Validators[0].APIAddress+"/eth/rpc/ws/", "http:")
			// rpcaddr := net.Validators[0].APIAddress + "/eth/rpc"
			ws, err := gethrpc.DialWebsocket(context.Background(), wsaddr, "*")
			Expect(err).ToNot(HaveOccurred())
			wsClient = ethclient.NewClient(ws)
			Expect(err).ToNot(HaveOccurred())
		})
		It("should connect", func() {
			// Dial an Ethereum websocket Endpoint
			wsaddr := "ws:" + strings.TrimPrefix(net.Validators[0].APIAddress+"/eth/rpc/ws/", "http:")
			// rpcaddr := net.Validators[0].APIAddress + "/eth/rpc"
			ws, err := gethrpc.DialWebsocket(context.Background(), wsaddr, "*")
			Expect(err).ToNot(HaveOccurred())
			wsClient = ethclient.NewClient(ws)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should subscribe to new heads", func() {
			// Subscribe to new heads
			sub, err := wsClient.SubscribeNewHead(context.Background(), make(chan *gethtypes.Header))
			Expect(err).ToNot(HaveOccurred())
			Expect(sub).ToNot(BeNil())
		})

		It("should subscribe to logs", func() {
			// Subscribe to logs
			sub, err := wsClient.SubscribeFilterLogs(context.Background(), geth.FilterQuery{}, make(chan gethtypes.Log))
			Expect(err).ToNot(HaveOccurred())
			Expect(sub).ToNot(BeNil())
		})
		// TODO: add scenario tests for websockets
	})
})
