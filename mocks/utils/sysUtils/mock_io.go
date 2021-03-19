// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/utils/sysUtils (interfaces: IoI,IoUtilI)

// Package mock_sysUtils is a generated GoMock package.
package mock_sysUtils

import (
	io "io"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIoI is a mock of IoI interface.
type MockIoI struct {
	ctrl     *gomock.Controller
	recorder *MockIoIMockRecorder
}

// MockIoIMockRecorder is the mock recorder for MockIoI.
type MockIoIMockRecorder struct {
	mock *MockIoI
}

// NewMockIoI creates a new mock instance.
func NewMockIoI(ctrl *gomock.Controller) *MockIoI {
	mock := &MockIoI{ctrl: ctrl}
	mock.recorder = &MockIoIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIoI) EXPECT() *MockIoIMockRecorder {
	return m.recorder
}

// Copy mocks base method.
func (m *MockIoI) Copy(arg0 io.Writer, arg1 io.Reader) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Copy indicates an expected call of Copy.
func (mr *MockIoIMockRecorder) Copy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockIoI)(nil).Copy), arg0, arg1)
}

// MockIoUtilI is a mock of IoUtilI interface.
type MockIoUtilI struct {
	ctrl     *gomock.Controller
	recorder *MockIoUtilIMockRecorder
}

// MockIoUtilIMockRecorder is the mock recorder for MockIoUtilI.
type MockIoUtilIMockRecorder struct {
	mock *MockIoUtilI
}

// NewMockIoUtilI creates a new mock instance.
func NewMockIoUtilI(ctrl *gomock.Controller) *MockIoUtilI {
	mock := &MockIoUtilI{ctrl: ctrl}
	mock.recorder = &MockIoUtilIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIoUtilI) EXPECT() *MockIoUtilIMockRecorder {
	return m.recorder
}

// NopCloser mocks base method.
func (m *MockIoUtilI) NopCloser(arg0 io.Reader) io.ReadCloser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NopCloser", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	return ret0
}

// NopCloser indicates an expected call of NopCloser.
func (mr *MockIoUtilIMockRecorder) NopCloser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NopCloser", reflect.TypeOf((*MockIoUtilI)(nil).NopCloser), arg0)
}

// ReadAll mocks base method.
func (m *MockIoUtilI) ReadAll(arg0 io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAll", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAll indicates an expected call of ReadAll.
func (mr *MockIoUtilIMockRecorder) ReadAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAll", reflect.TypeOf((*MockIoUtilI)(nil).ReadAll), arg0)
}

// ReadFile mocks base method.
func (m *MockIoUtilI) ReadFile(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFile", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFile indicates an expected call of ReadFile.
func (mr *MockIoUtilIMockRecorder) ReadFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFile", reflect.TypeOf((*MockIoUtilI)(nil).ReadFile), arg0)
}

// WriteFile mocks base method.
func (m *MockIoUtilI) WriteFile(arg0 string, arg1 []byte, arg2 os.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteFile indicates an expected call of WriteFile.
func (mr *MockIoUtilIMockRecorder) WriteFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteFile", reflect.TypeOf((*MockIoUtilI)(nil).WriteFile), arg0, arg1, arg2)
}
