// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/olenka--91/reminder-app/internal/domain"
)

// MockRemind is a mock of Remind interface.
type MockRemind struct {
	ctrl     *gomock.Controller
	recorder *MockRemindMockRecorder
}

// MockRemindMockRecorder is the mock recorder for MockRemind.
type MockRemindMockRecorder struct {
	mock *MockRemind
}

// NewMockRemind creates a new mock instance.
func NewMockRemind(ctrl *gomock.Controller) *MockRemind {
	mock := &MockRemind{ctrl: ctrl}
	mock.recorder = &MockRemindMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemind) EXPECT() *MockRemindMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRemind) Create(userID int, remind domain.Remind) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userID, remind)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRemindMockRecorder) Create(userID, remind interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRemind)(nil).Create), userID, remind)
}

// Delete mocks base method.
func (m *MockRemind) Delete(userID, remindID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userID, remindID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRemindMockRecorder) Delete(userID, remindID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRemind)(nil).Delete), userID, remindID)
}

// GetAll mocks base method.
func (m *MockRemind) GetAll(userID int) ([]domain.Remind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userID)
	ret0, _ := ret[0].([]domain.Remind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRemindMockRecorder) GetAll(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRemind)(nil).GetAll), userID)
}

// GetByID mocks base method.
func (m *MockRemind) GetByID(userID, remindID int) (domain.Remind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", userID, remindID)
	ret0, _ := ret[0].(domain.Remind)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockRemindMockRecorder) GetByID(userID, remindID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockRemind)(nil).GetByID), userID, remindID)
}

// Update mocks base method.
func (m *MockRemind) Update(userID, remindID int, input domain.RemindUpdateInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userID, remindID, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockRemindMockRecorder) Update(userID, remindID, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRemind)(nil).Update), userID, remindID, input)
}

// MockAuthorization is a mock of Authorization interface.
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization.
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance.
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockAuthorization) CreateUser(user domain.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

// GenerateToken mocks base method.
func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

// ParseToken mocks base method.
func (m *MockAuthorization) ParseToken(accessToken string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", accessToken)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseToken indicates an expected call of ParseToken.
func (mr *MockAuthorizationMockRecorder) ParseToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), accessToken)
}