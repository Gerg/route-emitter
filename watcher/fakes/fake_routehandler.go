// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/bbs/models"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/route-emitter/routing_table/schema/endpoint"
	"code.cloudfoundry.org/route-emitter/watcher"
)

type FakeRouteHandler struct {
	HandleEventStub        func(logger lager.Logger, event models.Event)
	handleEventMutex       sync.RWMutex
	handleEventArgsForCall []struct {
		logger lager.Logger
		event  models.Event
	}
	SyncStub        func(logger lager.Logger, desired []*models.DesiredLRPSchedulingInfo, runningActual []*endpoint.ActualLRPRoutingInfo, domains models.DomainSet)
	syncMutex       sync.RWMutex
	syncArgsForCall []struct {
		logger        lager.Logger
		desired       []*models.DesiredLRPSchedulingInfo
		runningActual []*endpoint.ActualLRPRoutingInfo
		domains       models.DomainSet
	}
	ShouldRefreshDesiredStub        func(*endpoint.ActualLRPRoutingInfo) bool
	shouldRefreshDesiredMutex       sync.RWMutex
	shouldRefreshDesiredArgsForCall []struct {
		arg1 *endpoint.ActualLRPRoutingInfo
	}
	shouldRefreshDesiredReturns struct {
		result1 bool
	}
	shouldRefreshDesiredReturnsOnCall map[int]struct {
		result1 bool
	}
	RefreshDesiredStub        func(lager.Logger, []*models.DesiredLRPSchedulingInfo)
	refreshDesiredMutex       sync.RWMutex
	refreshDesiredArgsForCall []struct {
		arg1 lager.Logger
		arg2 []*models.DesiredLRPSchedulingInfo
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRouteHandler) HandleEvent(logger lager.Logger, event models.Event) {
	fake.handleEventMutex.Lock()
	fake.handleEventArgsForCall = append(fake.handleEventArgsForCall, struct {
		logger lager.Logger
		event  models.Event
	}{logger, event})
	fake.recordInvocation("HandleEvent", []interface{}{logger, event})
	fake.handleEventMutex.Unlock()
	if fake.HandleEventStub != nil {
		fake.HandleEventStub(logger, event)
	}
}

func (fake *FakeRouteHandler) HandleEventCallCount() int {
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	return len(fake.handleEventArgsForCall)
}

func (fake *FakeRouteHandler) HandleEventArgsForCall(i int) (lager.Logger, models.Event) {
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	return fake.handleEventArgsForCall[i].logger, fake.handleEventArgsForCall[i].event
}

func (fake *FakeRouteHandler) Sync(logger lager.Logger, desired []*models.DesiredLRPSchedulingInfo, runningActual []*endpoint.ActualLRPRoutingInfo, domains models.DomainSet) {
	var desiredCopy []*models.DesiredLRPSchedulingInfo
	if desired != nil {
		desiredCopy = make([]*models.DesiredLRPSchedulingInfo, len(desired))
		copy(desiredCopy, desired)
	}
	var runningActualCopy []*endpoint.ActualLRPRoutingInfo
	if runningActual != nil {
		runningActualCopy = make([]*endpoint.ActualLRPRoutingInfo, len(runningActual))
		copy(runningActualCopy, runningActual)
	}
	fake.syncMutex.Lock()
	fake.syncArgsForCall = append(fake.syncArgsForCall, struct {
		logger        lager.Logger
		desired       []*models.DesiredLRPSchedulingInfo
		runningActual []*endpoint.ActualLRPRoutingInfo
		domains       models.DomainSet
	}{logger, desiredCopy, runningActualCopy, domains})
	fake.recordInvocation("Sync", []interface{}{logger, desiredCopy, runningActualCopy, domains})
	fake.syncMutex.Unlock()
	if fake.SyncStub != nil {
		fake.SyncStub(logger, desired, runningActual, domains)
	}
}

func (fake *FakeRouteHandler) SyncCallCount() int {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return len(fake.syncArgsForCall)
}

func (fake *FakeRouteHandler) SyncArgsForCall(i int) (lager.Logger, []*models.DesiredLRPSchedulingInfo, []*endpoint.ActualLRPRoutingInfo, models.DomainSet) {
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	return fake.syncArgsForCall[i].logger, fake.syncArgsForCall[i].desired, fake.syncArgsForCall[i].runningActual, fake.syncArgsForCall[i].domains
}

func (fake *FakeRouteHandler) ShouldRefreshDesired(arg1 *endpoint.ActualLRPRoutingInfo) bool {
	fake.shouldRefreshDesiredMutex.Lock()
	ret, specificReturn := fake.shouldRefreshDesiredReturnsOnCall[len(fake.shouldRefreshDesiredArgsForCall)]
	fake.shouldRefreshDesiredArgsForCall = append(fake.shouldRefreshDesiredArgsForCall, struct {
		arg1 *endpoint.ActualLRPRoutingInfo
	}{arg1})
	fake.recordInvocation("ShouldRefreshDesired", []interface{}{arg1})
	fake.shouldRefreshDesiredMutex.Unlock()
	if fake.ShouldRefreshDesiredStub != nil {
		return fake.ShouldRefreshDesiredStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.shouldRefreshDesiredReturns.result1
}

func (fake *FakeRouteHandler) ShouldRefreshDesiredCallCount() int {
	fake.shouldRefreshDesiredMutex.RLock()
	defer fake.shouldRefreshDesiredMutex.RUnlock()
	return len(fake.shouldRefreshDesiredArgsForCall)
}

func (fake *FakeRouteHandler) ShouldRefreshDesiredArgsForCall(i int) *endpoint.ActualLRPRoutingInfo {
	fake.shouldRefreshDesiredMutex.RLock()
	defer fake.shouldRefreshDesiredMutex.RUnlock()
	return fake.shouldRefreshDesiredArgsForCall[i].arg1
}

func (fake *FakeRouteHandler) ShouldRefreshDesiredReturns(result1 bool) {
	fake.ShouldRefreshDesiredStub = nil
	fake.shouldRefreshDesiredReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRouteHandler) ShouldRefreshDesiredReturnsOnCall(i int, result1 bool) {
	fake.ShouldRefreshDesiredStub = nil
	if fake.shouldRefreshDesiredReturnsOnCall == nil {
		fake.shouldRefreshDesiredReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.shouldRefreshDesiredReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRouteHandler) RefreshDesired(arg1 lager.Logger, arg2 []*models.DesiredLRPSchedulingInfo) {
	var arg2Copy []*models.DesiredLRPSchedulingInfo
	if arg2 != nil {
		arg2Copy = make([]*models.DesiredLRPSchedulingInfo, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.refreshDesiredMutex.Lock()
	fake.refreshDesiredArgsForCall = append(fake.refreshDesiredArgsForCall, struct {
		arg1 lager.Logger
		arg2 []*models.DesiredLRPSchedulingInfo
	}{arg1, arg2Copy})
	fake.recordInvocation("RefreshDesired", []interface{}{arg1, arg2Copy})
	fake.refreshDesiredMutex.Unlock()
	if fake.RefreshDesiredStub != nil {
		fake.RefreshDesiredStub(arg1, arg2)
	}
}

func (fake *FakeRouteHandler) RefreshDesiredCallCount() int {
	fake.refreshDesiredMutex.RLock()
	defer fake.refreshDesiredMutex.RUnlock()
	return len(fake.refreshDesiredArgsForCall)
}

func (fake *FakeRouteHandler) RefreshDesiredArgsForCall(i int) (lager.Logger, []*models.DesiredLRPSchedulingInfo) {
	fake.refreshDesiredMutex.RLock()
	defer fake.refreshDesiredMutex.RUnlock()
	return fake.refreshDesiredArgsForCall[i].arg1, fake.refreshDesiredArgsForCall[i].arg2
}

func (fake *FakeRouteHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleEventMutex.RLock()
	defer fake.handleEventMutex.RUnlock()
	fake.syncMutex.RLock()
	defer fake.syncMutex.RUnlock()
	fake.shouldRefreshDesiredMutex.RLock()
	defer fake.shouldRefreshDesiredMutex.RUnlock()
	fake.refreshDesiredMutex.RLock()
	defer fake.refreshDesiredMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRouteHandler) recordInvocation(key string, args []interface{}) {
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

var _ watcher.RouteHandler = new(FakeRouteHandler)
