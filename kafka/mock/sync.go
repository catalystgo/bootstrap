// Code generated by MockGen. DO NOT EDIT.
// Source: kafka/sync.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	kafka "github.com/catalystgo/bootstrap/kafka"
	gomock "github.com/golang/mock/gomock"
)

// MockSyncProducer is a mock of SyncProducer interface.
type MockSyncProducer struct {
	ctrl     *gomock.Controller
	recorder *MockSyncProducerMockRecorder
}

// MockSyncProducerMockRecorder is the mock recorder for MockSyncProducer.
type MockSyncProducerMockRecorder struct {
	mock *MockSyncProducer
}

// NewMockSyncProducer creates a new mock instance.
func NewMockSyncProducer(ctrl *gomock.Controller) *MockSyncProducer {
	mock := &MockSyncProducer{ctrl: ctrl}
	mock.recorder = &MockSyncProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncProducer) EXPECT() *MockSyncProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSyncProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSyncProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncProducer)(nil).Close))
}

// Produce mocks base method.
func (m *MockSyncProducer) Produce(topic string, message []byte, opts ...kafka.MessageOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{topic, message}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Produce", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Produce indicates an expected call of Produce.
func (mr *MockSyncProducerMockRecorder) Produce(topic, message interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{topic, message}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockSyncProducer)(nil).Produce), varargs...)
}

// ProduceWithContext mocks base method.
func (m *MockSyncProducer) ProduceWithContext(ctx context.Context, topic string, message []byte, opts ...kafka.MessageOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, topic, message}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProduceWithContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceWithContext indicates an expected call of ProduceWithContext.
func (mr *MockSyncProducerMockRecorder) ProduceWithContext(ctx, topic, message interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, topic, message}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceWithContext", reflect.TypeOf((*MockSyncProducer)(nil).ProduceWithContext), varargs...)
}