// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-msoffice/server/api (interfaces: Availability)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	gomock "github.com/golang/mock/gomock"
	remote "github.com/mattermost/mattermost-plugin-msoffice/server/remote"
	reflect "reflect"
)

// MockAvailability is a mock of Availability interface
type MockAvailability struct {
	ctrl     *gomock.Controller
	recorder *MockAvailabilityMockRecorder
}

// MockAvailabilityMockRecorder is the mock recorder for MockAvailability
type MockAvailabilityMockRecorder struct {
	mock *MockAvailability
}

// NewMockAvailability creates a new mock instance
func NewMockAvailability(ctrl *gomock.Controller) *MockAvailability {
	mock := &MockAvailability{ctrl: ctrl}
	mock.recorder = &MockAvailabilityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAvailability) EXPECT() *MockAvailabilityMockRecorder {
	return m.recorder
}

// GetUserAvailabilities mocks base method
func (m *MockAvailability) GetUserAvailabilities(arg0 string, arg1 []string) ([]*remote.ScheduleInformation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAvailabilities", arg0, arg1)
	ret0, _ := ret[0].([]*remote.ScheduleInformation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAvailabilities indicates an expected call of GetUserAvailabilities
func (mr *MockAvailabilityMockRecorder) GetUserAvailabilities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAvailabilities", reflect.TypeOf((*MockAvailability)(nil).GetUserAvailabilities), arg0, arg1)
}

// SyncStatusForAllUsers mocks base method
func (m *MockAvailability) SyncStatusForAllUsers() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatusForAllUsers")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncStatusForAllUsers indicates an expected call of SyncStatusForAllUsers
func (mr *MockAvailabilityMockRecorder) SyncStatusForAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatusForAllUsers", reflect.TypeOf((*MockAvailability)(nil).SyncStatusForAllUsers))
}

// SyncStatusForSingleUser mocks base method
func (m *MockAvailability) SyncStatusForSingleUser(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatusForSingleUser", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncStatusForSingleUser indicates an expected call of SyncStatusForSingleUser
func (mr *MockAvailabilityMockRecorder) SyncStatusForSingleUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatusForSingleUser", reflect.TypeOf((*MockAvailability)(nil).SyncStatusForSingleUser), arg0)
}
