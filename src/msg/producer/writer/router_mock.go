// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3msg/producer/writer/router.go

package writer

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockackRouter is a mock of ackRouter interface
type MockackRouter struct {
	ctrl     *gomock.Controller
	recorder *MockackRouterMockRecorder
}

// MockackRouterMockRecorder is the mock recorder for MockackRouter
type MockackRouterMockRecorder struct {
	mock *MockackRouter
}

// NewMockackRouter creates a new mock instance
func NewMockackRouter(ctrl *gomock.Controller) *MockackRouter {
	mock := &MockackRouter{ctrl: ctrl}
	mock.recorder = &MockackRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockackRouter) EXPECT() *MockackRouterMockRecorder {
	return _m.recorder
}

// Ack mocks base method
func (_m *MockackRouter) Ack(ack metadata) error {
	ret := _m.ctrl.Call(_m, "Ack", ack)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack
func (_mr *MockackRouterMockRecorder) Ack(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Ack", reflect.TypeOf((*MockackRouter)(nil).Ack), arg0)
}
