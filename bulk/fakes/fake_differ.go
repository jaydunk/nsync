// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/nsync/bulk"
	"github.com/cloudfoundry-incubator/runtime-schema/cc_messages"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
)

type FakeDiffer struct {
	DiffStub        func(existing []models.DesiredLRP, newChan <-chan cc_messages.DesireAppRequestFromCC) <-chan models.DesiredLRPChange
	diffMutex       sync.RWMutex
	diffArgsForCall []struct {
		existing []models.DesiredLRP
		newChan  <-chan cc_messages.DesireAppRequestFromCC
	}
	diffReturns struct {
		result1 <-chan models.DesiredLRPChange
	}
}

func (fake *FakeDiffer) Diff(existing []models.DesiredLRP, newChan <-chan cc_messages.DesireAppRequestFromCC) <-chan models.DesiredLRPChange {
	fake.diffMutex.Lock()
	defer fake.diffMutex.Unlock()
	fake.diffArgsForCall = append(fake.diffArgsForCall, struct {
		existing []models.DesiredLRP
		newChan  <-chan cc_messages.DesireAppRequestFromCC
	}{existing, newChan})
	if fake.DiffStub != nil {
		return fake.DiffStub(existing, newChan)
	} else {
		return fake.diffReturns.result1
	}
}

func (fake *FakeDiffer) DiffCallCount() int {
	fake.diffMutex.RLock()
	defer fake.diffMutex.RUnlock()
	return len(fake.diffArgsForCall)
}

func (fake *FakeDiffer) DiffArgsForCall(i int) ([]models.DesiredLRP, <-chan cc_messages.DesireAppRequestFromCC) {
	fake.diffMutex.RLock()
	defer fake.diffMutex.RUnlock()
	return fake.diffArgsForCall[i].existing, fake.diffArgsForCall[i].newChan
}

func (fake *FakeDiffer) DiffReturns(result1 <-chan models.DesiredLRPChange) {
	fake.diffReturns = struct {
		result1 <-chan models.DesiredLRPChange
	}{result1}
}

var _ bulk.Differ = new(FakeDiffer)