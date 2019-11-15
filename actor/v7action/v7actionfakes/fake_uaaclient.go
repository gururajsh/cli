// Code generated by counterfeiter. DO NOT EDIT.
package v7actionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/uaa"
)

type FakeUAAClient struct {
	CreateUserStub        func(string, string, string) (uaa.User, error)
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	createUserReturns struct {
		result1 uaa.User
		result2 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 uaa.User
		result2 error
	}
	DeleteUserStub        func(string) (uaa.User, error)
	deleteUserMutex       sync.RWMutex
	deleteUserArgsForCall []struct {
		arg1 string
	}
	deleteUserReturns struct {
		result1 uaa.User
		result2 error
	}
	deleteUserReturnsOnCall map[int]struct {
		result1 uaa.User
		result2 error
	}
	GetSSHPasscodeStub        func(string, string) (string, error)
	getSSHPasscodeMutex       sync.RWMutex
	getSSHPasscodeArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSSHPasscodeReturns struct {
		result1 string
		result2 error
	}
	getSSHPasscodeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	ListUsersStub        func(string, string) ([]uaa.User, error)
	listUsersMutex       sync.RWMutex
	listUsersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	listUsersReturns struct {
		result1 []uaa.User
		result2 error
	}
	listUsersReturnsOnCall map[int]struct {
		result1 []uaa.User
		result2 error
	}
	RefreshAccessTokenStub        func(string) (uaa.RefreshedTokens, error)
	refreshAccessTokenMutex       sync.RWMutex
	refreshAccessTokenArgsForCall []struct {
		arg1 string
	}
	refreshAccessTokenReturns struct {
		result1 uaa.RefreshedTokens
		result2 error
	}
	refreshAccessTokenReturnsOnCall map[int]struct {
		result1 uaa.RefreshedTokens
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAAClient) CreateUser(arg1 string, arg2 string, arg3 string) (uaa.User, error) {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateUser", []interface{}{arg1, arg2, arg3})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeUAAClient) CreateUserCalls(stub func(string, string, string) (uaa.User, error)) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = stub
}

func (fake *FakeUAAClient) CreateUserArgsForCall(i int) (string, string, string) {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	argsForCall := fake.createUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUAAClient) CreateUserReturns(result1 uaa.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) CreateUserReturnsOnCall(i int, result1 uaa.User, result2 error) {
	fake.createUserMutex.Lock()
	defer fake.createUserMutex.Unlock()
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 uaa.User
			result2 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) DeleteUser(arg1 string) (uaa.User, error) {
	fake.deleteUserMutex.Lock()
	ret, specificReturn := fake.deleteUserReturnsOnCall[len(fake.deleteUserArgsForCall)]
	fake.deleteUserArgsForCall = append(fake.deleteUserArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteUser", []interface{}{arg1})
	fake.deleteUserMutex.Unlock()
	if fake.DeleteUserStub != nil {
		return fake.DeleteUserStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) DeleteUserCallCount() int {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	return len(fake.deleteUserArgsForCall)
}

func (fake *FakeUAAClient) DeleteUserCalls(stub func(string) (uaa.User, error)) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = stub
}

func (fake *FakeUAAClient) DeleteUserArgsForCall(i int) string {
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	argsForCall := fake.deleteUserArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) DeleteUserReturns(result1 uaa.User, result2 error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = nil
	fake.deleteUserReturns = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) DeleteUserReturnsOnCall(i int, result1 uaa.User, result2 error) {
	fake.deleteUserMutex.Lock()
	defer fake.deleteUserMutex.Unlock()
	fake.DeleteUserStub = nil
	if fake.deleteUserReturnsOnCall == nil {
		fake.deleteUserReturnsOnCall = make(map[int]struct {
			result1 uaa.User
			result2 error
		})
	}
	fake.deleteUserReturnsOnCall[i] = struct {
		result1 uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetSSHPasscode(arg1 string, arg2 string) (string, error) {
	fake.getSSHPasscodeMutex.Lock()
	ret, specificReturn := fake.getSSHPasscodeReturnsOnCall[len(fake.getSSHPasscodeArgsForCall)]
	fake.getSSHPasscodeArgsForCall = append(fake.getSSHPasscodeArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSSHPasscode", []interface{}{arg1, arg2})
	fake.getSSHPasscodeMutex.Unlock()
	if fake.GetSSHPasscodeStub != nil {
		return fake.GetSSHPasscodeStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getSSHPasscodeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) GetSSHPasscodeCallCount() int {
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	return len(fake.getSSHPasscodeArgsForCall)
}

func (fake *FakeUAAClient) GetSSHPasscodeCalls(stub func(string, string) (string, error)) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = stub
}

func (fake *FakeUAAClient) GetSSHPasscodeArgsForCall(i int) (string, string) {
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	argsForCall := fake.getSSHPasscodeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUAAClient) GetSSHPasscodeReturns(result1 string, result2 error) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = nil
	fake.getSSHPasscodeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) GetSSHPasscodeReturnsOnCall(i int, result1 string, result2 error) {
	fake.getSSHPasscodeMutex.Lock()
	defer fake.getSSHPasscodeMutex.Unlock()
	fake.GetSSHPasscodeStub = nil
	if fake.getSSHPasscodeReturnsOnCall == nil {
		fake.getSSHPasscodeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getSSHPasscodeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) ListUsers(arg1 string, arg2 string) ([]uaa.User, error) {
	fake.listUsersMutex.Lock()
	ret, specificReturn := fake.listUsersReturnsOnCall[len(fake.listUsersArgsForCall)]
	fake.listUsersArgsForCall = append(fake.listUsersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ListUsers", []interface{}{arg1, arg2})
	fake.listUsersMutex.Unlock()
	if fake.ListUsersStub != nil {
		return fake.ListUsersStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listUsersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) ListUsersCallCount() int {
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	return len(fake.listUsersArgsForCall)
}

func (fake *FakeUAAClient) ListUsersCalls(stub func(string, string) ([]uaa.User, error)) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = stub
}

func (fake *FakeUAAClient) ListUsersArgsForCall(i int) (string, string) {
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	argsForCall := fake.listUsersArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUAAClient) ListUsersReturns(result1 []uaa.User, result2 error) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = nil
	fake.listUsersReturns = struct {
		result1 []uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) ListUsersReturnsOnCall(i int, result1 []uaa.User, result2 error) {
	fake.listUsersMutex.Lock()
	defer fake.listUsersMutex.Unlock()
	fake.ListUsersStub = nil
	if fake.listUsersReturnsOnCall == nil {
		fake.listUsersReturnsOnCall = make(map[int]struct {
			result1 []uaa.User
			result2 error
		})
	}
	fake.listUsersReturnsOnCall[i] = struct {
		result1 []uaa.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) RefreshAccessToken(arg1 string) (uaa.RefreshedTokens, error) {
	fake.refreshAccessTokenMutex.Lock()
	ret, specificReturn := fake.refreshAccessTokenReturnsOnCall[len(fake.refreshAccessTokenArgsForCall)]
	fake.refreshAccessTokenArgsForCall = append(fake.refreshAccessTokenArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RefreshAccessToken", []interface{}{arg1})
	fake.refreshAccessTokenMutex.Unlock()
	if fake.RefreshAccessTokenStub != nil {
		return fake.RefreshAccessTokenStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.refreshAccessTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAAClient) RefreshAccessTokenCallCount() int {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	return len(fake.refreshAccessTokenArgsForCall)
}

func (fake *FakeUAAClient) RefreshAccessTokenCalls(stub func(string) (uaa.RefreshedTokens, error)) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = stub
}

func (fake *FakeUAAClient) RefreshAccessTokenArgsForCall(i int) string {
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	argsForCall := fake.refreshAccessTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUAAClient) RefreshAccessTokenReturns(result1 uaa.RefreshedTokens, result2 error) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = nil
	fake.refreshAccessTokenReturns = struct {
		result1 uaa.RefreshedTokens
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) RefreshAccessTokenReturnsOnCall(i int, result1 uaa.RefreshedTokens, result2 error) {
	fake.refreshAccessTokenMutex.Lock()
	defer fake.refreshAccessTokenMutex.Unlock()
	fake.RefreshAccessTokenStub = nil
	if fake.refreshAccessTokenReturnsOnCall == nil {
		fake.refreshAccessTokenReturnsOnCall = make(map[int]struct {
			result1 uaa.RefreshedTokens
			result2 error
		})
	}
	fake.refreshAccessTokenReturnsOnCall[i] = struct {
		result1 uaa.RefreshedTokens
		result2 error
	}{result1, result2}
}

func (fake *FakeUAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.deleteUserMutex.RLock()
	defer fake.deleteUserMutex.RUnlock()
	fake.getSSHPasscodeMutex.RLock()
	defer fake.getSSHPasscodeMutex.RUnlock()
	fake.listUsersMutex.RLock()
	defer fake.listUsersMutex.RUnlock()
	fake.refreshAccessTokenMutex.RLock()
	defer fake.refreshAccessTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAAClient) recordInvocation(key string, args []interface{}) {
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

var _ v7action.UAAClient = new(FakeUAAClient)