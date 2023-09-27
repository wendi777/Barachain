// Code generated by mockery v2.34.1. DO NOT EDIT.

package mock

import (
	client "github.com/cosmos/cosmos-sdk/client"
	common "github.com/ethereum/go-ethereum/common"

	core "pkg.berachain.dev/polaris/eth/core"

	coretxpool "github.com/ethereum/go-ethereum/core/txpool"

	log "cosmossdk.io/log"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

type Plugin_Expecter struct {
	mock *mock.Mock
}

func (_m *Plugin) EXPECT() *Plugin_Expecter {
	return &Plugin_Expecter{mock: &_m.Mock}
}

// GetHandler provides a mock function with given fields:
func (_m *Plugin) GetHandler() core.Handler {
	ret := _m.Called()

	var r0 core.Handler
	if rf, ok := ret.Get(0).(func() core.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Handler)
		}
	}

	return r0
}

// Plugin_GetHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHandler'
type Plugin_GetHandler_Call struct {
	*mock.Call
}

// GetHandler is a helper method to define mock.On call
func (_e *Plugin_Expecter) GetHandler() *Plugin_GetHandler_Call {
	return &Plugin_GetHandler_Call{Call: _e.mock.On("GetHandler")}
}

func (_c *Plugin_GetHandler_Call) Run(run func()) *Plugin_GetHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Plugin_GetHandler_Call) Return(_a0 core.Handler) *Plugin_GetHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_GetHandler_Call) RunAndReturn(run func() core.Handler) *Plugin_GetHandler_Call {
	_c.Call.Return(run)
	return _c
}

// Pending provides a mock function with given fields: enforceTips
func (_m *Plugin) Pending(enforceTips bool) map[common.Address][]*coretxpool.LazyTransaction {
	ret := _m.Called(enforceTips)

	var r0 map[common.Address][]*coretxpool.LazyTransaction
	if rf, ok := ret.Get(0).(func(bool) map[common.Address][]*coretxpool.LazyTransaction); ok {
		r0 = rf(enforceTips)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[common.Address][]*coretxpool.LazyTransaction)
		}
	}

	return r0
}

// Plugin_Pending_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Pending'
type Plugin_Pending_Call struct {
	*mock.Call
}

// Pending is a helper method to define mock.On call
//   - enforceTips bool
func (_e *Plugin_Expecter) Pending(enforceTips interface{}) *Plugin_Pending_Call {
	return &Plugin_Pending_Call{Call: _e.mock.On("Pending", enforceTips)}
}

func (_c *Plugin_Pending_Call) Run(run func(enforceTips bool)) *Plugin_Pending_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *Plugin_Pending_Call) Return(_a0 map[common.Address][]*coretxpool.LazyTransaction) *Plugin_Pending_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_Pending_Call) RunAndReturn(run func(bool) map[common.Address][]*coretxpool.LazyTransaction) *Plugin_Pending_Call {
	_c.Call.Return(run)
	return _c
}

// SerializeToBytes provides a mock function with given fields: signedTx
func (_m *Plugin) SerializeToBytes(signedTx *types.Transaction) ([]byte, error) {
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

// Plugin_SerializeToBytes_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SerializeToBytes'
type Plugin_SerializeToBytes_Call struct {
	*mock.Call
}

// SerializeToBytes is a helper method to define mock.On call
//   - signedTx *types.Transaction
func (_e *Plugin_Expecter) SerializeToBytes(signedTx interface{}) *Plugin_SerializeToBytes_Call {
	return &Plugin_SerializeToBytes_Call{Call: _e.mock.On("SerializeToBytes", signedTx)}
}

func (_c *Plugin_SerializeToBytes_Call) Run(run func(signedTx *types.Transaction)) *Plugin_SerializeToBytes_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Transaction))
	})
	return _c
}

func (_c *Plugin_SerializeToBytes_Call) Return(_a0 []byte, _a1 error) *Plugin_SerializeToBytes_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Plugin_SerializeToBytes_Call) RunAndReturn(run func(*types.Transaction) ([]byte, error)) *Plugin_SerializeToBytes_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: _a0, _a1, _a2
func (_m *Plugin) Start(_a0 log.Logger, _a1 *coretxpool.TxPool, _a2 client.Context) {
	_m.Called(_a0, _a1, _a2)
}

// Plugin_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Plugin_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - _a0 log.Logger
//   - _a1 *coretxpool.TxPool
//   - _a2 client.Context
func (_e *Plugin_Expecter) Start(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Plugin_Start_Call {
	return &Plugin_Start_Call{Call: _e.mock.On("Start", _a0, _a1, _a2)}
}

func (_c *Plugin_Start_Call) Run(run func(_a0 log.Logger, _a1 *coretxpool.TxPool, _a2 client.Context)) *Plugin_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(log.Logger), args[1].(*coretxpool.TxPool), args[2].(client.Context))
	})
	return _c
}

func (_c *Plugin_Start_Call) Return() *Plugin_Start_Call {
	_c.Call.Return()
	return _c
}

func (_c *Plugin_Start_Call) RunAndReturn(run func(log.Logger, *coretxpool.TxPool, client.Context)) *Plugin_Start_Call {
	_c.Call.Return(run)
	return _c
}

// TxPool provides a mock function with given fields:
func (_m *Plugin) TxPool() *coretxpool.TxPool {
	ret := _m.Called()

	var r0 *coretxpool.TxPool
	if rf, ok := ret.Get(0).(func() *coretxpool.TxPool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*coretxpool.TxPool)
		}
	}

	return r0
}

// Plugin_TxPool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TxPool'
type Plugin_TxPool_Call struct {
	*mock.Call
}

// TxPool is a helper method to define mock.On call
func (_e *Plugin_Expecter) TxPool() *Plugin_TxPool_Call {
	return &Plugin_TxPool_Call{Call: _e.mock.On("TxPool")}
}

func (_c *Plugin_TxPool_Call) Run(run func()) *Plugin_TxPool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Plugin_TxPool_Call) Return(_a0 *coretxpool.TxPool) *Plugin_TxPool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Plugin_TxPool_Call) RunAndReturn(run func() *coretxpool.TxPool) *Plugin_TxPool_Call {
	_c.Call.Return(run)
	return _c
}

// NewPlugin creates a new instance of Plugin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlugin(t interface {
	mock.TestingT
	Cleanup(func())
}) *Plugin {
	mock := &Plugin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}