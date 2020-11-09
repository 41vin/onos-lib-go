// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/controller/filter.go

// Package controller is a generated GoMock package.
package controller

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFilter is a mock of Filter interface
type MockFilter struct {
	ctrl     *gomock.Controller
	recorder *MockFilterMockRecorder
}

// MockFilterMockRecorder is the mock recorder for MockFilter
type MockFilterMockRecorder struct {
	mock *MockFilter
}

// NewMockFilter creates a new mock instance
func NewMockFilter(ctrl *gomock.Controller) *MockFilter {
	mock := &MockFilter{ctrl: ctrl}
	mock.recorder = &MockFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFilter) EXPECT() *MockFilterMockRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockFilter) Accept(id ID) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Accept indicates an expected call of Accept
func (mr *MockFilterMockRecorder) Accept(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockFilter)(nil).Accept), id)
}