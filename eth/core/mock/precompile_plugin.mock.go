// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"pkg.berachain.dev/polaris/eth/core"
	libtypes "pkg.berachain.dev/polaris/lib/types"
	"sync"
)

// Ensure, that PrecompilePluginMock does implement core.PrecompilePlugin.
// If this is not the case, regenerate this file with moq.
var _ core.PrecompilePlugin = &PrecompilePluginMock{}

// PrecompilePluginMock is a mock implementation of core.PrecompilePlugin.
//
//	func TestSomethingThatUsesPrecompilePlugin(t *testing.T) {
//
//		// make and configure a mocked core.PrecompilePlugin
//		mockedPrecompilePlugin := &PrecompilePluginMock{
//			GetFunc: func(addr common.Address) vm.PrecompiledContract {
//				panic("mock out the Get method")
//			},
//			GetPrecompilesFunc: func(rules *params.Rules) []libtypes.Registrable[Address] {
//				panic("mock out the GetPrecompiles method")
//			},
//			HasFunc: func(addr common.Address) bool {
//				panic("mock out the Has method")
//			},
//			RegisterFunc: func(precompiledContract vm.PrecompiledContract) error {
//				panic("mock out the Register method")
//			},
//			RunFunc: func(sdb state.StateDBI, p vm.PrecompiledContract, input []byte, caller common.Address, value *big.Int, suppliedGas uint64, readonly bool) ([]byte, uint64, error) {
//				panic("mock out the Run method")
//			},
//		}
//
//		// use mockedPrecompilePlugin in code that requires core.PrecompilePlugin
//		// and then make assertions.
//
//	}
type PrecompilePluginMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(addr common.Address) vm.PrecompiledContract

	// GetPrecompilesFunc mocks the GetPrecompiles method.
	GetPrecompilesFunc func(rules *params.Rules) []libtypes.Registrable[Address]

	// HasFunc mocks the Has method.
	HasFunc func(addr common.Address) bool

	// RegisterFunc mocks the Register method.
	RegisterFunc func(precompiledContract vm.PrecompiledContract) error

	// RunFunc mocks the Run method.
	RunFunc func(sdb state.StateDBI, p vm.PrecompiledContract, input []byte, caller common.Address, value *big.Int, suppliedGas uint64, readonly bool) ([]byte, uint64, error)

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Addr is the addr argument value.
			Addr common.Address
		}
		// GetPrecompiles holds details about calls to the GetPrecompiles method.
		GetPrecompiles []struct {
			// Rules is the rules argument value.
			Rules *params.Rules
		}
		// Has holds details about calls to the Has method.
		Has []struct {
			// Addr is the addr argument value.
			Addr common.Address
		}
		// Register holds details about calls to the Register method.
		Register []struct {
			// PrecompiledContract is the precompiledContract argument value.
			PrecompiledContract vm.PrecompiledContract
		}
		// Run holds details about calls to the Run method.
		Run []struct {
			// Sdb is the sdb argument value.
			Sdb state.StateDBI
			// P is the p argument value.
			P vm.PrecompiledContract
			// Input is the input argument value.
			Input []byte
			// Caller is the caller argument value.
			Caller common.Address
			// Value is the value argument value.
			Value *big.Int
			// SuppliedGas is the suppliedGas argument value.
			SuppliedGas uint64
			// Readonly is the readonly argument value.
			Readonly bool
		}
	}
	lockGet            sync.RWMutex
	lockGetPrecompiles sync.RWMutex
	lockHas            sync.RWMutex
	lockRegister       sync.RWMutex
	lockRun            sync.RWMutex
}

// Get calls GetFunc.
func (mock *PrecompilePluginMock) Get(addr common.Address) vm.PrecompiledContract {
	if mock.GetFunc == nil {
		panic("PrecompilePluginMock.GetFunc: method is nil but PrecompilePlugin.Get was just called")
	}
	callInfo := struct {
		Addr common.Address
	}{
		Addr: addr,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(addr)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedPrecompilePlugin.GetCalls())
func (mock *PrecompilePluginMock) GetCalls() []struct {
	Addr common.Address
} {
	var calls []struct {
		Addr common.Address
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetPrecompiles calls GetPrecompilesFunc.
func (mock *PrecompilePluginMock) GetPrecompiles(rules *params.Rules) []libtypes.Registrable[Address] {
	if mock.GetPrecompilesFunc == nil {
		panic("PrecompilePluginMock.GetPrecompilesFunc: method is nil but PrecompilePlugin.GetPrecompiles was just called")
	}
	callInfo := struct {
		Rules *params.Rules
	}{
		Rules: rules,
	}
	mock.lockGetPrecompiles.Lock()
	mock.calls.GetPrecompiles = append(mock.calls.GetPrecompiles, callInfo)
	mock.lockGetPrecompiles.Unlock()
	return mock.GetPrecompilesFunc(rules)
}

// GetPrecompilesCalls gets all the calls that were made to GetPrecompiles.
// Check the length with:
//
//	len(mockedPrecompilePlugin.GetPrecompilesCalls())
func (mock *PrecompilePluginMock) GetPrecompilesCalls() []struct {
	Rules *params.Rules
} {
	var calls []struct {
		Rules *params.Rules
	}
	mock.lockGetPrecompiles.RLock()
	calls = mock.calls.GetPrecompiles
	mock.lockGetPrecompiles.RUnlock()
	return calls
}

// Has calls HasFunc.
func (mock *PrecompilePluginMock) Has(addr common.Address) bool {
	if mock.HasFunc == nil {
		panic("PrecompilePluginMock.HasFunc: method is nil but PrecompilePlugin.Has was just called")
	}
	callInfo := struct {
		Addr common.Address
	}{
		Addr: addr,
	}
	mock.lockHas.Lock()
	mock.calls.Has = append(mock.calls.Has, callInfo)
	mock.lockHas.Unlock()
	return mock.HasFunc(addr)
}

// HasCalls gets all the calls that were made to Has.
// Check the length with:
//
//	len(mockedPrecompilePlugin.HasCalls())
func (mock *PrecompilePluginMock) HasCalls() []struct {
	Addr common.Address
} {
	var calls []struct {
		Addr common.Address
	}
	mock.lockHas.RLock()
	calls = mock.calls.Has
	mock.lockHas.RUnlock()
	return calls
}

// Register calls RegisterFunc.
func (mock *PrecompilePluginMock) Register(precompiledContract vm.PrecompiledContract) error {
	if mock.RegisterFunc == nil {
		panic("PrecompilePluginMock.RegisterFunc: method is nil but PrecompilePlugin.Register was just called")
	}
	callInfo := struct {
		PrecompiledContract vm.PrecompiledContract
	}{
		PrecompiledContract: precompiledContract,
	}
	mock.lockRegister.Lock()
	mock.calls.Register = append(mock.calls.Register, callInfo)
	mock.lockRegister.Unlock()
	return mock.RegisterFunc(precompiledContract)
}

// RegisterCalls gets all the calls that were made to Register.
// Check the length with:
//
//	len(mockedPrecompilePlugin.RegisterCalls())
func (mock *PrecompilePluginMock) RegisterCalls() []struct {
	PrecompiledContract vm.PrecompiledContract
} {
	var calls []struct {
		PrecompiledContract vm.PrecompiledContract
	}
	mock.lockRegister.RLock()
	calls = mock.calls.Register
	mock.lockRegister.RUnlock()
	return calls
}

// Run calls RunFunc.
func (mock *PrecompilePluginMock) Run(sdb state.StateDBI, p vm.PrecompiledContract, input []byte, caller common.Address, value *big.Int, suppliedGas uint64, readonly bool) ([]byte, uint64, error) {
	if mock.RunFunc == nil {
		panic("PrecompilePluginMock.RunFunc: method is nil but PrecompilePlugin.Run was just called")
	}
	callInfo := struct {
		Sdb         state.StateDBI
		P           vm.PrecompiledContract
		Input       []byte
		Caller      common.Address
		Value       *big.Int
		SuppliedGas uint64
		Readonly    bool
	}{
		Sdb:         sdb,
		P:           p,
		Input:       input,
		Caller:      caller,
		Value:       value,
		SuppliedGas: suppliedGas,
		Readonly:    readonly,
	}
	mock.lockRun.Lock()
	mock.calls.Run = append(mock.calls.Run, callInfo)
	mock.lockRun.Unlock()
	return mock.RunFunc(sdb, p, input, caller, value, suppliedGas, readonly)
}

// RunCalls gets all the calls that were made to Run.
// Check the length with:
//
//	len(mockedPrecompilePlugin.RunCalls())
func (mock *PrecompilePluginMock) RunCalls() []struct {
	Sdb         state.StateDBI
	P           vm.PrecompiledContract
	Input       []byte
	Caller      common.Address
	Value       *big.Int
	SuppliedGas uint64
	Readonly    bool
} {
	var calls []struct {
		Sdb         state.StateDBI
		P           vm.PrecompiledContract
		Input       []byte
		Caller      common.Address
		Value       *big.Int
		SuppliedGas uint64
		Readonly    bool
	}
	mock.lockRun.RLock()
	calls = mock.calls.Run
	mock.lockRun.RUnlock()
	return calls
}
