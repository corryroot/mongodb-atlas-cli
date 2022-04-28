// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: MaintenanceWindowUpdater,MaintenanceWindowClearer,MaintenanceWindowDeferrer,MaintenanceWindowDescriber,OpsManagerMaintenanceWindowCreator,OpsManagerMaintenanceWindowLister,OpsManagerMaintenanceWindowDeleter,OpsManagerMaintenanceWindowDescriber,OpsManagerMaintenanceWindowUpdater)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	opsmngr "go.mongodb.org/ops-manager/opsmngr"
)

// MockMaintenanceWindowUpdater is a mock of MaintenanceWindowUpdater interface.
type MockMaintenanceWindowUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowUpdaterMockRecorder
}

// MockMaintenanceWindowUpdaterMockRecorder is the mock recorder for MockMaintenanceWindowUpdater.
type MockMaintenanceWindowUpdaterMockRecorder struct {
	mock *MockMaintenanceWindowUpdater
}

// NewMockMaintenanceWindowUpdater creates a new mock instance.
func NewMockMaintenanceWindowUpdater(ctrl *gomock.Controller) *MockMaintenanceWindowUpdater {
	mock := &MockMaintenanceWindowUpdater{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenanceWindowUpdater) EXPECT() *MockMaintenanceWindowUpdaterMockRecorder {
	return m.recorder
}

// UpdateMaintenanceWindow mocks base method.
func (m *MockMaintenanceWindowUpdater) UpdateMaintenanceWindow(arg0 string, arg1 *mongodbatlas.MaintenanceWindow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMaintenanceWindow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMaintenanceWindow indicates an expected call of UpdateMaintenanceWindow.
func (mr *MockMaintenanceWindowUpdaterMockRecorder) UpdateMaintenanceWindow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowUpdater)(nil).UpdateMaintenanceWindow), arg0, arg1)
}

// MockMaintenanceWindowClearer is a mock of MaintenanceWindowClearer interface.
type MockMaintenanceWindowClearer struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowClearerMockRecorder
}

// MockMaintenanceWindowClearerMockRecorder is the mock recorder for MockMaintenanceWindowClearer.
type MockMaintenanceWindowClearerMockRecorder struct {
	mock *MockMaintenanceWindowClearer
}

// NewMockMaintenanceWindowClearer creates a new mock instance.
func NewMockMaintenanceWindowClearer(ctrl *gomock.Controller) *MockMaintenanceWindowClearer {
	mock := &MockMaintenanceWindowClearer{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowClearerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenanceWindowClearer) EXPECT() *MockMaintenanceWindowClearerMockRecorder {
	return m.recorder
}

// ClearMaintenanceWindow mocks base method.
func (m *MockMaintenanceWindowClearer) ClearMaintenanceWindow(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearMaintenanceWindow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearMaintenanceWindow indicates an expected call of ClearMaintenanceWindow.
func (mr *MockMaintenanceWindowClearerMockRecorder) ClearMaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowClearer)(nil).ClearMaintenanceWindow), arg0)
}

// MockMaintenanceWindowDeferrer is a mock of MaintenanceWindowDeferrer interface.
type MockMaintenanceWindowDeferrer struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowDeferrerMockRecorder
}

// MockMaintenanceWindowDeferrerMockRecorder is the mock recorder for MockMaintenanceWindowDeferrer.
type MockMaintenanceWindowDeferrerMockRecorder struct {
	mock *MockMaintenanceWindowDeferrer
}

// NewMockMaintenanceWindowDeferrer creates a new mock instance.
func NewMockMaintenanceWindowDeferrer(ctrl *gomock.Controller) *MockMaintenanceWindowDeferrer {
	mock := &MockMaintenanceWindowDeferrer{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowDeferrerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenanceWindowDeferrer) EXPECT() *MockMaintenanceWindowDeferrerMockRecorder {
	return m.recorder
}

// DeferMaintenanceWindow mocks base method.
func (m *MockMaintenanceWindowDeferrer) DeferMaintenanceWindow(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeferMaintenanceWindow", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeferMaintenanceWindow indicates an expected call of DeferMaintenanceWindow.
func (mr *MockMaintenanceWindowDeferrerMockRecorder) DeferMaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeferMaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowDeferrer)(nil).DeferMaintenanceWindow), arg0)
}

// MockMaintenanceWindowDescriber is a mock of MaintenanceWindowDescriber interface.
type MockMaintenanceWindowDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockMaintenanceWindowDescriberMockRecorder
}

// MockMaintenanceWindowDescriberMockRecorder is the mock recorder for MockMaintenanceWindowDescriber.
type MockMaintenanceWindowDescriberMockRecorder struct {
	mock *MockMaintenanceWindowDescriber
}

// NewMockMaintenanceWindowDescriber creates a new mock instance.
func NewMockMaintenanceWindowDescriber(ctrl *gomock.Controller) *MockMaintenanceWindowDescriber {
	mock := &MockMaintenanceWindowDescriber{ctrl: ctrl}
	mock.recorder = &MockMaintenanceWindowDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMaintenanceWindowDescriber) EXPECT() *MockMaintenanceWindowDescriberMockRecorder {
	return m.recorder
}

// MaintenanceWindow mocks base method.
func (m *MockMaintenanceWindowDescriber) MaintenanceWindow(arg0 string) (*mongodbatlas.MaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaintenanceWindow", arg0)
	ret0, _ := ret[0].(*mongodbatlas.MaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaintenanceWindow indicates an expected call of MaintenanceWindow.
func (mr *MockMaintenanceWindowDescriberMockRecorder) MaintenanceWindow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaintenanceWindow", reflect.TypeOf((*MockMaintenanceWindowDescriber)(nil).MaintenanceWindow), arg0)
}

// MockOpsManagerMaintenanceWindowCreator is a mock of OpsManagerMaintenanceWindowCreator interface.
type MockOpsManagerMaintenanceWindowCreator struct {
	ctrl     *gomock.Controller
	recorder *MockOpsManagerMaintenanceWindowCreatorMockRecorder
}

// MockOpsManagerMaintenanceWindowCreatorMockRecorder is the mock recorder for MockOpsManagerMaintenanceWindowCreator.
type MockOpsManagerMaintenanceWindowCreatorMockRecorder struct {
	mock *MockOpsManagerMaintenanceWindowCreator
}

// NewMockOpsManagerMaintenanceWindowCreator creates a new mock instance.
func NewMockOpsManagerMaintenanceWindowCreator(ctrl *gomock.Controller) *MockOpsManagerMaintenanceWindowCreator {
	mock := &MockOpsManagerMaintenanceWindowCreator{ctrl: ctrl}
	mock.recorder = &MockOpsManagerMaintenanceWindowCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpsManagerMaintenanceWindowCreator) EXPECT() *MockOpsManagerMaintenanceWindowCreatorMockRecorder {
	return m.recorder
}

// CreateOpsManagerMaintenanceWindow mocks base method.
func (m *MockOpsManagerMaintenanceWindowCreator) CreateOpsManagerMaintenanceWindow(arg0 string, arg1 *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOpsManagerMaintenanceWindow", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.MaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOpsManagerMaintenanceWindow indicates an expected call of CreateOpsManagerMaintenanceWindow.
func (mr *MockOpsManagerMaintenanceWindowCreatorMockRecorder) CreateOpsManagerMaintenanceWindow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOpsManagerMaintenanceWindow", reflect.TypeOf((*MockOpsManagerMaintenanceWindowCreator)(nil).CreateOpsManagerMaintenanceWindow), arg0, arg1)
}

// MockOpsManagerMaintenanceWindowLister is a mock of OpsManagerMaintenanceWindowLister interface.
type MockOpsManagerMaintenanceWindowLister struct {
	ctrl     *gomock.Controller
	recorder *MockOpsManagerMaintenanceWindowListerMockRecorder
}

// MockOpsManagerMaintenanceWindowListerMockRecorder is the mock recorder for MockOpsManagerMaintenanceWindowLister.
type MockOpsManagerMaintenanceWindowListerMockRecorder struct {
	mock *MockOpsManagerMaintenanceWindowLister
}

// NewMockOpsManagerMaintenanceWindowLister creates a new mock instance.
func NewMockOpsManagerMaintenanceWindowLister(ctrl *gomock.Controller) *MockOpsManagerMaintenanceWindowLister {
	mock := &MockOpsManagerMaintenanceWindowLister{ctrl: ctrl}
	mock.recorder = &MockOpsManagerMaintenanceWindowListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpsManagerMaintenanceWindowLister) EXPECT() *MockOpsManagerMaintenanceWindowListerMockRecorder {
	return m.recorder
}

// OpsManagerMaintenanceWindows mocks base method.
func (m *MockOpsManagerMaintenanceWindowLister) OpsManagerMaintenanceWindows(arg0 string) (*opsmngr.MaintenanceWindows, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpsManagerMaintenanceWindows", arg0)
	ret0, _ := ret[0].(*opsmngr.MaintenanceWindows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpsManagerMaintenanceWindows indicates an expected call of OpsManagerMaintenanceWindows.
func (mr *MockOpsManagerMaintenanceWindowListerMockRecorder) OpsManagerMaintenanceWindows(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpsManagerMaintenanceWindows", reflect.TypeOf((*MockOpsManagerMaintenanceWindowLister)(nil).OpsManagerMaintenanceWindows), arg0)
}

// MockOpsManagerMaintenanceWindowDeleter is a mock of OpsManagerMaintenanceWindowDeleter interface.
type MockOpsManagerMaintenanceWindowDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockOpsManagerMaintenanceWindowDeleterMockRecorder
}

// MockOpsManagerMaintenanceWindowDeleterMockRecorder is the mock recorder for MockOpsManagerMaintenanceWindowDeleter.
type MockOpsManagerMaintenanceWindowDeleterMockRecorder struct {
	mock *MockOpsManagerMaintenanceWindowDeleter
}

// NewMockOpsManagerMaintenanceWindowDeleter creates a new mock instance.
func NewMockOpsManagerMaintenanceWindowDeleter(ctrl *gomock.Controller) *MockOpsManagerMaintenanceWindowDeleter {
	mock := &MockOpsManagerMaintenanceWindowDeleter{ctrl: ctrl}
	mock.recorder = &MockOpsManagerMaintenanceWindowDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpsManagerMaintenanceWindowDeleter) EXPECT() *MockOpsManagerMaintenanceWindowDeleterMockRecorder {
	return m.recorder
}

// DeleteOpsManagerMaintenanceWindow mocks base method.
func (m *MockOpsManagerMaintenanceWindowDeleter) DeleteOpsManagerMaintenanceWindow(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOpsManagerMaintenanceWindow", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOpsManagerMaintenanceWindow indicates an expected call of DeleteOpsManagerMaintenanceWindow.
func (mr *MockOpsManagerMaintenanceWindowDeleterMockRecorder) DeleteOpsManagerMaintenanceWindow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOpsManagerMaintenanceWindow", reflect.TypeOf((*MockOpsManagerMaintenanceWindowDeleter)(nil).DeleteOpsManagerMaintenanceWindow), arg0, arg1)
}

// MockOpsManagerMaintenanceWindowDescriber is a mock of OpsManagerMaintenanceWindowDescriber interface.
type MockOpsManagerMaintenanceWindowDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockOpsManagerMaintenanceWindowDescriberMockRecorder
}

// MockOpsManagerMaintenanceWindowDescriberMockRecorder is the mock recorder for MockOpsManagerMaintenanceWindowDescriber.
type MockOpsManagerMaintenanceWindowDescriberMockRecorder struct {
	mock *MockOpsManagerMaintenanceWindowDescriber
}

// NewMockOpsManagerMaintenanceWindowDescriber creates a new mock instance.
func NewMockOpsManagerMaintenanceWindowDescriber(ctrl *gomock.Controller) *MockOpsManagerMaintenanceWindowDescriber {
	mock := &MockOpsManagerMaintenanceWindowDescriber{ctrl: ctrl}
	mock.recorder = &MockOpsManagerMaintenanceWindowDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpsManagerMaintenanceWindowDescriber) EXPECT() *MockOpsManagerMaintenanceWindowDescriberMockRecorder {
	return m.recorder
}

// OpsManagerMaintenanceWindow mocks base method.
func (m *MockOpsManagerMaintenanceWindowDescriber) OpsManagerMaintenanceWindow(arg0, arg1 string) (*opsmngr.MaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpsManagerMaintenanceWindow", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.MaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpsManagerMaintenanceWindow indicates an expected call of OpsManagerMaintenanceWindow.
func (mr *MockOpsManagerMaintenanceWindowDescriberMockRecorder) OpsManagerMaintenanceWindow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpsManagerMaintenanceWindow", reflect.TypeOf((*MockOpsManagerMaintenanceWindowDescriber)(nil).OpsManagerMaintenanceWindow), arg0, arg1)
}

// MockOpsManagerMaintenanceWindowUpdater is a mock of OpsManagerMaintenanceWindowUpdater interface.
type MockOpsManagerMaintenanceWindowUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockOpsManagerMaintenanceWindowUpdaterMockRecorder
}

// MockOpsManagerMaintenanceWindowUpdaterMockRecorder is the mock recorder for MockOpsManagerMaintenanceWindowUpdater.
type MockOpsManagerMaintenanceWindowUpdaterMockRecorder struct {
	mock *MockOpsManagerMaintenanceWindowUpdater
}

// NewMockOpsManagerMaintenanceWindowUpdater creates a new mock instance.
func NewMockOpsManagerMaintenanceWindowUpdater(ctrl *gomock.Controller) *MockOpsManagerMaintenanceWindowUpdater {
	mock := &MockOpsManagerMaintenanceWindowUpdater{ctrl: ctrl}
	mock.recorder = &MockOpsManagerMaintenanceWindowUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpsManagerMaintenanceWindowUpdater) EXPECT() *MockOpsManagerMaintenanceWindowUpdaterMockRecorder {
	return m.recorder
}

// UpdateOpsManagerMaintenanceWindow mocks base method.
func (m *MockOpsManagerMaintenanceWindowUpdater) UpdateOpsManagerMaintenanceWindow(arg0 string, arg1 *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOpsManagerMaintenanceWindow", arg0, arg1)
	ret0, _ := ret[0].(*opsmngr.MaintenanceWindow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOpsManagerMaintenanceWindow indicates an expected call of UpdateOpsManagerMaintenanceWindow.
func (mr *MockOpsManagerMaintenanceWindowUpdaterMockRecorder) UpdateOpsManagerMaintenanceWindow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOpsManagerMaintenanceWindow", reflect.TypeOf((*MockOpsManagerMaintenanceWindowUpdater)(nil).UpdateOpsManagerMaintenanceWindow), arg0, arg1)
}