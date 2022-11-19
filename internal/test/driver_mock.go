// Code generated by MockGen. DO NOT EDIT.
// Source: ./service/driverService.go

// Package test is a generated GoMock package.
package test

import (
	model "InnoTaxi-Driver/internal/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddNewDriver mocks base method.
func (m *MockRepository) AddNewDriver(ctx context.Context, driver model.Driver) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewDriver", ctx, driver)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNewDriver indicates an expected call of AddNewDriver.
func (mr *MockRepositoryMockRecorder) AddNewDriver(ctx, driver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewDriver", reflect.TypeOf((*MockRepository)(nil).AddNewDriver), ctx, driver)
}

// Exists mocks base method.
func (m *MockRepository) Exists(ctx context.Context, user model.Login) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, user)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockRepositoryMockRecorder) Exists(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockRepository)(nil).Exists), ctx, user)
}

// GetAllDrivers mocks base method.
func (m *MockRepository) GetAllDrivers(ctx context.Context) ([]model.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDrivers", ctx)
	ret0, _ := ret[0].([]model.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDrivers indicates an expected call of GetAllDrivers.
func (mr *MockRepositoryMockRecorder) GetAllDrivers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDrivers", reflect.TypeOf((*MockRepository)(nil).GetAllDrivers), ctx)
}

// GetAllFreeDrivers mocks base method.
func (m *MockRepository) GetAllFreeDrivers(ctx context.Context) ([]model.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFreeDrivers", ctx)
	ret0, _ := ret[0].([]model.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFreeDrivers indicates an expected call of GetAllFreeDrivers.
func (mr *MockRepositoryMockRecorder) GetAllFreeDrivers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFreeDrivers", reflect.TypeOf((*MockRepository)(nil).GetAllFreeDrivers), ctx)
}

// GetDriverPhoneAndPasswordByPhone mocks base method.
func (m *MockRepository) GetDriverPhoneAndPasswordByPhone(ctx context.Context, phone string) (model.Driver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDriverPhoneAndPasswordByPhone", ctx, phone)
	ret0, _ := ret[0].(model.Driver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDriverPhoneAndPasswordByPhone indicates an expected call of GetDriverPhoneAndPasswordByPhone.
func (mr *MockRepositoryMockRecorder) GetDriverPhoneAndPasswordByPhone(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDriverPhoneAndPasswordByPhone", reflect.TypeOf((*MockRepository)(nil).GetDriverPhoneAndPasswordByPhone), ctx, phone)
}

// UpdateDriverStatus mocks base method.
func (m *MockRepository) UpdateDriverStatus(ctx context.Context, user model.Driver) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDriverStatus", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDriverStatus indicates an expected call of UpdateDriverStatus.
func (mr *MockRepositoryMockRecorder) UpdateDriverStatus(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDriverStatus", reflect.TypeOf((*MockRepository)(nil).UpdateDriverStatus), ctx, user)
}

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// SendMessage mocks base method.
func (m *MockBroker) SendMessage(value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", value)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockBrokerMockRecorder) SendMessage(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockBroker)(nil).SendMessage), value)
}