// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/peer"
)

type QueryProvider struct {
	TransactionStatusStub        func(string, string) (peer.TxValidationCode, error)
	transactionStatusMutex       sync.RWMutex
	transactionStatusArgsForCall []struct {
		arg1 string
		arg2 string
	}
	transactionStatusReturns struct {
		result1 peer.TxValidationCode
		result2 error
	}
	transactionStatusReturnsOnCall map[int]struct {
		result1 peer.TxValidationCode
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *QueryProvider) TransactionStatus(arg1 string, arg2 string) (peer.TxValidationCode, error) {
	fake.transactionStatusMutex.Lock()
	ret, specificReturn := fake.transactionStatusReturnsOnCall[len(fake.transactionStatusArgsForCall)]
	fake.transactionStatusArgsForCall = append(fake.transactionStatusArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.TransactionStatusStub
	fakeReturns := fake.transactionStatusReturns
	fake.recordInvocation("TransactionStatus", []interface{}{arg1, arg2})
	fake.transactionStatusMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *QueryProvider) TransactionStatusCallCount() int {
	fake.transactionStatusMutex.RLock()
	defer fake.transactionStatusMutex.RUnlock()
	return len(fake.transactionStatusArgsForCall)
}

func (fake *QueryProvider) TransactionStatusCalls(stub func(string, string) (peer.TxValidationCode, error)) {
	fake.transactionStatusMutex.Lock()
	defer fake.transactionStatusMutex.Unlock()
	fake.TransactionStatusStub = stub
}

func (fake *QueryProvider) TransactionStatusArgsForCall(i int) (string, string) {
	fake.transactionStatusMutex.RLock()
	defer fake.transactionStatusMutex.RUnlock()
	argsForCall := fake.transactionStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *QueryProvider) TransactionStatusReturns(result1 peer.TxValidationCode, result2 error) {
	fake.transactionStatusMutex.Lock()
	defer fake.transactionStatusMutex.Unlock()
	fake.TransactionStatusStub = nil
	fake.transactionStatusReturns = struct {
		result1 peer.TxValidationCode
		result2 error
	}{result1, result2}
}

func (fake *QueryProvider) TransactionStatusReturnsOnCall(i int, result1 peer.TxValidationCode, result2 error) {
	fake.transactionStatusMutex.Lock()
	defer fake.transactionStatusMutex.Unlock()
	fake.TransactionStatusStub = nil
	if fake.transactionStatusReturnsOnCall == nil {
		fake.transactionStatusReturnsOnCall = make(map[int]struct {
			result1 peer.TxValidationCode
			result2 error
		})
	}
	fake.transactionStatusReturnsOnCall[i] = struct {
		result1 peer.TxValidationCode
		result2 error
	}{result1, result2}
}

func (fake *QueryProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.transactionStatusMutex.RLock()
	defer fake.transactionStatusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *QueryProvider) recordInvocation(key string, args []interface{}) {
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
