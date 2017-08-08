// This file was generated by counterfeiter
package modelsfakes

import (
	"gcp-service-broker/brokerapi/brokers/models"
	"sync"
)

type FakeAccountManager struct {
	CreateCredentialsStub        func(instanceID string, bindingID string, details models.BindDetails, instance models.ServiceInstanceDetails) (models.ServiceBindingCredentials, error)
	createCredentialsMutex       sync.RWMutex
	createCredentialsArgsForCall []struct {
		instanceID string
		bindingID  string
		details    models.BindDetails
		instance   models.ServiceInstanceDetails
	}
	createCredentialsReturns struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}
	DeleteCredentialsStub        func(creds models.ServiceBindingCredentials) error
	deleteCredentialsMutex       sync.RWMutex
	deleteCredentialsArgsForCall []struct {
		creds models.ServiceBindingCredentials
	}
	deleteCredentialsReturns struct {
		result1 error
	}
	BuildInstanceCredentialsStub        func(bindDetails map[string]string, instanceDetails map[string]string) map[string]string
	buildInstanceCredentialsMutex       sync.RWMutex
	buildInstanceCredentialsArgsForCall []struct {
		bindDetails     map[string]string
		instanceDetails map[string]string
	}
	buildInstanceCredentialsReturns struct {
		result1 map[string]string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAccountManager) CreateCredentials(instanceID string, bindingID string, details models.BindDetails, instance models.ServiceInstanceDetails) (models.ServiceBindingCredentials, error) {
	fake.createCredentialsMutex.Lock()
	fake.createCredentialsArgsForCall = append(fake.createCredentialsArgsForCall, struct {
		instanceID string
		bindingID  string
		details    models.BindDetails
		instance   models.ServiceInstanceDetails
	}{instanceID, bindingID, details, instance})
	fake.recordInvocation("CreateCredentials", []interface{}{instanceID, bindingID, details, instance})
	fake.createCredentialsMutex.Unlock()
	if fake.CreateCredentialsStub != nil {
		return fake.CreateCredentialsStub(instanceID, bindingID, details, instance)
	} else {
		return fake.createCredentialsReturns.result1, fake.createCredentialsReturns.result2
	}
}

func (fake *FakeAccountManager) CreateCredentialsCallCount() int {
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	return len(fake.createCredentialsArgsForCall)
}

func (fake *FakeAccountManager) CreateCredentialsArgsForCall(i int) (string, string, models.BindDetails, models.ServiceInstanceDetails) {
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	return fake.createCredentialsArgsForCall[i].instanceID, fake.createCredentialsArgsForCall[i].bindingID, fake.createCredentialsArgsForCall[i].details, fake.createCredentialsArgsForCall[i].instance
}

func (fake *FakeAccountManager) CreateCredentialsReturns(result1 models.ServiceBindingCredentials, result2 error) {
	fake.CreateCredentialsStub = nil
	fake.createCredentialsReturns = struct {
		result1 models.ServiceBindingCredentials
		result2 error
	}{result1, result2}
}

func (fake *FakeAccountManager) DeleteCredentials(creds models.ServiceBindingCredentials) error {
	fake.deleteCredentialsMutex.Lock()
	fake.deleteCredentialsArgsForCall = append(fake.deleteCredentialsArgsForCall, struct {
		creds models.ServiceBindingCredentials
	}{creds})
	fake.recordInvocation("DeleteCredentials", []interface{}{creds})
	fake.deleteCredentialsMutex.Unlock()
	if fake.DeleteCredentialsStub != nil {
		return fake.DeleteCredentialsStub(creds)
	} else {
		return fake.deleteCredentialsReturns.result1
	}
}

func (fake *FakeAccountManager) DeleteCredentialsCallCount() int {
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	return len(fake.deleteCredentialsArgsForCall)
}

func (fake *FakeAccountManager) DeleteCredentialsArgsForCall(i int) models.ServiceBindingCredentials {
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	return fake.deleteCredentialsArgsForCall[i].creds
}

func (fake *FakeAccountManager) DeleteCredentialsReturns(result1 error) {
	fake.DeleteCredentialsStub = nil
	fake.deleteCredentialsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAccountManager) BuildInstanceCredentials(bindDetails map[string]string, instanceDetails map[string]string) map[string]string {
	fake.buildInstanceCredentialsMutex.Lock()
	fake.buildInstanceCredentialsArgsForCall = append(fake.buildInstanceCredentialsArgsForCall, struct {
		bindDetails     map[string]string
		instanceDetails map[string]string
	}{bindDetails, instanceDetails})
	fake.recordInvocation("BuildInstanceCredentials", []interface{}{bindDetails, instanceDetails})
	fake.buildInstanceCredentialsMutex.Unlock()
	if fake.BuildInstanceCredentialsStub != nil {
		return fake.BuildInstanceCredentialsStub(bindDetails, instanceDetails)
	} else {
		return fake.buildInstanceCredentialsReturns.result1
	}
}

func (fake *FakeAccountManager) BuildInstanceCredentialsCallCount() int {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return len(fake.buildInstanceCredentialsArgsForCall)
}

func (fake *FakeAccountManager) BuildInstanceCredentialsArgsForCall(i int) (map[string]string, map[string]string) {
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return fake.buildInstanceCredentialsArgsForCall[i].bindDetails, fake.buildInstanceCredentialsArgsForCall[i].instanceDetails
}

func (fake *FakeAccountManager) BuildInstanceCredentialsReturns(result1 map[string]string) {
	fake.BuildInstanceCredentialsStub = nil
	fake.buildInstanceCredentialsReturns = struct {
		result1 map[string]string
	}{result1}
}

func (fake *FakeAccountManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createCredentialsMutex.RLock()
	defer fake.createCredentialsMutex.RUnlock()
	fake.deleteCredentialsMutex.RLock()
	defer fake.deleteCredentialsMutex.RUnlock()
	fake.buildInstanceCredentialsMutex.RLock()
	defer fake.buildInstanceCredentialsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAccountManager) recordInvocation(key string, args []interface{}) {
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

var _ models.AccountManager = new(FakeAccountManager)
