// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/class.go
//mockgen -source=internal/repositories/class.go -destination=internal/mocks/mock_class_repository.go -package=mocks

// Package mocks is a generated GoMock package.
package mocks

import (
	models "glofox-backend/internal/models"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockClassRepository is a mock of ClassRepository interface.
type MockClassRepository struct {
	ctrl     *gomock.Controller
	recorder *MockClassRepositoryMockRecorder
}

// MockClassRepositoryMockRecorder is the mock recorder for MockClassRepository.
type MockClassRepositoryMockRecorder struct {
	mock *MockClassRepository
}

// NewMockClassRepository creates a new mock instance.
func NewMockClassRepository(ctrl *gomock.Controller) *MockClassRepository {
	mock := &MockClassRepository{ctrl: ctrl}
	mock.recorder = &MockClassRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClassRepository) EXPECT() *MockClassRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClassRepository) Create(class *models.Class) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", class)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockClassRepositoryMockRecorder) Create(class interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClassRepository)(nil).Create), class)
}

// GetAll mocks base method.
func (m *MockClassRepository) GetAll() []*models.Class {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*models.Class)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockClassRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockClassRepository)(nil).GetAll))
}

// GetByDate mocks base method.
func (m *MockClassRepository) GetByDate(date time.Time) []*models.Class {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDate", date)
	ret0, _ := ret[0].([]*models.Class)
	return ret0
}

// GetByDate indicates an expected call of GetByDate.
func (mr *MockClassRepositoryMockRecorder) GetByDate(date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDate", reflect.TypeOf((*MockClassRepository)(nil).GetByDate), date)
}

// GetByID mocks base method.
func (m *MockClassRepository) GetByID(id string) (*models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockClassRepositoryMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockClassRepository)(nil).GetByID), id)
}
