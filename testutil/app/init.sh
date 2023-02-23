#!/bin/bash
# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
# Use of this software is govered by the Business Source License included
# in the LICENSE file of this repository and at www.mariadb.com/bsl11.
#
# ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
# TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
# VERSIONS OF THE LICENSED WORK.
#
# THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
# LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
# LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
#
# TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
# AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
# EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
# TITLE.

KEYS[0]="dev0"
KEYS[1]="dev1"
KEYS[2]="dev2"
CHAINID="stargazer-2061"
MONIKER="localtestnet"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="test"
KEYALGO="secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the stargazerd instance
HOMEDIR="$HOME/.stargazerd"
# to trace evm
#TRACE="--trace"
TRACE=""

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# used to exit on first error (any non-zero exit code)
set -e

# Reinstall daemon
# mage install

# # User prompt if an existing local node configuration is found.
# if [ -d "$HOMEDIR" ]; then
# 	printf "\nAn existing folder at '%s' was found. You can choose to delete this folder and start a new local node with new keys from genesis. When declined, the existing local node is started. \n" "$HOMEDIR"
# 	echo "Overwrite the existing configuration and start a new local node? [y/n]"
# 	read -r overwrite
# else
overwrite="Y"
# fi


# Setup local node if overwrite is set to Yes, otherwise skip setup
if [[ $overwrite == "y" || $overwrite == "Y" ]]; then
	# Remove the previous folder
	rm -rf "$HOMEDIR"

    # Set moniker and chain-id (Moniker can be anything, chain-id must be an integer)
	stargazerd init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"

	# Set client config
	stargazerd config set client keyring-backend $KEYRING --home "$HOMEDIR"
	stargazerd config set client chain-id "$CHAINID" --home "$HOMEDIR"

	# If keys exist they should be deleted
	for KEY in "${KEYS[@]}"; do
		stargazerd keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
	done

	# Change parameter token denominations to abera
	jq '.app_state["staking"]["params"]["bond_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["crisis"]["constant_fee"]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["evm"]["params"]["evm_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["mint"]["params"]["mint_denom"]="abera"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Set gas limit in genesis
	jq '.consensus_params["block"]["max_gas"]="10000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "$CONFIG"
        sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "1s"/g' "$CONFIG"
        sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' "$CONFIG"
        sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "1s"/g' "$CONFIG"
        sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' "$CONFIG"
        sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "1s"/g' "$CONFIG"
        sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$CONFIG"
        sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "15s"/g' "$CONFIG"
    else
        sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "$CONFIG"
        sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "1s"/g' "$CONFIG"
        sed -i 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' "$CONFIG"
        sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "1s"/g' "$CONFIG"
        sed -i 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' "$CONFIG"
        sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "1s"/g' "$CONFIG"
        sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$CONFIG"
        sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "15s"/g' "$CONFIG"
    fi
	# Allocate genesis accounts (cosmos formatted addresses)
	for KEY in "${KEYS[@]}"; do
		stargazerd genesis add-genesis-account $KEY 100000000000000000000000000abera --keyring-backend $KEYRING --home "$HOMEDIR"
	done

	# Sign genesis transaction
	stargazerd genesis gentx ${KEYS[0]} 1000000000000000000000abera --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
	## In case you want to create multiple validators at genesis
	## 1. Back to `stargazerd keys add` step, init more keys
	## 2. Back to `stargazerd add-genesis-account` step, add balance for those
	## 3. Clone this ~/.stargazerd home directory into some others, let's say `~/.clonedstargazerd`
	## 4. Run `gentx` in each of those folders
	## 5. Copy the `gentx-*` folders under `~/.clonedstargazerd/config/gentx/` folders into the original `~/.stargazerd/config/gentx`

	# Collect genesis tx
	stargazerd genesis collect-gentxs --home "$HOMEDIR"

	# Run this to ensure everything worked and that the genesis file is setup correctly
	stargazerd genesis validate-genesis --home "$HOMEDIR"

	if [[ $1 == "pending" ]]; then
		echo "pending mode is on, please wait for the first block committed."
	fi
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
stargazerd start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --minimum-gas-prices=0.0001abera --home "$HOMEDIR"