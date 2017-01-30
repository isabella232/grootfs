// This file was generated by counterfeiter
package base_image_pullerfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/base_image_puller"
	"code.cloudfoundry.org/lager"
)

type FakeVolumeDriver struct {
	PathStub        func(logger lager.Logger, id string) (string, error)
	pathMutex       sync.RWMutex
	pathArgsForCall []struct {
		logger lager.Logger
		id     string
	}
	pathReturns struct {
		result1 string
		result2 error
	}
	CreateStub func(logger lager.Logger, parentID, id string) (string, error)

	createMutex       sync.RWMutex
	createArgsForCall []struct {
		logger   lager.Logger
		parentID string
		id       string
	}
	createReturns struct {
		result1 string
		result2 error
	}
	DestroyVolumeStub        func(logger lager.Logger, id string) error
	destroyVolumeMutex       sync.RWMutex
	destroyVolumeArgsForCall []struct {
		logger lager.Logger
		id     string
	}
	destroyVolumeReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeDriver) Path(logger lager.Logger, id string) (string, error) {
	fake.pathMutex.Lock()
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct {
		logger lager.Logger
		id     string
	}{logger, id})
	fake.recordInvocation("Path", []interface{}{logger, id})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub(logger, id)
	} else {
		return fake.pathReturns.result1, fake.pathReturns.result2
	}
}

func (fake *FakeVolumeDriver) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *FakeVolumeDriver) PathArgsForCall(i int) (lager.Logger, string) {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return fake.pathArgsForCall[i].logger, fake.pathArgsForCall[i].id
}

func (fake *FakeVolumeDriver) PathReturns(result1 string, result2 error) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeDriver) Create(logger lager.Logger, parentID string, id string) (string, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		logger   lager.Logger
		parentID string
		id       string
	}{logger, parentID, id})
	fake.recordInvocation("Create", []interface{}{logger, parentID, id})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(logger, parentID, id)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeVolumeDriver) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeVolumeDriver) CreateArgsForCall(i int) (lager.Logger, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].logger, fake.createArgsForCall[i].parentID, fake.createArgsForCall[i].id
}

func (fake *FakeVolumeDriver) CreateReturns(result1 string, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeDriver) DestroyVolume(logger lager.Logger, id string) error {
	fake.destroyVolumeMutex.Lock()
	fake.destroyVolumeArgsForCall = append(fake.destroyVolumeArgsForCall, struct {
		logger lager.Logger
		id     string
	}{logger, id})
	fake.recordInvocation("DestroyVolume", []interface{}{logger, id})
	fake.destroyVolumeMutex.Unlock()
	if fake.DestroyVolumeStub != nil {
		return fake.DestroyVolumeStub(logger, id)
	} else {
		return fake.destroyVolumeReturns.result1
	}
}

func (fake *FakeVolumeDriver) DestroyVolumeCallCount() int {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return len(fake.destroyVolumeArgsForCall)
}

func (fake *FakeVolumeDriver) DestroyVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return fake.destroyVolumeArgsForCall[i].logger, fake.destroyVolumeArgsForCall[i].id
}

func (fake *FakeVolumeDriver) DestroyVolumeReturns(result1 error) {
	fake.DestroyVolumeStub = nil
	fake.destroyVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVolumeDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeVolumeDriver) recordInvocation(key string, args []interface{}) {
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

var _ base_image_puller.VolumeDriver = new(FakeVolumeDriver)
