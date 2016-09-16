// This file was generated by counterfeiter
package remotefakes

import (
	"net/url"
	"sync"

	"code.cloudfoundry.org/grootfs/fetcher/remote"
	"code.cloudfoundry.org/lager"
	specsv1 "github.com/opencontainers/image-spec/specs-go/v1"
)

type FakeSource struct {
	ManifestStub        func(logger lager.Logger, imageURL *url.URL) (remote.Manifest, error)
	manifestMutex       sync.RWMutex
	manifestArgsForCall []struct {
		logger   lager.Logger
		imageURL *url.URL
	}
	manifestReturns struct {
		result1 remote.Manifest
		result2 error
	}
	ConfigStub        func(logger lager.Logger, imageURL *url.URL, manifest remote.Manifest) (specsv1.Image, error)
	configMutex       sync.RWMutex
	configArgsForCall []struct {
		logger   lager.Logger
		imageURL *url.URL
		manifest remote.Manifest
	}
	configReturns struct {
		result1 specsv1.Image
		result2 error
	}
	BlobStub        func(logger lager.Logger, imageURL *url.URL, digest string) ([]byte, int64, error)
	blobMutex       sync.RWMutex
	blobArgsForCall []struct {
		logger   lager.Logger
		imageURL *url.URL
		digest   string
	}
	blobReturns struct {
		result1 []byte
		result2 int64
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSource) Manifest(logger lager.Logger, imageURL *url.URL) (remote.Manifest, error) {
	fake.manifestMutex.Lock()
	fake.manifestArgsForCall = append(fake.manifestArgsForCall, struct {
		logger   lager.Logger
		imageURL *url.URL
	}{logger, imageURL})
	fake.recordInvocation("Manifest", []interface{}{logger, imageURL})
	fake.manifestMutex.Unlock()
	if fake.ManifestStub != nil {
		return fake.ManifestStub(logger, imageURL)
	} else {
		return fake.manifestReturns.result1, fake.manifestReturns.result2
	}
}

func (fake *FakeSource) ManifestCallCount() int {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return len(fake.manifestArgsForCall)
}

func (fake *FakeSource) ManifestArgsForCall(i int) (lager.Logger, *url.URL) {
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	return fake.manifestArgsForCall[i].logger, fake.manifestArgsForCall[i].imageURL
}

func (fake *FakeSource) ManifestReturns(result1 remote.Manifest, result2 error) {
	fake.ManifestStub = nil
	fake.manifestReturns = struct {
		result1 remote.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) Config(logger lager.Logger, imageURL *url.URL, manifest remote.Manifest) (specsv1.Image, error) {
	fake.configMutex.Lock()
	fake.configArgsForCall = append(fake.configArgsForCall, struct {
		logger   lager.Logger
		imageURL *url.URL
		manifest remote.Manifest
	}{logger, imageURL, manifest})
	fake.recordInvocation("Config", []interface{}{logger, imageURL, manifest})
	fake.configMutex.Unlock()
	if fake.ConfigStub != nil {
		return fake.ConfigStub(logger, imageURL, manifest)
	} else {
		return fake.configReturns.result1, fake.configReturns.result2
	}
}

func (fake *FakeSource) ConfigCallCount() int {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return len(fake.configArgsForCall)
}

func (fake *FakeSource) ConfigArgsForCall(i int) (lager.Logger, *url.URL, remote.Manifest) {
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	return fake.configArgsForCall[i].logger, fake.configArgsForCall[i].imageURL, fake.configArgsForCall[i].manifest
}

func (fake *FakeSource) ConfigReturns(result1 specsv1.Image, result2 error) {
	fake.ConfigStub = nil
	fake.configReturns = struct {
		result1 specsv1.Image
		result2 error
	}{result1, result2}
}

func (fake *FakeSource) Blob(logger lager.Logger, imageURL *url.URL, digest string) ([]byte, int64, error) {
	fake.blobMutex.Lock()
	fake.blobArgsForCall = append(fake.blobArgsForCall, struct {
		logger   lager.Logger
		imageURL *url.URL
		digest   string
	}{logger, imageURL, digest})
	fake.recordInvocation("Blob", []interface{}{logger, imageURL, digest})
	fake.blobMutex.Unlock()
	if fake.BlobStub != nil {
		return fake.BlobStub(logger, imageURL, digest)
	} else {
		return fake.blobReturns.result1, fake.blobReturns.result2, fake.blobReturns.result3
	}
}

func (fake *FakeSource) BlobCallCount() int {
	fake.blobMutex.RLock()
	defer fake.blobMutex.RUnlock()
	return len(fake.blobArgsForCall)
}

func (fake *FakeSource) BlobArgsForCall(i int) (lager.Logger, *url.URL, string) {
	fake.blobMutex.RLock()
	defer fake.blobMutex.RUnlock()
	return fake.blobArgsForCall[i].logger, fake.blobArgsForCall[i].imageURL, fake.blobArgsForCall[i].digest
}

func (fake *FakeSource) BlobReturns(result1 []byte, result2 int64, result3 error) {
	fake.BlobStub = nil
	fake.blobReturns = struct {
		result1 []byte
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.manifestMutex.RLock()
	defer fake.manifestMutex.RUnlock()
	fake.configMutex.RLock()
	defer fake.configMutex.RUnlock()
	fake.blobMutex.RLock()
	defer fake.blobMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSource) recordInvocation(key string, args []interface{}) {
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

var _ remote.Source = new(FakeSource)
