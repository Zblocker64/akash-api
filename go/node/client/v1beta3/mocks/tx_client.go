// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	context "context"

	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"

	v1beta3 "pkg.akt.dev/go/node/client/v1beta3"
)

// TxClient is an autogenerated mock type for the TxClient type
type TxClient struct {
	mock.Mock
}

type TxClient_Expecter struct {
	mock *mock.Mock
}

func (_m *TxClient) EXPECT() *TxClient_Expecter {
	return &TxClient_Expecter{mock: &_m.Mock}
}

// BroadcastMsgs provides a mock function with given fields: _a0, _a1, _a2
func (_m *TxClient) BroadcastMsgs(_a0 context.Context, _a1 []types.Msg, _a2 ...v1beta3.BroadcastOption) (interface{}, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BroadcastMsgs")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.Msg, ...v1beta3.BroadcastOption) (interface{}, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []types.Msg, ...v1beta3.BroadcastOption) interface{}); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []types.Msg, ...v1beta3.BroadcastOption) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxClient_BroadcastMsgs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BroadcastMsgs'
type TxClient_BroadcastMsgs_Call struct {
	*mock.Call
}

// BroadcastMsgs is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 []types.Msg
//   - _a2 ...v1beta3.BroadcastOption
func (_e *TxClient_Expecter) BroadcastMsgs(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *TxClient_BroadcastMsgs_Call {
	return &TxClient_BroadcastMsgs_Call{Call: _e.mock.On("BroadcastMsgs",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *TxClient_BroadcastMsgs_Call) Run(run func(_a0 context.Context, _a1 []types.Msg, _a2 ...v1beta3.BroadcastOption)) *TxClient_BroadcastMsgs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]v1beta3.BroadcastOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(v1beta3.BroadcastOption)
			}
		}
		run(args[0].(context.Context), args[1].([]types.Msg), variadicArgs...)
	})
	return _c
}

func (_c *TxClient_BroadcastMsgs_Call) Return(_a0 interface{}, _a1 error) *TxClient_BroadcastMsgs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxClient_BroadcastMsgs_Call) RunAndReturn(run func(context.Context, []types.Msg, ...v1beta3.BroadcastOption) (interface{}, error)) *TxClient_BroadcastMsgs_Call {
	_c.Call.Return(run)
	return _c
}

// BroadcastTx provides a mock function with given fields: _a0, _a1, _a2
func (_m *TxClient) BroadcastTx(_a0 context.Context, _a1 types.Tx, _a2 ...v1beta3.BroadcastOption) (interface{}, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for BroadcastTx")
	}

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx, ...v1beta3.BroadcastOption) (interface{}, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Tx, ...v1beta3.BroadcastOption) interface{}); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Tx, ...v1beta3.BroadcastOption) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TxClient_BroadcastTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BroadcastTx'
type TxClient_BroadcastTx_Call struct {
	*mock.Call
}

// BroadcastTx is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 types.Tx
//   - _a2 ...v1beta3.BroadcastOption
func (_e *TxClient_Expecter) BroadcastTx(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *TxClient_BroadcastTx_Call {
	return &TxClient_BroadcastTx_Call{Call: _e.mock.On("BroadcastTx",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *TxClient_BroadcastTx_Call) Run(run func(_a0 context.Context, _a1 types.Tx, _a2 ...v1beta3.BroadcastOption)) *TxClient_BroadcastTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]v1beta3.BroadcastOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(v1beta3.BroadcastOption)
			}
		}
		run(args[0].(context.Context), args[1].(types.Tx), variadicArgs...)
	})
	return _c
}

func (_c *TxClient_BroadcastTx_Call) Return(_a0 interface{}, _a1 error) *TxClient_BroadcastTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TxClient_BroadcastTx_Call) RunAndReturn(run func(context.Context, types.Tx, ...v1beta3.BroadcastOption) (interface{}, error)) *TxClient_BroadcastTx_Call {
	_c.Call.Return(run)
	return _c
}

// NewTxClient creates a new instance of TxClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTxClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *TxClient {
	mock := &TxClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
