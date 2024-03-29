// Code generated by MockGen. DO NOT EDIT.
// Source: familytree/familytree.go
//
// Generated by this command:
//
//	mockgen -source=familytree/familytree.go -destination=familytree/mock/familytree.go
//

// Package mock_familytree is a generated GoMock package.
package mock_familytree

import (
	context "context"
	reflect "reflect"

	entity "github.com/GeovaneCavalcante/tree-genealogical/internal/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockGenealogyInterface is a mock of GenealogyInterface interface.
type MockGenealogyInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGenealogyInterfaceMockRecorder
}

// MockGenealogyInterfaceMockRecorder is the mock recorder for MockGenealogyInterface.
type MockGenealogyInterfaceMockRecorder struct {
	mock *MockGenealogyInterface
}

// NewMockGenealogyInterface creates a new mock instance.
func NewMockGenealogyInterface(ctrl *gomock.Controller) *MockGenealogyInterface {
	mock := &MockGenealogyInterface{ctrl: ctrl}
	mock.recorder = &MockGenealogyInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenealogyInterface) EXPECT() *MockGenealogyInterfaceMockRecorder {
	return m.recorder
}

// BuildFamilyTree mocks base method.
func (m *MockGenealogyInterface) BuildFamilyTree(ctx context.Context, rootPerson *entity.Person, persons []*entity.Person, level int) []*entity.Relative {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildFamilyTree", ctx, rootPerson, persons, level)
	ret0, _ := ret[0].([]*entity.Relative)
	return ret0
}

// BuildFamilyTree indicates an expected call of BuildFamilyTree.
func (mr *MockGenealogyInterfaceMockRecorder) BuildFamilyTree(ctx, rootPerson, persons, level any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildFamilyTree", reflect.TypeOf((*MockGenealogyInterface)(nil).BuildFamilyTree), ctx, rootPerson, persons, level)
}

// GetRelatives mocks base method.
func (m *MockGenealogyInterface) GetRelatives(ctx context.Context) []*entity.Relative {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRelatives", ctx)
	ret0, _ := ret[0].([]*entity.Relative)
	return ret0
}

// GetRelatives indicates an expected call of GetRelatives.
func (mr *MockGenealogyInterfaceMockRecorder) GetRelatives(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelatives", reflect.TypeOf((*MockGenealogyInterface)(nil).GetRelatives), ctx)
}

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// CalculateKinshipDistance mocks base method.
func (m *MockUseCase) CalculateKinshipDistance(ctx context.Context, firstPersonName, secondPersonName string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateKinshipDistance", ctx, firstPersonName, secondPersonName)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateKinshipDistance indicates an expected call of CalculateKinshipDistance.
func (mr *MockUseCaseMockRecorder) CalculateKinshipDistance(ctx, firstPersonName, secondPersonName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateKinshipDistance", reflect.TypeOf((*MockUseCase)(nil).CalculateKinshipDistance), ctx, firstPersonName, secondPersonName)
}

// DetermineRelationship mocks base method.
func (m *MockUseCase) DetermineRelationship(ctx context.Context, firstPersonName, secondPersonName string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetermineRelationship", ctx, firstPersonName, secondPersonName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetermineRelationship indicates an expected call of DetermineRelationship.
func (mr *MockUseCaseMockRecorder) DetermineRelationship(ctx, firstPersonName, secondPersonName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetermineRelationship", reflect.TypeOf((*MockUseCase)(nil).DetermineRelationship), ctx, firstPersonName, secondPersonName)
}

// GetAllFamilyMembers mocks base method.
func (m *MockUseCase) GetAllFamilyMembers(ctx context.Context, personName string) ([]*entity.Relative, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFamilyMembers", ctx, personName)
	ret0, _ := ret[0].([]*entity.Relative)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFamilyMembers indicates an expected call of GetAllFamilyMembers.
func (mr *MockUseCaseMockRecorder) GetAllFamilyMembers(ctx, personName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFamilyMembers", reflect.TypeOf((*MockUseCase)(nil).GetAllFamilyMembers), ctx, personName)
}
