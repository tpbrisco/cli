// Code generated by counterfeiter. DO NOT EDIT.
package compositefakes

import (
	sync "sync"

	v2action "code.cloudfoundry.org/cli/actor/v2action"
	wrappers "code.cloudfoundry.org/cli/actor/v2action/composite"
)

type FakeOrganizationActor struct {
	GetOrganizationStub        func(string) (v2action.Organization, v2action.Warnings, error)
	getOrganizationMutex       sync.RWMutex
	getOrganizationArgsForCall []struct {
		arg1 string
	}
	getOrganizationReturns struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	getOrganizationReturnsOnCall map[int]struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrganizationActor) GetOrganization(arg1 string) (v2action.Organization, v2action.Warnings, error) {
	fake.getOrganizationMutex.Lock()
	ret, specificReturn := fake.getOrganizationReturnsOnCall[len(fake.getOrganizationArgsForCall)]
	fake.getOrganizationArgsForCall = append(fake.getOrganizationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetOrganization", []interface{}{arg1})
	fake.getOrganizationMutex.Unlock()
	if fake.GetOrganizationStub != nil {
		return fake.GetOrganizationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeOrganizationActor) GetOrganizationCallCount() int {
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	return len(fake.getOrganizationArgsForCall)
}

func (fake *FakeOrganizationActor) GetOrganizationCalls(stub func(string) (v2action.Organization, v2action.Warnings, error)) {
	fake.getOrganizationMutex.Lock()
	defer fake.getOrganizationMutex.Unlock()
	fake.GetOrganizationStub = stub
}

func (fake *FakeOrganizationActor) GetOrganizationArgsForCall(i int) string {
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	argsForCall := fake.getOrganizationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeOrganizationActor) GetOrganizationReturns(result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.getOrganizationMutex.Lock()
	defer fake.getOrganizationMutex.Unlock()
	fake.GetOrganizationStub = nil
	fake.getOrganizationReturns = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrganizationActor) GetOrganizationReturnsOnCall(i int, result1 v2action.Organization, result2 v2action.Warnings, result3 error) {
	fake.getOrganizationMutex.Lock()
	defer fake.getOrganizationMutex.Unlock()
	fake.GetOrganizationStub = nil
	if fake.getOrganizationReturnsOnCall == nil {
		fake.getOrganizationReturnsOnCall = make(map[int]struct {
			result1 v2action.Organization
			result2 v2action.Warnings
			result3 error
		})
	}
	fake.getOrganizationReturnsOnCall[i] = struct {
		result1 v2action.Organization
		result2 v2action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeOrganizationActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOrganizationActor) recordInvocation(key string, args []interface{}) {
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

var _ wrappers.OrganizationActor = new(FakeOrganizationActor)
