// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mbaxamb3/pantopia/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/mbaxamb3/pantopia/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreateLead mocks base method.
func (m *MockStore) CreateLead(arg0 context.Context, arg1 db.CreateLeadParams) (db.Lead, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLead", arg0, arg1)
	ret0, _ := ret[0].(db.Lead)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLead indicates an expected call of CreateLead.
func (mr *MockStoreMockRecorder) CreateLead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLead", reflect.TypeOf((*MockStore)(nil).CreateLead), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// DeleteLead mocks base method.
func (m *MockStore) DeleteLead(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLead", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLead indicates an expected call of DeleteLead.
func (mr *MockStoreMockRecorder) DeleteLead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLead", reflect.TypeOf((*MockStore)(nil).DeleteLead), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockStore) GetAccount(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockStoreMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockStore)(nil).GetAccount), arg0, arg1)
}

// GetLead mocks base method.
func (m *MockStore) GetLead(arg0 context.Context, arg1 int64) (db.Lead, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLead", arg0, arg1)
	ret0, _ := ret[0].(db.Lead)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLead indicates an expected call of GetLead.
func (mr *MockStoreMockRecorder) GetLead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLead", reflect.TypeOf((*MockStore)(nil).GetLead), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// ListLeads mocks base method.
func (m *MockStore) ListLeads(arg0 context.Context, arg1 db.ListLeadsParams) ([]db.Lead, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLeads", arg0, arg1)
	ret0, _ := ret[0].([]db.Lead)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLeads indicates an expected call of ListLeads.
func (mr *MockStoreMockRecorder) ListLeads(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLeads", reflect.TypeOf((*MockStore)(nil).ListLeads), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdateLead mocks base method.
func (m *MockStore) UpdateLead(arg0 context.Context, arg1 db.UpdateLeadParams) (db.Lead, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLead", arg0, arg1)
	ret0, _ := ret[0].(db.Lead)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLead indicates an expected call of UpdateLead.
func (mr *MockStoreMockRecorder) UpdateLead(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLead", reflect.TypeOf((*MockStore)(nil).UpdateLead), arg0, arg1)
}
