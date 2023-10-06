// Code generated by mockery v2.35.2. DO NOT EDIT.

package mocks

import (
	cosmos_sdktypes "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// TxSerializer is an autogenerated mock type for the TxSerializer type
type TxSerializer struct {
	mock.Mock
}

type TxSerializer_Expecter struct {
	mock *mock.Mock
}

func (_m *TxSerializer) EXPECT() *TxSerializer_Expecter {
	return &TxSerializer_Expecter{mock: &_m.Mock}
}

// TxToSdkTxBytes provides a mock function with given fields: signedTx
func (_m *TxSerializer) TxToSdkTxBytes(signedTx *types.Transaction) ([]byte, error) {
	ret := _m.Called(signedTx)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.Transaction) ([]byte, error)); ok {
		return rf(signedTx)
	}
	if rf, ok := ret.Get(0).(func(*types.Transaction) []byte); ok {
		r0 = rf(signedTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.Transaction) error); ok {
		r1 = rf(signedTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxSerializer_TxToSdkTxBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxToSdkTxBytes'
type TxSerializer_TxToSdkTxBytes_Call struct {
	*mock.Call
}

// TxToSdkTxBytes is a helper method to define mock.On call
//   - signedTx *types.Transaction
func (_e *TxSerializer_Expecter) TxToSdkTxBytes(signedTx interface{}) *TxSerializer_TxToSdkTxBytes_Call {
	return &TxSerializer_TxToSdkTxBytes_Call{Call: _e.mock.On("TxToSdkTxBytes", signedTx)}
}

func (_c *TxSerializer_TxToSdkTxBytes_Call) Run(run func(signedTx *types.Transaction)) *TxSerializer_TxToSdkTxBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Transaction))
	})
	return _c
}

func (_c *TxSerializer_TxToSdkTxBytes_Call) Return(_a0 []byte, _a1 error) *TxSerializer_TxToSdkTxBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxSerializer_TxToSdkTxBytes_Call) RunAndReturn(run func(*types.Transaction) ([]byte, error)) *TxSerializer_TxToSdkTxBytes_Call {
	_c.Call.Return(run)
	return _c
}

// TxToSdkTx provides a mock function with given fields: signedTx
func (_m *TxSerializer) TxToSdkTx(signedTx *types.Transaction) (cosmos_sdktypes.Tx, error) {
	ret := _m.Called(signedTx)

	var r0 cosmos_sdktypes.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.Transaction) (cosmos_sdktypes.Tx, error)); ok {
		return rf(signedTx)
	}
	if rf, ok := ret.Get(0).(func(*types.Transaction) cosmos_sdktypes.Tx); ok {
		r0 = rf(signedTx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cosmos_sdktypes.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.Transaction) error); ok {
		r1 = rf(signedTx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxSerializer_TxToSdkTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxToSdkTx'
type TxSerializer_TxToSdkTx_Call struct {
	*mock.Call
}

// TxToSdkTx is a helper method to define mock.On call
//   - signedTx *types.Transaction
func (_e *TxSerializer_Expecter) TxToSdkTx(signedTx interface{}) *TxSerializer_TxToSdkTx_Call {
	return &TxSerializer_TxToSdkTx_Call{Call: _e.mock.On("TxToSdkTx", signedTx)}
}

func (_c *TxSerializer_TxToSdkTx_Call) Run(run func(signedTx *types.Transaction)) *TxSerializer_TxToSdkTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Transaction))
	})
	return _c
}

func (_c *TxSerializer_TxToSdkTx_Call) Return(_a0 cosmos_sdktypes.Tx, _a1 error) *TxSerializer_TxToSdkTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxSerializer_TxToSdkTx_Call) RunAndReturn(run func(*types.Transaction) (cosmos_sdktypes.Tx, error)) *TxSerializer_TxToSdkTx_Call {
	_c.Call.Return(run)
	return _c
}

// NewTxSerializer creates a new instance of TxSerializer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxSerializer(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxSerializer {
	mock := &TxSerializer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
