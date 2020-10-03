// Code generated by MockGen. DO NOT EDIT.
// Source: dbmodels/twitter_user/querier.go

// Package mock_twitter_user is a generated GoMock package.
package mock_twitter_user

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	twitter_user "github.com/rickypai/cutie/examples/twitter/dbmodels/twitter_user"
	reflect "reflect"
)

// MockQuerier is a mock of Querier interface
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// GetIDByScreenName mocks base method
func (m *MockQuerier) GetIDByScreenName(ctx context.Context, screenName string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDByScreenName", ctx, screenName)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDByScreenName indicates an expected call of GetIDByScreenName
func (mr *MockQuerierMockRecorder) GetIDByScreenName(ctx, screenName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDByScreenName", reflect.TypeOf((*MockQuerier)(nil).GetIDByScreenName), ctx, screenName)
}

// GetScreenNameByID mocks base method
func (m *MockQuerier) GetScreenNameByID(ctx context.Context, id int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScreenNameByID", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScreenNameByID indicates an expected call of GetScreenNameByID
func (mr *MockQuerierMockRecorder) GetScreenNameByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScreenNameByID", reflect.TypeOf((*MockQuerier)(nil).GetScreenNameByID), ctx, id)
}

// GetTwitterUserByID mocks base method
func (m *MockQuerier) GetTwitterUserByID(ctx context.Context, id int64) (twitter_user.TwitterUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTwitterUserByID", ctx, id)
	ret0, _ := ret[0].(twitter_user.TwitterUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTwitterUserByID indicates an expected call of GetTwitterUserByID
func (mr *MockQuerierMockRecorder) GetTwitterUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTwitterUserByID", reflect.TypeOf((*MockQuerier)(nil).GetTwitterUserByID), ctx, id)
}

// GetTwitterUserByScreenName mocks base method
func (m *MockQuerier) GetTwitterUserByScreenName(ctx context.Context, screenName string) (twitter_user.TwitterUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTwitterUserByScreenName", ctx, screenName)
	ret0, _ := ret[0].(twitter_user.TwitterUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTwitterUserByScreenName indicates an expected call of GetTwitterUserByScreenName
func (mr *MockQuerierMockRecorder) GetTwitterUserByScreenName(ctx, screenName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTwitterUserByScreenName", reflect.TypeOf((*MockQuerier)(nil).GetTwitterUserByScreenName), ctx, screenName)
}
