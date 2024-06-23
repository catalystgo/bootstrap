// Code generated by MockGen. DO NOT EDIT.
// Source: kafka/async.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	kafka "github.com/catalystgo/bootstrap/kafka"
	gomock "github.com/golang/mock/gomock"
)

// MockAsyncProducer is a mock of AsyncProducer interface.
type MockAsyncProducer struct {
	ctrl     *gomock.Controller
	recorder *MockAsyncProducerMockRecorder
}

// MockAsyncProducerMockRecorder is the mock recorder for MockAsyncProducer.
type MockAsyncProducerMockRecorder struct {
	mock *MockAsyncProducer
}

// NewMockAsyncProducer creates a new mock instance.
func NewMockAsyncProducer(ctrl *gomock.Controller) *MockAsyncProducer {
	mock := &MockAsyncProducer{ctrl: ctrl}
	mock.recorder = &MockAsyncProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAsyncProducer) EXPECT() *MockAsyncProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockAsyncProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAsyncProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAsyncProducer)(nil).Close))
}

// Produce mocks base method.
func (m *MockAsyncProducer) Produce(topic string, message []byte, opts ...kafka.MessageOption) {
	m.ctrl.T.Helper()
	varargs := []interface{}{topic, message}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Produce", varargs...)
}

// Produce indicates an expected call of Produce.
func (mr *MockAsyncProducerMockRecorder) Produce(topic, message interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{topic, message}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockAsyncProducer)(nil).Produce), varargs...)
}

// ProduceWithContext mocks base method.
func (m *MockAsyncProducer) ProduceWithContext(ctx context.Context, topic string, message []byte, opts ...kafka.MessageOption) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, topic, message}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "ProduceWithContext", varargs...)
}

// ProduceWithContext indicates an expected call of ProduceWithContext.
func (mr *MockAsyncProducerMockRecorder) ProduceWithContext(ctx, topic, message interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, topic, message}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceWithContext", reflect.TypeOf((*MockAsyncProducer)(nil).ProduceWithContext), varargs...)
}
