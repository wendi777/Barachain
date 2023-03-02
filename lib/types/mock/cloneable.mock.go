// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"sync"

	"pkg.berachain.dev/stargazer/lib/types"
)

// Ensure, that CloneableMock does implement types.Cloneable.
// If this is not the case, regenerate this file with moq.
var _ types.Cloneable[any] = &CloneableMock[any]{}

// CloneableMock is a mock implementation of types.Cloneable.
//
//	func TestSomethingThatUsesCloneable(t *testing.T) {
//
//		// make and configure a mocked types.Cloneable
//		mockedCloneable := &CloneableMock{
//			CloneFunc: func() T {
//				panic("mock out the Clone method")
//			},
//		}
//
//		// use mockedCloneable in code that requires types.Cloneable
//		// and then make assertions.
//
//	}
type CloneableMock[T any] struct {
	// CloneFunc mocks the Clone method.
	CloneFunc func() T

	// calls tracks calls to the methods.
	calls struct {
		// Clone holds details about calls to the Clone method.
		Clone []struct {
		}
	}
	lockClone sync.RWMutex
}

// Clone calls CloneFunc.
func (mock *CloneableMock[T]) Clone() T {
	if mock.CloneFunc == nil {
		panic("CloneableMock.CloneFunc: method is nil but Cloneable.Clone was just called")
	}
	callInfo := struct {
	}{}
	mock.lockClone.Lock()
	mock.calls.Clone = append(mock.calls.Clone, callInfo)
	mock.lockClone.Unlock()
	return mock.CloneFunc()
}

// CloneCalls gets all the calls that were made to Clone.
// Check the length with:
//
//	len(mockedCloneable.CloneCalls())
func (mock *CloneableMock[T]) CloneCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockClone.RLock()
	calls = mock.calls.Clone
	mock.lockClone.RUnlock()
	return calls
}
