// Code generated by MockGen. DO NOT EDIT.
// Source: article.go

// Package internal is a generated GoMock package.
package internal

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockArticleStorage is a mock of ArticleStorage interface.
type MockArticleStorage struct {
	ctrl     *gomock.Controller
	recorder *MockArticleStorageMockRecorder
}

// MockArticleStorageMockRecorder is the mock recorder for MockArticleStorage.
type MockArticleStorageMockRecorder struct {
	mock *MockArticleStorage
}

// NewMockArticleStorage creates a new mock instance.
func NewMockArticleStorage(ctrl *gomock.Controller) *MockArticleStorage {
	mock := &MockArticleStorage{ctrl: ctrl}
	mock.recorder = &MockArticleStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleStorage) EXPECT() *MockArticleStorageMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockArticleStorage) Delete(id ArticleId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockArticleStorageMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockArticleStorage)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockArticleStorage) Get(id ArticleId) (*Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockArticleStorageMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockArticleStorage)(nil).Get), id)
}

// Set mocks base method.
func (m *MockArticleStorage) Set(id ArticleId, title, content string) (*Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", id, title, content)
	ret0, _ := ret[0].(*Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Set indicates an expected call of Set.
func (mr *MockArticleStorageMockRecorder) Set(id, title, content interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockArticleStorage)(nil).Set), id, title, content)
}
