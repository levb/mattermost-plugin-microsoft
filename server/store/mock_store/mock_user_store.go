// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-msoffice/server/store (interfaces: UserStore)

// Package mock_store is a generated GoMock package.
package mock_store

import (
	gomock "github.com/golang/mock/gomock"
	store "github.com/mattermost/mattermost-plugin-msoffice/server/store"
	reflect "reflect"
)

// MockUserStore is a mock of UserStore interface
type MockUserStore struct {
	ctrl     *gomock.Controller
	recorder *MockUserStoreMockRecorder
}

// MockUserStoreMockRecorder is the mock recorder for MockUserStore
type MockUserStoreMockRecorder struct {
	mock *MockUserStore
}

// NewMockUserStore creates a new mock instance
func NewMockUserStore(ctrl *gomock.Controller) *MockUserStore {
	mock := &MockUserStore{ctrl: ctrl}
	mock.recorder = &MockUserStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserStore) EXPECT() *MockUserStoreMockRecorder {
	return m.recorder
}

// DeleteUser mocks base method
func (m *MockUserStore) DeleteUser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUserStoreMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserStore)(nil).DeleteUser), arg0)
}

// LoadMattermostUserId mocks base method
func (m *MockUserStore) LoadMattermostUserId(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadMattermostUserId", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadMattermostUserId indicates an expected call of LoadMattermostUserId
func (mr *MockUserStoreMockRecorder) LoadMattermostUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadMattermostUserId", reflect.TypeOf((*MockUserStore)(nil).LoadMattermostUserId), arg0)
}

// LoadUser mocks base method
func (m *MockUserStore) LoadUser(arg0 string) (*store.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUser", arg0)
	ret0, _ := ret[0].(*store.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUser indicates an expected call of LoadUser
func (mr *MockUserStoreMockRecorder) LoadUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUser", reflect.TypeOf((*MockUserStore)(nil).LoadUser), arg0)
}

// LoadUserIndex mocks base method
func (m *MockUserStore) LoadUserIndex() ([]*store.UserShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadUserIndex")
	ret0, _ := ret[0].([]*store.UserShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadUserIndex indicates an expected call of LoadUserIndex
func (mr *MockUserStoreMockRecorder) LoadUserIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadUserIndex", reflect.TypeOf((*MockUserStore)(nil).LoadUserIndex))
}

// StoreUser mocks base method
func (m *MockUserStore) StoreUser(arg0 *store.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreUser indicates an expected call of StoreUser
func (mr *MockUserStoreMockRecorder) StoreUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreUser", reflect.TypeOf((*MockUserStore)(nil).StoreUser), arg0)
}
