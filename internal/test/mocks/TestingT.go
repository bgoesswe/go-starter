// Initial code generated by mockery using a local installation.
// The methods for the testing.T interface will not change that often so no
// automatic generation of the mock is required.

package mocks

import mock "github.com/stretchr/testify/mock"

// TestingT is an autogenerated mock type for the TestingT type
type TestingT struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: _a0
func (_m *TestingT) Cleanup(_a0 func()) {
	_m.Called(_a0)
}

// Error provides a mock function with given fields: args
func (_m *TestingT) Error(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Errorf provides a mock function with given fields: format, args
func (_m *TestingT) Errorf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fail provides a mock function with given fields:
func (_m *TestingT) Fail() {
	_m.Called()
}

// FailNow provides a mock function with given fields:
func (_m *TestingT) FailNow() {
	_m.Called()
}

// Failed provides a mock function with given fields:
func (_m *TestingT) Failed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Fatal provides a mock function with given fields: args
func (_m *TestingT) Fatal(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatalf provides a mock function with given fields: format, args
func (_m *TestingT) Fatalf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Helper provides a mock function with given fields:
func (_m *TestingT) Helper() {
	_m.Called()
}

// Log provides a mock function with given fields: args
func (_m *TestingT) Log(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Logf provides a mock function with given fields: format, args
func (_m *TestingT) Logf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Name provides a mock function with given fields:
func (_m *TestingT) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Skip provides a mock function with given fields: args
func (_m *TestingT) Skip(args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// SkipNow provides a mock function with given fields:
func (_m *TestingT) SkipNow() {
	_m.Called()
}

// Skipf provides a mock function with given fields: format, args
func (_m *TestingT) Skipf(format string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Skipped provides a mock function with given fields:
func (_m *TestingT) Skipped() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// TempDir provides a mock function with given fields:
func (_m *TestingT) TempDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
