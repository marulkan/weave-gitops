// Code generated by counterfeiter. DO NOT EDIT.
package kubefakes

import (
	"context"
	"sync"

	"github.com/weaveworks/weave-gitops/pkg/kube"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeKube struct {
	ApplyStub        func(context.Context, []byte, string) error
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		arg1 context.Context
		arg2 []byte
		arg3 string
	}
	applyReturns struct {
		result1 error
	}
	applyReturnsOnCall map[int]struct {
		result1 error
	}
	GetClusterNameStub        func(context.Context) (string, error)
	getClusterNameMutex       sync.RWMutex
	getClusterNameArgsForCall []struct {
		arg1 context.Context
	}
	getClusterNameReturns struct {
		result1 string
		result2 error
	}
	getClusterNameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetWegoConfigStub        func(context.Context, string) (*kube.WegoConfig, error)
	getWegoConfigMutex       sync.RWMutex
	getWegoConfigArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getWegoConfigReturns struct {
		result1 *kube.WegoConfig
		result2 error
	}
	getWegoConfigReturnsOnCall map[int]struct {
		result1 *kube.WegoConfig
		result2 error
	}
	RawStub        func() client.Client
	rawMutex       sync.RWMutex
	rawArgsForCall []struct {
	}
	rawReturns struct {
		result1 client.Client
	}
	rawReturnsOnCall map[int]struct {
		result1 client.Client
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeKube) Apply(arg1 context.Context, arg2 []byte, arg3 string) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		arg1 context.Context
		arg2 []byte
		arg3 string
	}{arg1, arg2Copy, arg3})
	stub := fake.ApplyStub
	fakeReturns := fake.applyReturns
	fake.recordInvocation("Apply", []interface{}{arg1, arg2Copy, arg3})
	fake.applyMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeKube) ApplyCalls(stub func(context.Context, []byte, string) error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = stub
}

func (fake *FakeKube) ApplyArgsForCall(i int) (context.Context, []byte, string) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	argsForCall := fake.applyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeKube) ApplyReturns(result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) ApplyReturnsOnCall(i int, result1 error) {
	fake.applyMutex.Lock()
	defer fake.applyMutex.Unlock()
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeKube) GetClusterName(arg1 context.Context) (string, error) {
	fake.getClusterNameMutex.Lock()
	ret, specificReturn := fake.getClusterNameReturnsOnCall[len(fake.getClusterNameArgsForCall)]
	fake.getClusterNameArgsForCall = append(fake.getClusterNameArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.GetClusterNameStub
	fakeReturns := fake.getClusterNameReturns
	fake.recordInvocation("GetClusterName", []interface{}{arg1})
	fake.getClusterNameMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetClusterNameCallCount() int {
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	return len(fake.getClusterNameArgsForCall)
}

func (fake *FakeKube) GetClusterNameCalls(stub func(context.Context) (string, error)) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = stub
}

func (fake *FakeKube) GetClusterNameArgsForCall(i int) context.Context {
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	argsForCall := fake.getClusterNameArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeKube) GetClusterNameReturns(result1 string, result2 error) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = nil
	fake.getClusterNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetClusterNameReturnsOnCall(i int, result1 string, result2 error) {
	fake.getClusterNameMutex.Lock()
	defer fake.getClusterNameMutex.Unlock()
	fake.GetClusterNameStub = nil
	if fake.getClusterNameReturnsOnCall == nil {
		fake.getClusterNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getClusterNameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetWegoConfig(arg1 context.Context, arg2 string) (*kube.WegoConfig, error) {
	fake.getWegoConfigMutex.Lock()
	ret, specificReturn := fake.getWegoConfigReturnsOnCall[len(fake.getWegoConfigArgsForCall)]
	fake.getWegoConfigArgsForCall = append(fake.getWegoConfigArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetWegoConfigStub
	fakeReturns := fake.getWegoConfigReturns
	fake.recordInvocation("GetWegoConfig", []interface{}{arg1, arg2})
	fake.getWegoConfigMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeKube) GetWegoConfigCallCount() int {
	fake.getWegoConfigMutex.RLock()
	defer fake.getWegoConfigMutex.RUnlock()
	return len(fake.getWegoConfigArgsForCall)
}

func (fake *FakeKube) GetWegoConfigCalls(stub func(context.Context, string) (*kube.WegoConfig, error)) {
	fake.getWegoConfigMutex.Lock()
	defer fake.getWegoConfigMutex.Unlock()
	fake.GetWegoConfigStub = stub
}

func (fake *FakeKube) GetWegoConfigArgsForCall(i int) (context.Context, string) {
	fake.getWegoConfigMutex.RLock()
	defer fake.getWegoConfigMutex.RUnlock()
	argsForCall := fake.getWegoConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeKube) GetWegoConfigReturns(result1 *kube.WegoConfig, result2 error) {
	fake.getWegoConfigMutex.Lock()
	defer fake.getWegoConfigMutex.Unlock()
	fake.GetWegoConfigStub = nil
	fake.getWegoConfigReturns = struct {
		result1 *kube.WegoConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) GetWegoConfigReturnsOnCall(i int, result1 *kube.WegoConfig, result2 error) {
	fake.getWegoConfigMutex.Lock()
	defer fake.getWegoConfigMutex.Unlock()
	fake.GetWegoConfigStub = nil
	if fake.getWegoConfigReturnsOnCall == nil {
		fake.getWegoConfigReturnsOnCall = make(map[int]struct {
			result1 *kube.WegoConfig
			result2 error
		})
	}
	fake.getWegoConfigReturnsOnCall[i] = struct {
		result1 *kube.WegoConfig
		result2 error
	}{result1, result2}
}

func (fake *FakeKube) Raw() client.Client {
	fake.rawMutex.Lock()
	ret, specificReturn := fake.rawReturnsOnCall[len(fake.rawArgsForCall)]
	fake.rawArgsForCall = append(fake.rawArgsForCall, struct {
	}{})
	stub := fake.RawStub
	fakeReturns := fake.rawReturns
	fake.recordInvocation("Raw", []interface{}{})
	fake.rawMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeKube) RawCallCount() int {
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	return len(fake.rawArgsForCall)
}

func (fake *FakeKube) RawCalls(stub func() client.Client) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = stub
}

func (fake *FakeKube) RawReturns(result1 client.Client) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = nil
	fake.rawReturns = struct {
		result1 client.Client
	}{result1}
}

func (fake *FakeKube) RawReturnsOnCall(i int, result1 client.Client) {
	fake.rawMutex.Lock()
	defer fake.rawMutex.Unlock()
	fake.RawStub = nil
	if fake.rawReturnsOnCall == nil {
		fake.rawReturnsOnCall = make(map[int]struct {
			result1 client.Client
		})
	}
	fake.rawReturnsOnCall[i] = struct {
		result1 client.Client
	}{result1}
}

func (fake *FakeKube) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.getClusterNameMutex.RLock()
	defer fake.getClusterNameMutex.RUnlock()
	fake.getWegoConfigMutex.RLock()
	defer fake.getWegoConfigMutex.RUnlock()
	fake.rawMutex.RLock()
	defer fake.rawMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeKube) recordInvocation(key string, args []interface{}) {
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

var _ kube.Kube = new(FakeKube)
