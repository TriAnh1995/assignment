// Code generated by mockery v2.20.0. DO NOT EDIT.

package controller

import (
	model "assignment/internal/model"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockController is an autogenerated mock type for the Controller type
type MockController struct {
	mock.Mock
}

// AddFriends provides a mock function with given fields: _a0, _a1
func (_m *MockController) AddFriends(_a0 context.Context, _a1 []string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddUsers provides a mock function with given fields: _a0, _a1
func (_m *MockController) AddUsers(_a0 context.Context, _a1 model.User) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CommonFriends provides a mock function with given fields: _a0, _a1
func (_m *MockController) CommonFriends(_a0 context.Context, _a1 []string) (model.FriendshipInfo, error) {
	ret := _m.Called(_a0, _a1)

	var r0 model.FriendshipInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) (model.FriendshipInfo, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) model.FriendshipInfo); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.FriendshipInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FriendsList provides a mock function with given fields: _a0, _a1
func (_m *MockController) FriendsList(_a0 context.Context, _a1 string) (model.FriendshipInfo, error) {
	ret := _m.Called(_a0, _a1)

	var r0 model.FriendshipInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.FriendshipInfo, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.FriendshipInfo); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.FriendshipInfo)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTopic provides a mock function with given fields: _a0, _a1
func (_m *MockController) UpdateTopic(_a0 context.Context, _a1 model.UpdateInfo) ([]string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateInfo) ([]string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.UpdateInfo) []string); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.UpdateInfo) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockController interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockController creates a new instance of MockController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockController(t mockConstructorTestingTNewMockController) *MockController {
	mock := &MockController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
