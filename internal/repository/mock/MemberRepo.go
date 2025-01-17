// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	entity "tinder_like/internal/model/entity"

	mock "github.com/stretchr/testify/mock"

	repository "tinder_like/internal/repository"
)

// MemberRepo is an autogenerated mock type for the MemberRepo type
type MemberRepo struct {
	mock.Mock
}

// FetchMemberByGenderExceptSelf provides a mock function with given fields: ctx, gender, userID
func (_m *MemberRepo) FetchMemberByGenderExceptSelf(ctx context.Context, gender string, userID int64) ([]entity.Member, error) {
	ret := _m.Called(ctx, gender, userID)

	if len(ret) == 0 {
		panic("no return value specified for FetchMemberByGenderExceptSelf")
	}

	var r0 []entity.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) ([]entity.Member, error)); ok {
		return rf(ctx, gender, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []entity.Member); ok {
		r0 = rf(ctx, gender, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, gender, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMemberByUserID provides a mock function with given fields: ctx, userID
func (_m *MemberRepo) FindMemberByUserID(ctx context.Context, userID int64) (entity.Member, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for FindMemberByUserID")
	}

	var r0 entity.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (entity.Member, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) entity.Member); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(entity.Member)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertMember provides a mock function with given fields: ctx, tx, userID, name, gender
func (_m *MemberRepo) InsertMember(ctx context.Context, tx repository.TxProvider, userID int64, name string, gender string) error {
	ret := _m.Called(ctx, tx, userID, name, gender)

	if len(ret) == 0 {
		panic("no return value specified for InsertMember")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.TxProvider, int64, string, string) error); ok {
		r0 = rf(ctx, tx, userID, name, gender)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMemberRepo creates a new instance of MemberRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMemberRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MemberRepo {
	mock := &MemberRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
