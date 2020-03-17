// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-plugin-mscalendar/server/mscalendar (interfaces: MSCalendar)

// Package mock_mscalendar is a generated GoMock package.
package mock_mscalendar

import (
	gomock "github.com/golang/mock/gomock"
	mscalendar "github.com/mattermost/mattermost-plugin-mscalendar/server/mscalendar"
	remote "github.com/mattermost/mattermost-plugin-mscalendar/server/remote"
	store "github.com/mattermost/mattermost-plugin-mscalendar/server/store"
	reflect "reflect"
	time "time"
)

// MockMSCalendar is a mock of MSCalendar interface
type MockMSCalendar struct {
	ctrl     *gomock.Controller
	recorder *MockMSCalendarMockRecorder
}

// MockMSCalendarMockRecorder is the mock recorder for MockMSCalendar
type MockMSCalendarMockRecorder struct {
	mock *MockMSCalendar
}

// NewMockMSCalendar creates a new mock instance
func NewMockMSCalendar(ctrl *gomock.Controller) *MockMSCalendar {
	mock := &MockMSCalendar{ctrl: ctrl}
	mock.recorder = &MockMSCalendarMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMSCalendar) EXPECT() *MockMSCalendarMockRecorder {
	return m.recorder
}

// AcceptEvent mocks base method
func (m *MockMSCalendar) AcceptEvent(arg0 *mscalendar.User, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptEvent indicates an expected call of AcceptEvent
func (mr *MockMSCalendarMockRecorder) AcceptEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptEvent", reflect.TypeOf((*MockMSCalendar)(nil).AcceptEvent), arg0, arg1)
}

// AfterDisconnect mocks base method
func (m *MockMSCalendar) AfterDisconnect(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterDisconnect", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterDisconnect indicates an expected call of AfterDisconnect
func (mr *MockMSCalendarMockRecorder) AfterDisconnect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterDisconnect", reflect.TypeOf((*MockMSCalendar)(nil).AfterDisconnect), arg0)
}

// AfterSuccessfullyConnect mocks base method
func (m *MockMSCalendar) AfterSuccessfullyConnect(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterSuccessfullyConnect", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterSuccessfullyConnect indicates an expected call of AfterSuccessfullyConnect
func (mr *MockMSCalendarMockRecorder) AfterSuccessfullyConnect(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterSuccessfullyConnect", reflect.TypeOf((*MockMSCalendar)(nil).AfterSuccessfullyConnect), arg0, arg1)
}

// CreateCalendar mocks base method
func (m *MockMSCalendar) CreateCalendar(arg0 *mscalendar.User, arg1 *remote.Calendar) (*remote.Calendar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCalendar", arg0, arg1)
	ret0, _ := ret[0].(*remote.Calendar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCalendar indicates an expected call of CreateCalendar
func (mr *MockMSCalendarMockRecorder) CreateCalendar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCalendar", reflect.TypeOf((*MockMSCalendar)(nil).CreateCalendar), arg0, arg1)
}

// CreateEvent mocks base method
func (m *MockMSCalendar) CreateEvent(arg0 *mscalendar.User, arg1 *remote.Event, arg2 []string) (*remote.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(*remote.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent
func (mr *MockMSCalendarMockRecorder) CreateEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockMSCalendar)(nil).CreateEvent), arg0, arg1, arg2)
}

// CreateMyEventSubscription mocks base method
func (m *MockMSCalendar) CreateMyEventSubscription() (*store.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMyEventSubscription")
	ret0, _ := ret[0].(*store.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMyEventSubscription indicates an expected call of CreateMyEventSubscription
func (mr *MockMSCalendarMockRecorder) CreateMyEventSubscription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMyEventSubscription", reflect.TypeOf((*MockMSCalendar)(nil).CreateMyEventSubscription))
}

// DeclineEvent mocks base method
func (m *MockMSCalendar) DeclineEvent(arg0 *mscalendar.User, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeclineEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeclineEvent indicates an expected call of DeclineEvent
func (mr *MockMSCalendarMockRecorder) DeclineEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeclineEvent", reflect.TypeOf((*MockMSCalendar)(nil).DeclineEvent), arg0, arg1)
}

// DeleteCalendar mocks base method
func (m *MockMSCalendar) DeleteCalendar(arg0 *mscalendar.User, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCalendar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCalendar indicates an expected call of DeleteCalendar
func (mr *MockMSCalendarMockRecorder) DeleteCalendar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCalendar", reflect.TypeOf((*MockMSCalendar)(nil).DeleteCalendar), arg0, arg1)
}

// DeleteMyEventSubscription mocks base method
func (m *MockMSCalendar) DeleteMyEventSubscription() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMyEventSubscription")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMyEventSubscription indicates an expected call of DeleteMyEventSubscription
func (mr *MockMSCalendarMockRecorder) DeleteMyEventSubscription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMyEventSubscription", reflect.TypeOf((*MockMSCalendar)(nil).DeleteMyEventSubscription))
}

// DeleteOrphanedSubscription mocks base method
func (m *MockMSCalendar) DeleteOrphanedSubscription(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrphanedSubscription", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrphanedSubscription indicates an expected call of DeleteOrphanedSubscription
func (mr *MockMSCalendarMockRecorder) DeleteOrphanedSubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrphanedSubscription", reflect.TypeOf((*MockMSCalendar)(nil).DeleteOrphanedSubscription), arg0)
}

// DisconnectUser mocks base method
func (m *MockMSCalendar) DisconnectUser(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisconnectUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisconnectUser indicates an expected call of DisconnectUser
func (mr *MockMSCalendarMockRecorder) DisconnectUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisconnectUser", reflect.TypeOf((*MockMSCalendar)(nil).DisconnectUser), arg0)
}

// FindMeetingTimes mocks base method
func (m *MockMSCalendar) FindMeetingTimes(arg0 *mscalendar.User, arg1 *remote.FindMeetingTimesParameters) (*remote.MeetingTimeSuggestionResults, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMeetingTimes", arg0, arg1)
	ret0, _ := ret[0].(*remote.MeetingTimeSuggestionResults)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMeetingTimes indicates an expected call of FindMeetingTimes
func (mr *MockMSCalendarMockRecorder) FindMeetingTimes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMeetingTimes", reflect.TypeOf((*MockMSCalendar)(nil).FindMeetingTimes), arg0, arg1)
}

// GetActingUser mocks base method
func (m *MockMSCalendar) GetActingUser() *mscalendar.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActingUser")
	ret0, _ := ret[0].(*mscalendar.User)
	return ret0
}

// GetActingUser indicates an expected call of GetActingUser
func (mr *MockMSCalendarMockRecorder) GetActingUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActingUser", reflect.TypeOf((*MockMSCalendar)(nil).GetActingUser))
}

// GetAvailabilities mocks base method
func (m *MockMSCalendar) GetAvailabilities(arg0 string, arg1 []string) ([]*remote.ScheduleInformation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailabilities", arg0, arg1)
	ret0, _ := ret[0].([]*remote.ScheduleInformation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailabilities indicates an expected call of GetAvailabilities
func (mr *MockMSCalendarMockRecorder) GetAvailabilities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailabilities", reflect.TypeOf((*MockMSCalendar)(nil).GetAvailabilities), arg0, arg1)
}

// GetCalendars mocks base method
func (m *MockMSCalendar) GetCalendars(arg0 *mscalendar.User) ([]*remote.Calendar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCalendars", arg0)
	ret0, _ := ret[0].([]*remote.Calendar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCalendars indicates an expected call of GetCalendars
func (mr *MockMSCalendarMockRecorder) GetCalendars(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCalendars", reflect.TypeOf((*MockMSCalendar)(nil).GetCalendars), arg0)
}

// GetRemoteUser mocks base method
func (m *MockMSCalendar) GetRemoteUser(arg0 string) (*remote.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteUser", arg0)
	ret0, _ := ret[0].(*remote.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRemoteUser indicates an expected call of GetRemoteUser
func (mr *MockMSCalendarMockRecorder) GetRemoteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteUser", reflect.TypeOf((*MockMSCalendar)(nil).GetRemoteUser), arg0)
}

// GetTimezone mocks base method
func (m *MockMSCalendar) GetTimezone(arg0 *mscalendar.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimezone", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimezone indicates an expected call of GetTimezone
func (mr *MockMSCalendarMockRecorder) GetTimezone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimezone", reflect.TypeOf((*MockMSCalendar)(nil).GetTimezone), arg0)
}

// IsAuthorizedAdmin mocks base method
func (m *MockMSCalendar) IsAuthorizedAdmin(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorizedAdmin", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthorizedAdmin indicates an expected call of IsAuthorizedAdmin
func (mr *MockMSCalendarMockRecorder) IsAuthorizedAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorizedAdmin", reflect.TypeOf((*MockMSCalendar)(nil).IsAuthorizedAdmin), arg0)
}

// ListRemoteSubscriptions mocks base method
func (m *MockMSCalendar) ListRemoteSubscriptions() ([]*remote.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRemoteSubscriptions")
	ret0, _ := ret[0].([]*remote.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRemoteSubscriptions indicates an expected call of ListRemoteSubscriptions
func (mr *MockMSCalendarMockRecorder) ListRemoteSubscriptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRemoteSubscriptions", reflect.TypeOf((*MockMSCalendar)(nil).ListRemoteSubscriptions))
}

// LoadMyEventSubscription mocks base method
func (m *MockMSCalendar) LoadMyEventSubscription() (*store.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadMyEventSubscription")
	ret0, _ := ret[0].(*store.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadMyEventSubscription indicates an expected call of LoadMyEventSubscription
func (mr *MockMSCalendarMockRecorder) LoadMyEventSubscription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadMyEventSubscription", reflect.TypeOf((*MockMSCalendar)(nil).LoadMyEventSubscription))
}

// RenewMyEventSubscription mocks base method
func (m *MockMSCalendar) RenewMyEventSubscription() (*store.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenewMyEventSubscription")
	ret0, _ := ret[0].(*store.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenewMyEventSubscription indicates an expected call of RenewMyEventSubscription
func (mr *MockMSCalendarMockRecorder) RenewMyEventSubscription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenewMyEventSubscription", reflect.TypeOf((*MockMSCalendar)(nil).RenewMyEventSubscription))
}

// RespondToEvent mocks base method
func (m *MockMSCalendar) RespondToEvent(arg0 *mscalendar.User, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondToEvent", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondToEvent indicates an expected call of RespondToEvent
func (mr *MockMSCalendarMockRecorder) RespondToEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondToEvent", reflect.TypeOf((*MockMSCalendar)(nil).RespondToEvent), arg0, arg1, arg2)
}

// SyncStatus mocks base method
func (m *MockMSCalendar) SyncStatus(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncStatus indicates an expected call of SyncStatus
func (mr *MockMSCalendarMockRecorder) SyncStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockMSCalendar)(nil).SyncStatus), arg0)
}

// SyncStatusAll mocks base method
func (m *MockMSCalendar) SyncStatusAll() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatusAll")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncStatusAll indicates an expected call of SyncStatusAll
func (mr *MockMSCalendarMockRecorder) SyncStatusAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatusAll", reflect.TypeOf((*MockMSCalendar)(nil).SyncStatusAll))
}

// TentativelyAcceptEvent mocks base method
func (m *MockMSCalendar) TentativelyAcceptEvent(arg0 *mscalendar.User, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TentativelyAcceptEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TentativelyAcceptEvent indicates an expected call of TentativelyAcceptEvent
func (mr *MockMSCalendarMockRecorder) TentativelyAcceptEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TentativelyAcceptEvent", reflect.TypeOf((*MockMSCalendar)(nil).TentativelyAcceptEvent), arg0, arg1)
}

// ViewCalendar mocks base method
func (m *MockMSCalendar) ViewCalendar(arg0 *mscalendar.User, arg1, arg2 time.Time) ([]*remote.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewCalendar", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*remote.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewCalendar indicates an expected call of ViewCalendar
func (mr *MockMSCalendarMockRecorder) ViewCalendar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewCalendar", reflect.TypeOf((*MockMSCalendar)(nil).ViewCalendar), arg0, arg1, arg2)
}

// Welcome mocks base method
func (m *MockMSCalendar) Welcome(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Welcome", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Welcome indicates an expected call of Welcome
func (mr *MockMSCalendarMockRecorder) Welcome(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Welcome", reflect.TypeOf((*MockMSCalendar)(nil).Welcome), arg0)
}

// WelcomeFlowEnd mocks base method
func (m *MockMSCalendar) WelcomeFlowEnd(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WelcomeFlowEnd", arg0)
}

// WelcomeFlowEnd indicates an expected call of WelcomeFlowEnd
func (mr *MockMSCalendarMockRecorder) WelcomeFlowEnd(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WelcomeFlowEnd", reflect.TypeOf((*MockMSCalendar)(nil).WelcomeFlowEnd), arg0)
}
