// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"

	mock "github.com/stretchr/testify/mock"
)

// PromotionHandler is an autogenerated mock type for the PromotionHandler type
type PromotionHandler struct {
	mock.Mock
}

// CalculateDiscount provides a mock function with given fields: c
func (_m *PromotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPromotionHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewPromotionHandler creates a new instance of PromotionHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPromotionHandler(t mockConstructorTestingTNewPromotionHandler) *PromotionHandler {
	mock := &PromotionHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}