// Code generated by MockGen. DO NOT EDIT.
// Source: tech-db-forum/internal/app/usecase/awards (interfaces: Usecase)

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	io "io"
	models "tech-db-forum/internal/app/models"
	repository_files "tech-db-forum/internal/microservices/files/files/repository/files"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// AwardsUsecase is a mock of Usecase interface.
type AwardsUsecase struct {
	ctrl     *gomock.Controller
	recorder *AwardsUsecaseMockRecorder
}

// AwardsUsecaseMockRecorder is the mock recorder for AwardsUsecase.
type AwardsUsecaseMockRecorder struct {
	mock *AwardsUsecase
}

// NewAwardsUsecase creates a new mock instance.
func NewAwardsUsecase(ctrl *gomock.Controller) *AwardsUsecase {
	mock := &AwardsUsecase{ctrl: ctrl}
	mock.recorder = &AwardsUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *AwardsUsecase) EXPECT() *AwardsUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *AwardsUsecase) Create(arg0 *models.Award) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *AwardsUsecaseMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*AwardsUsecase)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *AwardsUsecase) Delete(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *AwardsUsecaseMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*AwardsUsecase)(nil).Delete), arg0)
}

// GetAwards mocks base method.
func (m *AwardsUsecase) GetAwards(arg0 int64) ([]models.Award, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAwards", arg0)
	ret0, _ := ret[0].([]models.Award)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAwards indicates an expected call of GetAwards.
func (mr *AwardsUsecaseMockRecorder) GetAwards(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAwards", reflect.TypeOf((*AwardsUsecase)(nil).GetAwards), arg0)
}

// GetCreatorId mocks base method.
func (m *AwardsUsecase) GetCreatorId(arg0 int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreatorId", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreatorId indicates an expected call of GetCreatorId.
func (mr *AwardsUsecaseMockRecorder) GetCreatorId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreatorId", reflect.TypeOf((*AwardsUsecase)(nil).GetCreatorId), arg0)
}

// Update mocks base method.
func (m *AwardsUsecase) Update(arg0 *models.Award) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *AwardsUsecaseMockRecorder) Update(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*AwardsUsecase)(nil).Update), arg0)
}

// UpdateCover mocks base method.
func (m *AwardsUsecase) UpdateCover(arg0 io.Reader, arg1 repository_files.FileName, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCover", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCover indicates an expected call of UpdateCover.
func (mr *AwardsUsecaseMockRecorder) UpdateCover(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCover", reflect.TypeOf((*AwardsUsecase)(nil).UpdateCover), arg0, arg1, arg2)
}