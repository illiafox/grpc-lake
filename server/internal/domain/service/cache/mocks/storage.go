// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "server/internal/domain/entity"
)

// MockCacheStorage is a mock of CacheStorage interface.
type MockCacheStorage struct {
	ctrl     *gomock.Controller
	recorder *MockCacheStorageMockRecorder
}

// MockCacheStorageMockRecorder is the mock recorder for MockCacheStorage.
type MockCacheStorageMockRecorder struct {
	mock *MockCacheStorage
}

// NewMockCacheStorage creates a new mock instance.
func NewMockCacheStorage(ctrl *gomock.Controller) *MockCacheStorage {
	mock := &MockCacheStorage{ctrl: ctrl}
	mock.recorder = &MockCacheStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheStorage) EXPECT() *MockCacheStorageMockRecorder {
	return m.recorder
}

// DeleteItem mocks base method.
func (m *MockCacheStorage) DeleteItem(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockCacheStorageMockRecorder) DeleteItem(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockCacheStorage)(nil).DeleteItem), ctx, id)
}

// GetItem mocks base method.
func (m *MockCacheStorage) GetItem(ctx context.Context, id string) (entity.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", ctx, id)
	ret0, _ := ret[0].(entity.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockCacheStorageMockRecorder) GetItem(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockCacheStorage)(nil).GetItem), ctx, id)
}

// SetItem mocks base method.
func (m *MockCacheStorage) SetItem(ctx context.Context, id string, item entity.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetItem", ctx, id, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetItem indicates an expected call of SetItem.
func (mr *MockCacheStorageMockRecorder) SetItem(ctx, id, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetItem", reflect.TypeOf((*MockCacheStorage)(nil).SetItem), ctx, id, item)
}

// MockItemStorage is a mock of ItemStorage interface.
type MockItemStorage struct {
	ctrl     *gomock.Controller
	recorder *MockItemStorageMockRecorder
}

// MockItemStorageMockRecorder is the mock recorder for MockItemStorage.
type MockItemStorageMockRecorder struct {
	mock *MockItemStorage
}

// NewMockItemStorage creates a new mock instance.
func NewMockItemStorage(ctrl *gomock.Controller) *MockItemStorage {
	mock := &MockItemStorage{ctrl: ctrl}
	mock.recorder = &MockItemStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemStorage) EXPECT() *MockItemStorageMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockItemStorage) CreateItem(ctx context.Context, name string, data []byte, description string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", ctx, name, data, description)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockItemStorageMockRecorder) CreateItem(ctx, name, data, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockItemStorage)(nil).CreateItem), ctx, name, data, description)
}

// DeleteItem mocks base method.
func (m *MockItemStorage) DeleteItem(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockItemStorageMockRecorder) DeleteItem(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockItemStorage)(nil).DeleteItem), ctx, id)
}

// GetItem mocks base method.
func (m *MockItemStorage) GetItem(ctx context.Context, id string) (entity.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", ctx, id)
	ret0, _ := ret[0].(entity.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockItemStorageMockRecorder) GetItem(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockItemStorage)(nil).GetItem), ctx, id)
}

// UpdateItem mocks base method.
func (m *MockItemStorage) UpdateItem(ctx context.Context, id string, item entity.Item) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", ctx, id, item)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockItemStorageMockRecorder) UpdateItem(ctx, id, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockItemStorage)(nil).UpdateItem), ctx, id, item)
}