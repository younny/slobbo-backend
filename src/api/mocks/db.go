package mocks

import (
	reflect "reflect"

	types "slobbo/src/types"

	gomock "github.com/golang/mock/gomock"
)

type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

func (m *MockClientInterface) Connect(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClientInterfaceMockRecorder) Connect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockClientInterface)(nil).Connect), arg0)
}

func (m *MockClientInterface) GetPosts(arg0 int) *types.PostList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", arg0)
	ret0, _ := ret[0].(*types.PostList)
	return ret0
}

func (mr *MockClientInterfaceMockRecorder) GetPosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockClientInterface)(nil).GetPosts), arg0)
}

func (m *MockClientInterface) GetPostByID(arg0 uint) *types.Post {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", arg0)
	ret0, _ := ret[0].(*types.Post)
	return ret0
}

func (mr *MockClientInterfaceMockRecorder) GetPostByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockClientInterface)(nil).GetPostByID), arg0)
}

func (m *MockClientInterface) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClientInterfaceMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockClientInterface)(nil).Ping), arg0)
}

func (m *MockClientInterface) CreatePost(arg0 *types.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClientInterfaceMockRecorder) CreatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockClientInterface)(nil).CreatePost), arg0)
}

func (m *MockClientInterface) UpdatePost(arg0 *types.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClientInterfaceMockRecorder) UpdatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockClientInterface)(nil).UpdatePost), arg0)
}

func (m *MockClientInterface) DeletePost(arg0 uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockClientInterfaceMockRecorder) DeletePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockClientInterface)(nil).DeletePost), arg0)
}
