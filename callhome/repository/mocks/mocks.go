package mocks

import (
	"context"

	"github.com/mainflux/callhome/callhome"
	"github.com/stretchr/testify/mock"
)

var _ callhome.TelemetryRepo = (*mockRepo)(nil)

type mockRepo struct {
	mock.Mock
}

func (mr *mockRepo) RetrieveAll(ctx context.Context, pm callhome.PageMetadata) (callhome.TelemetryPage, error) {
	ret := mr.Called(ctx, pm)
	return ret.Get(0).(callhome.TelemetryPage), ret.Error(1)
}

func (mr *mockRepo) RetrieveByIP(ctx context.Context, email string) (callhome.Telemetry, error) {
	ret := mr.Called(ctx, email)
	return ret.Get(0).(callhome.Telemetry), ret.Error(1)
}

func (mr *mockRepo) Save(ctx context.Context, t callhome.Telemetry) error {
	ret := mr.Called(ctx, t)
	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, callhome.Telemetry) error); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (mr *mockRepo) Update(ctx context.Context, u callhome.Telemetry) error {
	ret := mr.Called(ctx, u)
	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, callhome.Telemetry) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewTelemetryRepo interface {
	mock.TestingT
	Cleanup(func())
}

func NewTelemetryRepo(t mockConstructorTestingTNewTelemetryRepo) *mockRepo {
	mock := &mockRepo{}
	mock.Mock.Test(t)

	return mock
}