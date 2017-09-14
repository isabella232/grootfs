// This file was generated by counterfeiter
package grootfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/grootfs/groot"
	"code.cloudfoundry.org/lager"
)

type FakeMetricsEmitter struct {
	TryEmitUsageStub        func(logger lager.Logger, name string, usage int64, units string)
	tryEmitUsageMutex       sync.RWMutex
	tryEmitUsageArgsForCall []struct {
		logger lager.Logger
		name   string
		usage  int64
		units  string
	}
	TryEmitDurationFromStub        func(logger lager.Logger, name string, from time.Time)
	tryEmitDurationFromMutex       sync.RWMutex
	tryEmitDurationFromArgsForCall []struct {
		logger lager.Logger
		name   string
		from   time.Time
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMetricsEmitter) TryEmitUsage(logger lager.Logger, name string, usage int64, units string) {
	fake.tryEmitUsageMutex.Lock()
	fake.tryEmitUsageArgsForCall = append(fake.tryEmitUsageArgsForCall, struct {
		logger lager.Logger
		name   string
		usage  int64
		units  string
	}{logger, name, usage, units})
	fake.recordInvocation("TryEmitUsage", []interface{}{logger, name, usage, units})
	fake.tryEmitUsageMutex.Unlock()
	if fake.TryEmitUsageStub != nil {
		fake.TryEmitUsageStub(logger, name, usage, units)
	}
}

func (fake *FakeMetricsEmitter) TryEmitUsageCallCount() int {
	fake.tryEmitUsageMutex.RLock()
	defer fake.tryEmitUsageMutex.RUnlock()
	return len(fake.tryEmitUsageArgsForCall)
}

func (fake *FakeMetricsEmitter) TryEmitUsageArgsForCall(i int) (lager.Logger, string, int64, string) {
	fake.tryEmitUsageMutex.RLock()
	defer fake.tryEmitUsageMutex.RUnlock()
	return fake.tryEmitUsageArgsForCall[i].logger, fake.tryEmitUsageArgsForCall[i].name, fake.tryEmitUsageArgsForCall[i].usage, fake.tryEmitUsageArgsForCall[i].units
}

func (fake *FakeMetricsEmitter) TryEmitDurationFrom(logger lager.Logger, name string, from time.Time) {
	fake.tryEmitDurationFromMutex.Lock()
	fake.tryEmitDurationFromArgsForCall = append(fake.tryEmitDurationFromArgsForCall, struct {
		logger lager.Logger
		name   string
		from   time.Time
	}{logger, name, from})
	fake.recordInvocation("TryEmitDurationFrom", []interface{}{logger, name, from})
	fake.tryEmitDurationFromMutex.Unlock()
	if fake.TryEmitDurationFromStub != nil {
		fake.TryEmitDurationFromStub(logger, name, from)
	}
}

func (fake *FakeMetricsEmitter) TryEmitDurationFromCallCount() int {
	fake.tryEmitDurationFromMutex.RLock()
	defer fake.tryEmitDurationFromMutex.RUnlock()
	return len(fake.tryEmitDurationFromArgsForCall)
}

func (fake *FakeMetricsEmitter) TryEmitDurationFromArgsForCall(i int) (lager.Logger, string, time.Time) {
	fake.tryEmitDurationFromMutex.RLock()
	defer fake.tryEmitDurationFromMutex.RUnlock()
	return fake.tryEmitDurationFromArgsForCall[i].logger, fake.tryEmitDurationFromArgsForCall[i].name, fake.tryEmitDurationFromArgsForCall[i].from
}

func (fake *FakeMetricsEmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.tryEmitUsageMutex.RLock()
	defer fake.tryEmitUsageMutex.RUnlock()
	fake.tryEmitDurationFromMutex.RLock()
	defer fake.tryEmitDurationFromMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeMetricsEmitter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ groot.MetricsEmitter = new(FakeMetricsEmitter)
