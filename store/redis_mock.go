// Code generated by MockGen. DO NOT EDIT.
// Source: store/redis/redis.go

// Package redis is a generated GoMock package.
package redis

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	v9 "github.com/redis/go-redis/v9"
)

// MockRedisClientInterface is a mock of RedisClientInterface interface.
type MockRedisClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRedisClientInterfaceMockRecorder
}

// MockRedisClientInterfaceMockRecorder is the mock recorder for MockRedisClientInterface.
type MockRedisClientInterfaceMockRecorder struct {
	mock *MockRedisClientInterface
}

// NewMockRedisClientInterface creates a new mock instance.
func NewMockRedisClientInterface(ctrl *gomock.Controller) *MockRedisClientInterface {
	mock := &MockRedisClientInterface{ctrl: ctrl}
	mock.recorder = &MockRedisClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisClientInterface) EXPECT() *MockRedisClientInterfaceMockRecorder {
	return m.recorder
}

// Del mocks base method.
func (m *MockRedisClientInterface) Del(ctx context.Context, keys ...string) *v9.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Del", varargs...)
	ret0, _ := ret[0].(*v9.IntCmd)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockRedisClientInterfaceMockRecorder) Del(ctx interface{}, keys ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockRedisClientInterface)(nil).Del), varargs...)
}

// Expire mocks base method.
func (m *MockRedisClientInterface) Expire(ctx context.Context, key string, expiration time.Duration) *v9.BoolCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expire", ctx, key, expiration)
	ret0, _ := ret[0].(*v9.BoolCmd)
	return ret0
}

// Expire indicates an expected call of Expire.
func (mr *MockRedisClientInterfaceMockRecorder) Expire(ctx, key, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expire", reflect.TypeOf((*MockRedisClientInterface)(nil).Expire), ctx, key, expiration)
}

// FlushAll mocks base method.
func (m *MockRedisClientInterface) FlushAll(ctx context.Context) *v9.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushAll", ctx)
	ret0, _ := ret[0].(*v9.StatusCmd)
	return ret0
}

// FlushAll indicates an expected call of FlushAll.
func (mr *MockRedisClientInterfaceMockRecorder) FlushAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushAll", reflect.TypeOf((*MockRedisClientInterface)(nil).FlushAll), ctx)
}

// Get mocks base method.
func (m *MockRedisClientInterface) Get(ctx context.Context, key string) *v9.StringCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key)
	ret0, _ := ret[0].(*v9.StringCmd)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisClientInterfaceMockRecorder) Get(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisClientInterface)(nil).Get), ctx, key)
}

// SAdd mocks base method.
func (m *MockRedisClientInterface) SAdd(ctx context.Context, key string, members ...any) *v9.IntCmd {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, key}
	for _, a := range members {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SAdd", varargs...)
	ret0, _ := ret[0].(*v9.IntCmd)
	return ret0
}

// SAdd indicates an expected call of SAdd.
func (mr *MockRedisClientInterfaceMockRecorder) SAdd(ctx, key interface{}, members ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, key}, members...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SAdd", reflect.TypeOf((*MockRedisClientInterface)(nil).SAdd), varargs...)
}

// SMembers mocks base method.
func (m *MockRedisClientInterface) SMembers(ctx context.Context, key string) *v9.StringSliceCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SMembers", ctx, key)
	ret0, _ := ret[0].(*v9.StringSliceCmd)
	return ret0
}

// SMembers indicates an expected call of SMembers.
func (mr *MockRedisClientInterfaceMockRecorder) SMembers(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SMembers", reflect.TypeOf((*MockRedisClientInterface)(nil).SMembers), ctx, key)
}

// Set mocks base method.
func (m *MockRedisClientInterface) Set(ctx context.Context, key string, values any, expiration time.Duration) *v9.StatusCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, values, expiration)
	ret0, _ := ret[0].(*v9.StatusCmd)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisClientInterfaceMockRecorder) Set(ctx, key, values, expiration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisClientInterface)(nil).Set), ctx, key, values, expiration)
}

// TTL mocks base method.
func (m *MockRedisClientInterface) TTL(ctx context.Context, key string) *v9.DurationCmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TTL", ctx, key)
	ret0, _ := ret[0].(*v9.DurationCmd)
	return ret0
}

// TTL indicates an expected call of TTL.
func (mr *MockRedisClientInterfaceMockRecorder) TTL(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TTL", reflect.TypeOf((*MockRedisClientInterface)(nil).TTL), ctx, key)
}
