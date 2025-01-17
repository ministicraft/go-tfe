// Code generated by MockGen. DO NOT EDIT.
// Source: workspace.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockWorkspaces is a mock of Workspaces interface.
type MockWorkspaces struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspacesMockRecorder
}

// MockWorkspacesMockRecorder is the mock recorder for MockWorkspaces.
type MockWorkspacesMockRecorder struct {
	mock *MockWorkspaces
}

// NewMockWorkspaces creates a new mock instance.
func NewMockWorkspaces(ctrl *gomock.Controller) *MockWorkspaces {
	mock := &MockWorkspaces{ctrl: ctrl}
	mock.recorder = &MockWorkspacesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspaces) EXPECT() *MockWorkspacesMockRecorder {
	return m.recorder
}

// AddRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) AddRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceAddRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRemoteStateConsumers indicates an expected call of AddRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) AddRemoteStateConsumers(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).AddRemoteStateConsumers), ctx, workspaceID, options)
}

// AddTags mocks base method.
func (m *MockWorkspaces) AddTags(ctx context.Context, workspaceID string, options tfe.WorkspaceAddTagsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTags", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTags indicates an expected call of AddTags.
func (mr *MockWorkspacesMockRecorder) AddTags(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTags", reflect.TypeOf((*MockWorkspaces)(nil).AddTags), ctx, workspaceID, options)
}

// AssignSSHKey mocks base method.
func (m *MockWorkspaces) AssignSSHKey(ctx context.Context, workspaceID string, options tfe.WorkspaceAssignSSHKeyOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignSSHKey", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignSSHKey indicates an expected call of AssignSSHKey.
func (mr *MockWorkspacesMockRecorder) AssignSSHKey(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignSSHKey", reflect.TypeOf((*MockWorkspaces)(nil).AssignSSHKey), ctx, workspaceID, options)
}

// Create mocks base method.
func (m *MockWorkspaces) Create(ctx context.Context, organization string, options tfe.WorkspaceCreateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWorkspacesMockRecorder) Create(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWorkspaces)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockWorkspaces) Delete(ctx context.Context, organization, workspace string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, organization, workspace)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWorkspacesMockRecorder) Delete(ctx, organization, workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWorkspaces)(nil).Delete), ctx, organization, workspace)
}

// DeleteByID mocks base method.
func (m *MockWorkspaces) DeleteByID(ctx context.Context, workspaceID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, workspaceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockWorkspacesMockRecorder) DeleteByID(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockWorkspaces)(nil).DeleteByID), ctx, workspaceID)
}

// ForceUnlock mocks base method.
func (m *MockWorkspaces) ForceUnlock(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceUnlock", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ForceUnlock indicates an expected call of ForceUnlock.
func (mr *MockWorkspacesMockRecorder) ForceUnlock(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceUnlock", reflect.TypeOf((*MockWorkspaces)(nil).ForceUnlock), ctx, workspaceID)
}

// List mocks base method.
func (m *MockWorkspaces) List(ctx context.Context, organization string, options tfe.WorkspaceListOptions) (*tfe.WorkspaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*tfe.WorkspaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockWorkspacesMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWorkspaces)(nil).List), ctx, organization, options)
}

// Lock mocks base method.
func (m *MockWorkspaces) Lock(ctx context.Context, workspaceID string, options tfe.WorkspaceLockOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockWorkspacesMockRecorder) Lock(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockWorkspaces)(nil).Lock), ctx, workspaceID, options)
}

// Read mocks base method.
func (m *MockWorkspaces) Read(ctx context.Context, organization, workspace string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, organization, workspace)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockWorkspacesMockRecorder) Read(ctx, organization, workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockWorkspaces)(nil).Read), ctx, organization, workspace)
}

// ReadByID mocks base method.
func (m *MockWorkspaces) ReadByID(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByID", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByID indicates an expected call of ReadByID.
func (mr *MockWorkspacesMockRecorder) ReadByID(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByID", reflect.TypeOf((*MockWorkspaces)(nil).ReadByID), ctx, workspaceID)
}

// ReadByIDWithOptions mocks base method.
func (m *MockWorkspaces) ReadByIDWithOptions(ctx context.Context, workspaceID string, options *tfe.WorkspaceReadOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadByIDWithOptions", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadByIDWithOptions indicates an expected call of ReadByIDWithOptions.
func (mr *MockWorkspacesMockRecorder) ReadByIDWithOptions(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadByIDWithOptions", reflect.TypeOf((*MockWorkspaces)(nil).ReadByIDWithOptions), ctx, workspaceID, options)
}

// ReadWithOptions mocks base method.
func (m *MockWorkspaces) ReadWithOptions(ctx context.Context, organization, workspace string, options *tfe.WorkspaceReadOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadWithOptions", ctx, organization, workspace, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadWithOptions indicates an expected call of ReadWithOptions.
func (mr *MockWorkspacesMockRecorder) ReadWithOptions(ctx, organization, workspace, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadWithOptions", reflect.TypeOf((*MockWorkspaces)(nil).ReadWithOptions), ctx, organization, workspace, options)
}

// Readme mocks base method.
func (m *MockWorkspaces) Readme(ctx context.Context, workspaceID string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Readme", ctx, workspaceID)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Readme indicates an expected call of Readme.
func (mr *MockWorkspacesMockRecorder) Readme(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Readme", reflect.TypeOf((*MockWorkspaces)(nil).Readme), ctx, workspaceID)
}

// RemoteStateConsumers mocks base method.
func (m *MockWorkspaces) RemoteStateConsumers(ctx context.Context, workspaceID string, options *tfe.RemoteStateConsumersListOptions) (*tfe.WorkspaceList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.WorkspaceList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteStateConsumers indicates an expected call of RemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) RemoteStateConsumers(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).RemoteStateConsumers), ctx, workspaceID, options)
}

// RemoveRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) RemoveRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceRemoveRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRemoteStateConsumers indicates an expected call of RemoveRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) RemoveRemoteStateConsumers(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).RemoveRemoteStateConsumers), ctx, workspaceID, options)
}

// RemoveTags mocks base method.
func (m *MockWorkspaces) RemoveTags(ctx context.Context, workspaceID string, options tfe.WorkspaceRemoveTagsOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTags", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTags indicates an expected call of RemoveTags.
func (mr *MockWorkspacesMockRecorder) RemoveTags(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTags", reflect.TypeOf((*MockWorkspaces)(nil).RemoveTags), ctx, workspaceID, options)
}

// RemoveVCSConnection mocks base method.
func (m *MockWorkspaces) RemoveVCSConnection(ctx context.Context, organization, workspace string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVCSConnection", ctx, organization, workspace)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVCSConnection indicates an expected call of RemoveVCSConnection.
func (mr *MockWorkspacesMockRecorder) RemoveVCSConnection(ctx, organization, workspace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVCSConnection", reflect.TypeOf((*MockWorkspaces)(nil).RemoveVCSConnection), ctx, organization, workspace)
}

// RemoveVCSConnectionByID mocks base method.
func (m *MockWorkspaces) RemoveVCSConnectionByID(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveVCSConnectionByID", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveVCSConnectionByID indicates an expected call of RemoveVCSConnectionByID.
func (mr *MockWorkspacesMockRecorder) RemoveVCSConnectionByID(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveVCSConnectionByID", reflect.TypeOf((*MockWorkspaces)(nil).RemoveVCSConnectionByID), ctx, workspaceID)
}

// Tags mocks base method.
func (m *MockWorkspaces) Tags(ctx context.Context, workspaceID string, options tfe.WorkspaceTagListOptions) (*tfe.TagList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tags", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.TagList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tags indicates an expected call of Tags.
func (mr *MockWorkspacesMockRecorder) Tags(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tags", reflect.TypeOf((*MockWorkspaces)(nil).Tags), ctx, workspaceID, options)
}

// UnassignSSHKey mocks base method.
func (m *MockWorkspaces) UnassignSSHKey(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnassignSSHKey", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnassignSSHKey indicates an expected call of UnassignSSHKey.
func (mr *MockWorkspacesMockRecorder) UnassignSSHKey(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnassignSSHKey", reflect.TypeOf((*MockWorkspaces)(nil).UnassignSSHKey), ctx, workspaceID)
}

// Unlock mocks base method.
func (m *MockWorkspaces) Unlock(ctx context.Context, workspaceID string) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", ctx, workspaceID)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unlock indicates an expected call of Unlock.
func (mr *MockWorkspacesMockRecorder) Unlock(ctx, workspaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockWorkspaces)(nil).Unlock), ctx, workspaceID)
}

// Update mocks base method.
func (m *MockWorkspaces) Update(ctx context.Context, organization, workspace string, options tfe.WorkspaceUpdateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, organization, workspace, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockWorkspacesMockRecorder) Update(ctx, organization, workspace, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWorkspaces)(nil).Update), ctx, organization, workspace, options)
}

// UpdateByID mocks base method.
func (m *MockWorkspaces) UpdateByID(ctx context.Context, workspaceID string, options tfe.WorkspaceUpdateOptions) (*tfe.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByID", ctx, workspaceID, options)
	ret0, _ := ret[0].(*tfe.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateByID indicates an expected call of UpdateByID.
func (mr *MockWorkspacesMockRecorder) UpdateByID(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByID", reflect.TypeOf((*MockWorkspaces)(nil).UpdateByID), ctx, workspaceID, options)
}

// UpdateRemoteStateConsumers mocks base method.
func (m *MockWorkspaces) UpdateRemoteStateConsumers(ctx context.Context, workspaceID string, options tfe.WorkspaceUpdateRemoteStateConsumersOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRemoteStateConsumers", ctx, workspaceID, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRemoteStateConsumers indicates an expected call of UpdateRemoteStateConsumers.
func (mr *MockWorkspacesMockRecorder) UpdateRemoteStateConsumers(ctx, workspaceID, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRemoteStateConsumers", reflect.TypeOf((*MockWorkspaces)(nil).UpdateRemoteStateConsumers), ctx, workspaceID, options)
}
