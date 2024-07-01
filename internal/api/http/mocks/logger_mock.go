// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"net/http"
	"sync"
)

// ServerMock is a mock implementation of http.Server.
//
//	func TestSomethingThatUsesServer(t *testing.T) {
//
//		// make and configure a mocked http.Server
//		mockedServer := &ServerMock{
//			RegisterHandlersFunc: func(mux *http.ServeMux)  {
//				panic("mock out the RegisterHandlers method")
//			},
//		}
//
//		// use mockedServer in code that requires http.Server
//		// and then make assertions.
//
//	}
type ServerMock struct {
	// RegisterHandlersFunc mocks the RegisterHandlers method.
	RegisterHandlersFunc func(mux *http.ServeMux)

	// calls tracks calls to the methods.
	calls struct {
		// RegisterHandlers holds details about calls to the RegisterHandlers method.
		RegisterHandlers []struct {
			// Mux is the mux argument value.
			Mux *http.ServeMux
		}
	}
	lockRegisterHandlers sync.RWMutex
}

// RegisterHandlers calls RegisterHandlersFunc.
func (mock *ServerMock) RegisterHandlers(mux *http.ServeMux) {
	callInfo := struct {
		Mux *http.ServeMux
	}{
		Mux: mux,
	}
	mock.lockRegisterHandlers.Lock()
	mock.calls.RegisterHandlers = append(mock.calls.RegisterHandlers, callInfo)
	mock.lockRegisterHandlers.Unlock()
	if mock.RegisterHandlersFunc == nil {
		return
	}
	mock.RegisterHandlersFunc(mux)
}

// RegisterHandlersCalls gets all the calls that were made to RegisterHandlers.
// Check the length with:
//
//	len(mockedServer.RegisterHandlersCalls())
func (mock *ServerMock) RegisterHandlersCalls() []struct {
	Mux *http.ServeMux
} {
	var calls []struct {
		Mux *http.ServeMux
	}
	mock.lockRegisterHandlers.RLock()
	calls = mock.calls.RegisterHandlers
	mock.lockRegisterHandlers.RUnlock()
	return calls
}
