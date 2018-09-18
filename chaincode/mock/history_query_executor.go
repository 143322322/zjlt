// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	commonledger "github.com/143322322/common/ledger"
)

type HistoryQueryExecutor struct {
	GetHistoryForKeyStub        func(namespace string, key string) (commonledger.ResultsIterator, error)
	getHistoryForKeyMutex       sync.RWMutex
	getHistoryForKeyArgsForCall []struct {
		namespace string
		key       string
	}
	getHistoryForKeyReturns struct {
		result1 commonledger.ResultsIterator
		result2 error
	}
	getHistoryForKeyReturnsOnCall map[int]struct {
		result1 commonledger.ResultsIterator
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HistoryQueryExecutor) GetHistoryForKey(namespace string, key string) (commonledger.ResultsIterator, error) {
	fake.getHistoryForKeyMutex.Lock()
	ret, specificReturn := fake.getHistoryForKeyReturnsOnCall[len(fake.getHistoryForKeyArgsForCall)]
	fake.getHistoryForKeyArgsForCall = append(fake.getHistoryForKeyArgsForCall, struct {
		namespace string
		key       string
	}{namespace, key})
	fake.recordInvocation("GetHistoryForKey", []interface{}{namespace, key})
	fake.getHistoryForKeyMutex.Unlock()
	if fake.GetHistoryForKeyStub != nil {
		return fake.GetHistoryForKeyStub(namespace, key)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getHistoryForKeyReturns.result1, fake.getHistoryForKeyReturns.result2
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyCallCount() int {
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	return len(fake.getHistoryForKeyArgsForCall)
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyArgsForCall(i int) (string, string) {
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	return fake.getHistoryForKeyArgsForCall[i].namespace, fake.getHistoryForKeyArgsForCall[i].key
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyReturns(result1 commonledger.ResultsIterator, result2 error) {
	fake.GetHistoryForKeyStub = nil
	fake.getHistoryForKeyReturns = struct {
		result1 commonledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *HistoryQueryExecutor) GetHistoryForKeyReturnsOnCall(i int, result1 commonledger.ResultsIterator, result2 error) {
	fake.GetHistoryForKeyStub = nil
	if fake.getHistoryForKeyReturnsOnCall == nil {
		fake.getHistoryForKeyReturnsOnCall = make(map[int]struct {
			result1 commonledger.ResultsIterator
			result2 error
		})
	}
	fake.getHistoryForKeyReturnsOnCall[i] = struct {
		result1 commonledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *HistoryQueryExecutor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getHistoryForKeyMutex.RLock()
	defer fake.getHistoryForKeyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *HistoryQueryExecutor) recordInvocation(key string, args []interface{}) {
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
