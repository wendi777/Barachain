// Code generated by rlpgen. DO NOT EDIT.

//go:build !norlpgen
// +build !norlpgen

package types

import (
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func (obj *TxLookupEntry) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	if err := obj.Tx.EncodeRLP(w); err != nil {
		return err
	}
	w.WriteUint64(obj.TxIndex)
	w.WriteUint64(obj.BlockNum)
	w.WriteBytes(obj.BlockHash[:])
	w.ListEnd(_tmp0)
	return w.Flush()
}

func (obj *TxLookupEntry) DecodeRLP(dec *rlp.Stream) error {
	var _tmp0 TxLookupEntry
	{
		if _, err := dec.List(); err != nil {
			return err
		}
		// Tx:
		_tmp1 := new(types.Transaction)
		if err := _tmp1.DecodeRLP(dec); err != nil {
			return err
		}
		_tmp0.Tx = _tmp1
		// TxIndex:
		_tmp2, err := dec.Uint64()
		if err != nil {
			return err
		}
		_tmp0.TxIndex = _tmp2
		// BlockNum:
		_tmp3, err := dec.Uint64()
		if err != nil {
			return err
		}
		_tmp0.BlockNum = _tmp3
		// BlockHash:
		var _tmp4 common.Hash
		if err := dec.ReadBytes(_tmp4[:]); err != nil {
			return err
		}
		_tmp0.BlockHash = _tmp4
		if err := dec.ListEnd(); err != nil {
			return err
		}
	}
	*obj = _tmp0
	return nil
}
