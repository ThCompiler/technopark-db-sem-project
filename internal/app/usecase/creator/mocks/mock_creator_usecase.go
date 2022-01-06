// Code generated by MockGen. DO NOT EDIT.
// Source: tech-db-forum/internal/app/usecase/creator (interfaces: Usecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	io "io"
	models "tech-db-forum/internal/app/models"
	repository_files "tech-db-forum/internal/microservices/files/files/repository/files"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// CreatorUsecase is a mock of Usecase interface.
type CreatorUsecase struct {
	ctrl     *gomock.Controller
	recorder *CreatorUsecaseMockRecorder
}

// CreatorUsecaseMockRecorder is the mock recorder for CreatorUsecase.
type CreatorUsecaseMockRecorder struct {
	mock *CreatorUsecase
}

// NewCreatorUsecase creates a new mock instance.
func NewCreatorUsecase(ctrl *gomock.Controller) *CreatorUsecase {
	mock := &CreatorUsecase{ctrl: ctrl}
	mock.recorder = &CreatorUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *CreatorUsecase) EXPECT() *CreatorUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *CreatorUsecase) Create(arg0 *models.Creator) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *CreatorUsecaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*CreatorUsecase)(nil).Create), arg0)
}

// GetCreator mocks base method.
func (m *CreatorUsecase) GetCreator(arg0, arg1 int64) (*models.CreatorWithAwards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreator", arg0, arg1)
	ret0, _ := ret[0].(*models.CreatorWithAwards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreator indicates an expected call of GetCreator.
func (mr *CreatorUsecaseMockRecorder) GetCreator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreator", reflect.TypeOf((*CreatorUsecase)(nil).GetCreator), arg0, arg1)
}

// GetCreators mocks base method.
func (m *CreatorUsecase) GetCreators() ([]models.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreators")
	ret0, _ := ret[0].([]models.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreators indicates an expected call of GetCreators.
func (mr *CreatorUsecaseMockRecorder) GetCreators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreators", reflect.TypeOf((*CreatorUsecase)(nil).GetCreators))
}

// SearchCreators mocks base method.
func (m *CreatorUsecase) SearchCreators(arg0 *models.Pagination, arg1 string, arg2 ...string) ([]models.Creator, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchCreators", varargs...)
	ret0, _ := ret[0].([]models.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchCreators indicates an expected call of SearchCreators.
func (mr *CreatorUsecaseMockRecorder) SearchCreators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchCreators", reflect.TypeOf((*CreatorUsecase)(nil).SearchCreators), varargs...)
}

// UpdateAvatar mocks base method.
func (m *CreatorUsecase) UpdateAvatar(arg0 io.Reader, arg1 repository_files.FileName, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAvatar", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAvatar indicates an expected call of UpdateAvatar.
func (mr *CreatorUsecaseMockRecorder) UpdateAvatar(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAvatar", reflect.TypeOf((*CreatorUsecase)(nil).UpdateAvatar), arg0, arg1, arg2)
}

// UpdateCover mocks base method.
func (m *CreatorUsecase) UpdateCover(arg0 io.Reader, arg1 repository_files.FileName, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCover", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCover indicates an expected call of UpdateCover.
func (mr *CreatorUsecaseMockRecorder) UpdateCover(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCover", reflect.TypeOf((*CreatorUsecase)(nil).UpdateCover), arg0, arg1, arg2)
}
