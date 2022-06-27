// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	netlink "github.com/vishvananda/netlink"
)

// NetlinkProvider is an autogenerated mock type for the NetlinkProvider type
type NetlinkProvider struct {
	mock.Mock
}

// GetDevLinkDeviceEswitchAttrs provides a mock function with given fields: ifName
func (_m *NetlinkProvider) GetDevLinkDeviceEswitchAttrs(ifName string) (*netlink.DevlinkDevEswitchAttr, error) {
	ret := _m.Called(ifName)

	var r0 *netlink.DevlinkDevEswitchAttr
	if rf, ok := ret.Get(0).(func(string) *netlink.DevlinkDevEswitchAttr); ok {
		r0 = rf(ifName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*netlink.DevlinkDevEswitchAttr)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ifName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLinkAttrs provides a mock function with given fields: ifName
func (_m *NetlinkProvider) GetLinkAttrs(ifName string) (*netlink.LinkAttrs, error) {
	ret := _m.Called(ifName)

	var r0 *netlink.LinkAttrs
	if rf, ok := ret.Get(0).(func(string) *netlink.LinkAttrs); ok {
		r0 = rf(ifName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*netlink.LinkAttrs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ifName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewNetlinkProviderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewNetlinkProvider creates a new instance of NetlinkProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewNetlinkProvider(t NewNetlinkProviderT) *NetlinkProvider {
	mock := &NetlinkProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}