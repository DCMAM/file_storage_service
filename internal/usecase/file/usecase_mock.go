// Code generated by MockGen. DO NOT EDIT.
// Source: ./usecase.go

// Package file is a generated GoMock package.
package file

import (
	models "file_storage_service/internal/models"
	multipart "mime/multipart"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockfileProvider is a mock of fileProvider interface.
type MockfileProvider struct {
	ctrl     *gomock.Controller
	recorder *MockfileProviderMockRecorder
}

// MockfileProviderMockRecorder is the mock recorder for MockfileProvider.
type MockfileProviderMockRecorder struct {
	mock *MockfileProvider
}

// NewMockfileProvider creates a new mock instance.
func NewMockfileProvider(ctrl *gomock.Controller) *MockfileProvider {
	mock := &MockfileProvider{ctrl: ctrl}
	mock.recorder = &MockfileProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfileProvider) EXPECT() *MockfileProviderMockRecorder {
	return m.recorder
}

// DonwloadFile mocks base method.
func (m *MockfileProvider) DonwloadFile(path string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DonwloadFile", path)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DonwloadFile indicates an expected call of DonwloadFile.
func (mr *MockfileProviderMockRecorder) DonwloadFile(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DonwloadFile", reflect.TypeOf((*MockfileProvider)(nil).DonwloadFile), path)
}

// UploadFile mocks base method.
func (m *MockfileProvider) UploadFile(file multipart.File, fileLocation string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", file, fileLocation)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockfileProviderMockRecorder) UploadFile(file, fileLocation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockfileProvider)(nil).UploadFile), file, fileLocation)
}

// MockfileDBProvider is a mock of fileDBProvider interface.
type MockfileDBProvider struct {
	ctrl     *gomock.Controller
	recorder *MockfileDBProviderMockRecorder
}

// MockfileDBProviderMockRecorder is the mock recorder for MockfileDBProvider.
type MockfileDBProviderMockRecorder struct {
	mock *MockfileDBProvider
}

// NewMockfileDBProvider creates a new mock instance.
func NewMockfileDBProvider(ctrl *gomock.Controller) *MockfileDBProvider {
	mock := &MockfileDBProvider{ctrl: ctrl}
	mock.recorder = &MockfileDBProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockfileDBProvider) EXPECT() *MockfileDBProviderMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockfileDBProvider) GetAll() ([]models.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockfileDBProviderMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockfileDBProvider)(nil).GetAll))
}

// SetFile mocks base method.
func (m *MockfileDBProvider) SetFile(url, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFile", url, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFile indicates an expected call of SetFile.
func (mr *MockfileDBProviderMockRecorder) SetFile(url, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFile", reflect.TypeOf((*MockfileDBProvider)(nil).SetFile), url, username)
}
