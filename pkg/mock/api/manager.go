// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeflow/katib/pkg/api (interfaces: ManagerClient)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	api "github.com/kubeflow/katib/pkg/api"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockManagerClient is a mock of ManagerClient interface
type MockManagerClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagerClientMockRecorder
}

// MockManagerClientMockRecorder is the mock recorder for MockManagerClient
type MockManagerClientMockRecorder struct {
	mock *MockManagerClient
}

// NewMockManagerClient creates a new mock instance
func NewMockManagerClient(ctrl *gomock.Controller) *MockManagerClient {
	mock := &MockManagerClient{ctrl: ctrl}
	mock.recorder = &MockManagerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManagerClient) EXPECT() *MockManagerClientMockRecorder {
	return m.recorder
}

// CreateStudy mocks base method
func (m *MockManagerClient) CreateStudy(arg0 context.Context, arg1 *api.CreateStudyRequest, arg2 ...grpc.CallOption) (*api.CreateStudyReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateStudy", varargs...)
	ret0, _ := ret[0].(*api.CreateStudyReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudy indicates an expected call of CreateStudy
func (mr *MockManagerClientMockRecorder) CreateStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudy", reflect.TypeOf((*MockManagerClient)(nil).CreateStudy), varargs...)
}

// CreateTrial mocks base method
func (m *MockManagerClient) CreateTrial(arg0 context.Context, arg1 *api.CreateTrialRequest, arg2 ...grpc.CallOption) (*api.CreateTrialReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTrial", varargs...)
	ret0, _ := ret[0].(*api.CreateTrialReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrial indicates an expected call of CreateTrial
func (mr *MockManagerClientMockRecorder) CreateTrial(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrial", reflect.TypeOf((*MockManagerClient)(nil).CreateTrial), varargs...)
}

// CreateWorker mocks base method
func (m *MockManagerClient) CreateWorker(arg0 context.Context, arg1 *api.CreateWorkerReauest, arg2 ...grpc.CallOption) (*api.CreateWorkerReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateWorker", varargs...)
	ret0, _ := ret[0].(*api.CreateWorkerReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorker indicates an expected call of CreateWorker
func (mr *MockManagerClientMockRecorder) CreateWorker(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorker", reflect.TypeOf((*MockManagerClient)(nil).CreateWorker), varargs...)
}

// GetEarlyStoppingParameterList mocks base method
func (m *MockManagerClient) GetEarlyStoppingParameterList(arg0 context.Context, arg1 *api.GetEarlyStoppingParameterListRequest, arg2 ...grpc.CallOption) (*api.GetEarlyStoppingParameterListReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEarlyStoppingParameterList", varargs...)
	ret0, _ := ret[0].(*api.GetEarlyStoppingParameterListReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEarlyStoppingParameterList indicates an expected call of GetEarlyStoppingParameterList
func (mr *MockManagerClientMockRecorder) GetEarlyStoppingParameterList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEarlyStoppingParameterList", reflect.TypeOf((*MockManagerClient)(nil).GetEarlyStoppingParameterList), varargs...)
}

// GetEarlyStoppingParameters mocks base method
func (m *MockManagerClient) GetEarlyStoppingParameters(arg0 context.Context, arg1 *api.GetEarlyStoppingParametersRequest, arg2 ...grpc.CallOption) (*api.GetEarlyStoppingParametersReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEarlyStoppingParameters", varargs...)
	ret0, _ := ret[0].(*api.GetEarlyStoppingParametersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEarlyStoppingParameters indicates an expected call of GetEarlyStoppingParameters
func (mr *MockManagerClientMockRecorder) GetEarlyStoppingParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEarlyStoppingParameters", reflect.TypeOf((*MockManagerClient)(nil).GetEarlyStoppingParameters), varargs...)
}

// GetMetrics mocks base method
func (m *MockManagerClient) GetMetrics(arg0 context.Context, arg1 *api.GetMetricsRequest, arg2 ...grpc.CallOption) (*api.GetMetricsReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMetrics", varargs...)
	ret0, _ := ret[0].(*api.GetMetricsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetrics indicates an expected call of GetMetrics
func (mr *MockManagerClientMockRecorder) GetMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetrics", reflect.TypeOf((*MockManagerClient)(nil).GetMetrics), varargs...)
}

// GetSavedModels mocks base method
func (m *MockManagerClient) GetSavedModels(arg0 context.Context, arg1 *api.GetSavedModelsRequest, arg2 ...grpc.CallOption) (*api.GetSavedModelsReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSavedModels", varargs...)
	ret0, _ := ret[0].(*api.GetSavedModelsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavedModels indicates an expected call of GetSavedModels
func (mr *MockManagerClientMockRecorder) GetSavedModels(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavedModels", reflect.TypeOf((*MockManagerClient)(nil).GetSavedModels), varargs...)
}

// GetSavedStudies mocks base method
func (m *MockManagerClient) GetSavedStudies(arg0 context.Context, arg1 *api.GetSavedStudiesRequest, arg2 ...grpc.CallOption) (*api.GetSavedStudiesReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSavedStudies", varargs...)
	ret0, _ := ret[0].(*api.GetSavedStudiesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavedStudies indicates an expected call of GetSavedStudies
func (mr *MockManagerClientMockRecorder) GetSavedStudies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavedStudies", reflect.TypeOf((*MockManagerClient)(nil).GetSavedStudies), varargs...)
}

// GetShouldStopWorkers mocks base method
func (m *MockManagerClient) GetShouldStopWorkers(arg0 context.Context, arg1 *api.GetShouldStopWorkersRequest, arg2 ...grpc.CallOption) (*api.GetShouldStopWorkersReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetShouldStopWorkers", varargs...)
	ret0, _ := ret[0].(*api.GetShouldStopWorkersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShouldStopWorkers indicates an expected call of GetShouldStopWorkers
func (mr *MockManagerClientMockRecorder) GetShouldStopWorkers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShouldStopWorkers", reflect.TypeOf((*MockManagerClient)(nil).GetShouldStopWorkers), varargs...)
}

// GetStudy mocks base method
func (m *MockManagerClient) GetStudy(arg0 context.Context, arg1 *api.GetStudyRequest, arg2 ...grpc.CallOption) (*api.GetStudyReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudy", varargs...)
	ret0, _ := ret[0].(*api.GetStudyReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudy indicates an expected call of GetStudy
func (mr *MockManagerClientMockRecorder) GetStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudy", reflect.TypeOf((*MockManagerClient)(nil).GetStudy), varargs...)
}

// GetStudyList mocks base method
func (m *MockManagerClient) GetStudyList(arg0 context.Context, arg1 *api.GetStudyListRequest, arg2 ...grpc.CallOption) (*api.GetStudyListReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetStudyList", varargs...)
	ret0, _ := ret[0].(*api.GetStudyListReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudyList indicates an expected call of GetStudyList
func (mr *MockManagerClientMockRecorder) GetStudyList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudyList", reflect.TypeOf((*MockManagerClient)(nil).GetStudyList), varargs...)
}

// GetSuggestionParameterList mocks base method
func (m *MockManagerClient) GetSuggestionParameterList(arg0 context.Context, arg1 *api.GetSuggestionParameterListRequest, arg2 ...grpc.CallOption) (*api.GetSuggestionParameterListReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSuggestionParameterList", varargs...)
	ret0, _ := ret[0].(*api.GetSuggestionParameterListReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuggestionParameterList indicates an expected call of GetSuggestionParameterList
func (mr *MockManagerClientMockRecorder) GetSuggestionParameterList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuggestionParameterList", reflect.TypeOf((*MockManagerClient)(nil).GetSuggestionParameterList), varargs...)
}

// GetSuggestionParameters mocks base method
func (m *MockManagerClient) GetSuggestionParameters(arg0 context.Context, arg1 *api.GetSuggestionParametersRequest, arg2 ...grpc.CallOption) (*api.GetSuggestionParametersReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSuggestionParameters", varargs...)
	ret0, _ := ret[0].(*api.GetSuggestionParametersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuggestionParameters indicates an expected call of GetSuggestionParameters
func (mr *MockManagerClientMockRecorder) GetSuggestionParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuggestionParameters", reflect.TypeOf((*MockManagerClient)(nil).GetSuggestionParameters), varargs...)
}

// GetSuggestions mocks base method
func (m *MockManagerClient) GetSuggestions(arg0 context.Context, arg1 *api.GetSuggestionsRequest, arg2 ...grpc.CallOption) (*api.GetSuggestionsReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSuggestions", varargs...)
	ret0, _ := ret[0].(*api.GetSuggestionsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSuggestions indicates an expected call of GetSuggestions
func (mr *MockManagerClientMockRecorder) GetSuggestions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuggestions", reflect.TypeOf((*MockManagerClient)(nil).GetSuggestions), varargs...)
}

// GetTrials mocks base method
func (m *MockManagerClient) GetTrials(arg0 context.Context, arg1 *api.GetTrialsRequest, arg2 ...grpc.CallOption) (*api.GetTrialsReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTrials", varargs...)
	ret0, _ := ret[0].(*api.GetTrialsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrials indicates an expected call of GetTrials
func (mr *MockManagerClientMockRecorder) GetTrials(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrials", reflect.TypeOf((*MockManagerClient)(nil).GetTrials), varargs...)
}

// GetWorkers mocks base method
func (m *MockManagerClient) GetWorkers(arg0 context.Context, arg1 *api.GetWorkersRequest, arg2 ...grpc.CallOption) (*api.GetWorkersReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetWorkers", varargs...)
	ret0, _ := ret[0].(*api.GetWorkersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWorkers indicates an expected call of GetWorkers
func (mr *MockManagerClientMockRecorder) GetWorkers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkers", reflect.TypeOf((*MockManagerClient)(nil).GetWorkers), varargs...)
}

// ReportMetrics mocks base method
func (m *MockManagerClient) ReportMetrics(arg0 context.Context, arg1 *api.ReportMetricsRequest, arg2 ...grpc.CallOption) (*api.ReportMetricsReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReportMetrics", varargs...)
	ret0, _ := ret[0].(*api.ReportMetricsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportMetrics indicates an expected call of ReportMetrics
func (mr *MockManagerClientMockRecorder) ReportMetrics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportMetrics", reflect.TypeOf((*MockManagerClient)(nil).ReportMetrics), varargs...)
}

// SaveModel mocks base method
func (m *MockManagerClient) SaveModel(arg0 context.Context, arg1 *api.SaveModelRequest, arg2 ...grpc.CallOption) (*api.SaveModelReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveModel", varargs...)
	ret0, _ := ret[0].(*api.SaveModelReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveModel indicates an expected call of SaveModel
func (mr *MockManagerClientMockRecorder) SaveModel(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveModel", reflect.TypeOf((*MockManagerClient)(nil).SaveModel), varargs...)
}

// SaveStudy mocks base method
func (m *MockManagerClient) SaveStudy(arg0 context.Context, arg1 *api.SaveStudyRequest, arg2 ...grpc.CallOption) (*api.SaveStudyReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SaveStudy", varargs...)
	ret0, _ := ret[0].(*api.SaveStudyReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveStudy indicates an expected call of SaveStudy
func (mr *MockManagerClientMockRecorder) SaveStudy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveStudy", reflect.TypeOf((*MockManagerClient)(nil).SaveStudy), varargs...)
}

// SetEarlyStoppingParameters mocks base method
func (m *MockManagerClient) SetEarlyStoppingParameters(arg0 context.Context, arg1 *api.SetEarlyStoppingParametersRequest, arg2 ...grpc.CallOption) (*api.SetEarlyStoppingParametersReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetEarlyStoppingParameters", varargs...)
	ret0, _ := ret[0].(*api.SetEarlyStoppingParametersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetEarlyStoppingParameters indicates an expected call of SetEarlyStoppingParameters
func (mr *MockManagerClientMockRecorder) SetEarlyStoppingParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEarlyStoppingParameters", reflect.TypeOf((*MockManagerClient)(nil).SetEarlyStoppingParameters), varargs...)
}

// SetSuggestionParameters mocks base method
func (m *MockManagerClient) SetSuggestionParameters(arg0 context.Context, arg1 *api.SetSuggestionParametersRequest, arg2 ...grpc.CallOption) (*api.SetSuggestionParametersReply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetSuggestionParameters", varargs...)
	ret0, _ := ret[0].(*api.SetSuggestionParametersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSuggestionParameters indicates an expected call of SetSuggestionParameters
func (mr *MockManagerClientMockRecorder) SetSuggestionParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSuggestionParameters", reflect.TypeOf((*MockManagerClient)(nil).SetSuggestionParameters), varargs...)
}
