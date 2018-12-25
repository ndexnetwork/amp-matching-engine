// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bson "github.com/globalsign/mgo/bson"
import common "github.com/ethereum/go-ethereum/common"

import mock "github.com/stretchr/testify/mock"
import types "github.com/Proofsuite/amp-matching-engine/types"

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// Create provides a mock function with given fields: token
func (_m *TokenService) Create(token *types.Token) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Token) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *TokenService) GetAll() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBaseTokens provides a mock function with given fields:
func (_m *TokenService) GetBaseTokens() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAddress provides a mock function with given fields: addr
func (_m *TokenService) GetByAddress(addr common.Address) (*types.Token, error) {
	ret := _m.Called(addr)

	var r0 *types.Token
	if rf, ok := ret.Get(0).(func(common.Address) *types.Token); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *TokenService) GetByID(id bson.ObjectId) (*types.Token, error) {
	ret := _m.Called(id)

	var r0 *types.Token
	if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Token); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQuoteTokens provides a mock function with given fields:
func (_m *TokenService) GetQuoteTokens() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
