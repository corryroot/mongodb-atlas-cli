// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: BackupConfigGetter,BackupConfigLister,BackupConfigUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	opsmngr "go.mongodb.org/ops-manager/opsmngr"
)

// MockBackupConfigGetter is a mock of BackupConfigGetter interface.
type MockBackupConfigGetter struct {
	ctrl     *gomock.Controller
	recorder *MockBackupConfigGetterMockRecorder
}

// MockBackupConfigGetterMockRecorder is the mock recorder for MockBackupConfigGetter.
type MockBackupConfigGetterMockRecorder struct {
	mock *MockBackupConfigGetter
}

// NewMockBackupConfigGetter creates a new mock instance.
func NewMockBackupConfigGetter(ctrl *gomock.Controller) *MockBackupConfigGetter {
	mock := &MockBackupConfigGetter{ctrl: ctrl}
	mock.recorder = &MockBackupConfigGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupConfigGetter) EXPECT() *MockBackupConfigGetterMockRecorder {
	return m.recorder
}

// GetBackupConfig mocks base method.
func (m *MockBackupConfigGetter) GetBackupConfig(arg0, arg1 string) (*opsmngr.BackupConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackupConfig", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.BackupConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackupConfig indicates an expected call of GetBackupConfig.
func (mr *MockBackupConfigGetterMockRecorder) GetBackupConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackupConfig", reflect.TypeOf((*MockBackupConfigGetter)(nil).GetBackupConfig), arg0, arg1)
}

// MockBackupConfigLister is a mock of BackupConfigLister interface.
type MockBackupConfigLister struct {
	ctrl     *gomock.Controller
	recorder *MockBackupConfigListerMockRecorder
}

// MockBackupConfigListerMockRecorder is the mock recorder for MockBackupConfigLister.
type MockBackupConfigListerMockRecorder struct {
	mock *MockBackupConfigLister
}

// NewMockBackupConfigLister creates a new mock instance.
func NewMockBackupConfigLister(ctrl *gomock.Controller) *MockBackupConfigLister {
	mock := &MockBackupConfigLister{ctrl: ctrl}
	mock.recorder = &MockBackupConfigListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupConfigLister) EXPECT() *MockBackupConfigListerMockRecorder {
	return m.recorder
}

// ListBackupConfigs mocks base method.
func (m *MockBackupConfigLister) ListBackupConfigs(arg0 string, arg1 *mongodbatlas.ListOptions) (*opsmngr.BackupConfigs, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBackupConfigs", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.BackupConfigs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackupConfigs indicates an expected call of ListBackupConfigs.
func (mr *MockBackupConfigListerMockRecorder) ListBackupConfigs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackupConfigs", reflect.TypeOf((*MockBackupConfigLister)(nil).ListBackupConfigs), arg0, arg1)
}

// MockBackupConfigUpdater is a mock of BackupConfigUpdater interface.
type MockBackupConfigUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockBackupConfigUpdaterMockRecorder
}

// MockBackupConfigUpdaterMockRecorder is the mock recorder for MockBackupConfigUpdater.
type MockBackupConfigUpdaterMockRecorder struct {
	mock *MockBackupConfigUpdater
}

// NewMockBackupConfigUpdater creates a new mock instance.
func NewMockBackupConfigUpdater(ctrl *gomock.Controller) *MockBackupConfigUpdater {
	mock := &MockBackupConfigUpdater{ctrl: ctrl}
	mock.recorder = &MockBackupConfigUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBackupConfigUpdater) EXPECT() *MockBackupConfigUpdaterMockRecorder {
	return m.recorder
}

// UpdateBackupConfig mocks base method.
func (m *MockBackupConfigUpdater) UpdateBackupConfig(arg0 *opsmngr.BackupConfig) (*opsmngr.BackupConfig, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBackupConfig", arg0)
	ret0, _ := ret[0].(*opsmngr.BackupConfig)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBackupConfig indicates an expected call of UpdateBackupConfig.
func (mr *MockBackupConfigUpdaterMockRecorder) UpdateBackupConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBackupConfig", reflect.TypeOf((*MockBackupConfigUpdater)(nil).UpdateBackupConfig), arg0)
}