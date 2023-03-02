// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"sync"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/types"

	"pkg.berachain.dev/stargazer/x/evm/plugins/state/events"
)

// Ensure, that PrecompileLogFactoryMock does implement events.PrecompileLogFactory.
// If this is not the case, regenerate this file with moq.
var _ events.PrecompileLogFactory = &PrecompileLogFactoryMock{}

// PrecompileLogFactoryMock is a mock implementation of events.PrecompileLogFactory.
//
//	func TestSomethingThatUsesPrecompileLogFactory(t *testing.T) {
//
//		// make and configure a mocked events.PrecompileLogFactory
//		mockedPrecompileLogFactory := &PrecompileLogFactoryMock{
//			BuildFunc: func(event *sdk.Event) (*types.Log, error) {
//				panic("mock out the Build method")
//			},
//		}
//
//		// use mockedPrecompileLogFactory in code that requires events.PrecompileLogFactory
//		// and then make assertions.
//
//	}
type PrecompileLogFactoryMock struct {
	// BuildFunc mocks the Build method.
	BuildFunc func(event *sdk.Event) (*types.Log, error)

	// calls tracks calls to the methods.
	calls struct {
		// Build holds details about calls to the Build method.
		Build []struct {
			// Event is the event argument value.
			Event *sdk.Event
		}
	}
	lockBuild sync.RWMutex
}

// Build calls BuildFunc.
func (mock *PrecompileLogFactoryMock) Build(event *sdk.Event) (*types.Log, error) {
	if mock.BuildFunc == nil {
		panic("PrecompileLogFactoryMock.BuildFunc: method is nil but PrecompileLogFactory.Build was just called")
	}
	callInfo := struct {
		Event *sdk.Event
	}{
		Event: event,
	}
	mock.lockBuild.Lock()
	mock.calls.Build = append(mock.calls.Build, callInfo)
	mock.lockBuild.Unlock()
	return mock.BuildFunc(event)
}

// BuildCalls gets all the calls that were made to Build.
// Check the length with:
//
//	len(mockedPrecompileLogFactory.BuildCalls())
func (mock *PrecompileLogFactoryMock) BuildCalls() []struct {
	Event *sdk.Event
} {
	var calls []struct {
		Event *sdk.Event
	}
	mock.lockBuild.RLock()
	calls = mock.calls.Build
	mock.lockBuild.RUnlock()
	return calls
}
