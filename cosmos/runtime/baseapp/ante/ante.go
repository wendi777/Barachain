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

package ante

import (
	builderdecorator "github.com/skip-mev/pob/x/builder/ante"
	builder "github.com/skip-mev/pob/x/builder/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"

	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool/mempool"
	"pkg.berachain.dev/polaris/lib/errors"
)

// NewAnteHandler returns an AnteHandler that checks and increments sequence
// numbers, checks signatures & account numbers, and deducts fees from the first
// signer.
func NewAnteHandler(options ante.HandlerOptions, builderKeeper builder.Keeper,
	txDecoder sdk.TxDecoder, txEncoder sdk.TxEncoder, mempool *mempool.EthTxPool,
) (sdk.AnteHandler, error) {
	// Validate the supplied options
	if err := ValidateOptions(options); err != nil {
		return nil, err
	}

	// Build the AnteDecorator pipeline
	anteDecorators := append(Decorators(options),
		builderdecorator.NewBuilderDecorator(builderKeeper, txDecoder, txEncoder, mempool),
	)

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

// NewProposalAnteHandler returns the set of ante handlers that are run when building
// and verifying block proposals. This is similar to the NewAnteHandler, however, it
// enforces that nonces are correctly sequenced and updated during Prepare and Process
// Proposal. This allows transactions to be broadcasted out of order but sequenced in
// the expected order when blocks are built/verified.
func NewProposalAnteHandler(options ante.HandlerOptions, builderKeeper builder.Keeper,
	txDecoder sdk.TxDecoder, txEncoder sdk.TxEncoder,
	mempool *mempool.EthTxPool) (sdk.AnteHandler, error) {
	// Validate the supplied options
	if err := ValidateOptions(options); err != nil {
		return nil, err
	}

	// Build the AnteDecorator pipeline
	anteDecorators := append(Decorators(options),
		// Validate the nonces of all accounts that are signers in the transaction.
		NewNonceDecorator(options.AccountKeeper),
		// Update the nonces of all accounts that are signers in the transaction.
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		builderdecorator.NewBuilderDecorator(builderKeeper, txDecoder, txEncoder, mempool),
	)

	return sdk.ChainAnteDecorators(anteDecorators...), nil
}

// ValidateOptions performs a basic validation of the provided options.
func ValidateOptions(options ante.HandlerOptions) error {
	if options.AccountKeeper == nil {
		return errors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}

	if options.BankKeeper == nil {
		return errors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}

	if options.FeegrantKeeper == nil {
		return errors.Wrap(sdkerrors.ErrLogic, "feegrant keeper is required for ante builder")
	}

	if options.TxFeeChecker == nil {
		return errors.Wrap(sdkerrors.ErrLogic, "tx fee checker is required for ante builder")
	}

	if options.SignModeHandler == nil {
		return errors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	return nil
}

// Decorators returns the default AnteDecorators that Polaris Chains should use.
func Decorators(options ante.HandlerOptions) []sdk.AnteDecorator {
	return []sdk.AnteDecorator{
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		// EthTransactions can skip consuming transaction gas as it will be done
		// in the StateTransition.
		EthSkipDecorator[ante.ConsumeTxSizeGasDecorator]{
			ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		},
		// EthTransaction can skip deduct fee transactions as they are done in the
		// StateTransition. // TODO: check to make sure this doesn't cause spam.
		EthSkipDecorator[ante.DeductFeeDecorator]{
			ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper,
				options.FeegrantKeeper, options.TxFeeChecker),
		},
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		// In order to match ethereum gas consumption, we do not consume any gas when
		// verifying the signature.
		EthSkipDecorator[ante.SigGasConsumeDecorator]{
			ante.NewSigGasConsumeDecorator(options.AccountKeeper, options.SigGasConsumer),
		},
		// EthTransaction can skip Signature Verification as we do this in the mempool.
		// TODO: // check with Marko to make sure this is okay (ties into the one below)
		EthSkipDecorator[ante.SigVerificationDecorator]{
			ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		},
	}
}
