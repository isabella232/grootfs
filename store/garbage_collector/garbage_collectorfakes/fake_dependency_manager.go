// This file was generated by counterfeiter
package garbage_collectorfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/store/garbage_collector"
)

type FakeDependencyManager struct {
	DependenciesStub        func(id string) ([]string, error)
	dependenciesMutex       sync.RWMutex
	dependenciesArgsForCall []struct {
		id string
	}
	dependenciesReturns struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDependencyManager) Dependencies(id string) ([]string, error) {
	fake.dependenciesMutex.Lock()
	fake.dependenciesArgsForCall = append(fake.dependenciesArgsForCall, struct {
		id string
	}{id})
	fake.recordInvocation("Dependencies", []interface{}{id})
	fake.dependenciesMutex.Unlock()
	if fake.DependenciesStub != nil {
		return fake.DependenciesStub(id)
	} else {
		return fake.dependenciesReturns.result1, fake.dependenciesReturns.result2
	}
}

func (fake *FakeDependencyManager) DependenciesCallCount() int {
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	return len(fake.dependenciesArgsForCall)
}

func (fake *FakeDependencyManager) DependenciesArgsForCall(i int) string {
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	return fake.dependenciesArgsForCall[i].id
}

func (fake *FakeDependencyManager) DependenciesReturns(result1 []string, result2 error) {
	fake.DependenciesStub = nil
	fake.dependenciesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDependencyManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDependencyManager) recordInvocation(key string, args []interface{}) {
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

var _ garbage_collector.DependencyManager = new(FakeDependencyManager)