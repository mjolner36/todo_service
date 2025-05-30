// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	context "context"
	model "todo_service/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: ctx, task
func (_m *Repository) CreateTask(ctx context.Context, task *model.Task) (int, error) {
	ret := _m.Called(ctx, task)

	if len(ret) == 0 {
		panic("no return value specified for CreateTask")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Task) (int, error)); ok {
		return rf(ctx, task)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Task) int); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Task) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
