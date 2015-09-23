// This file was generated by counterfeiter
package mocks

import (
	"io"
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/ssh/sshapi"
	"golang.org/x/crypto/ssh"
)

type FakeSSHSession struct {
	StdinPipeStub        func() (io.WriteCloser, error)
	stdinPipeMutex       sync.RWMutex
	stdinPipeArgsForCall []struct{}
	stdinPipeReturns     struct {
		result1 io.WriteCloser
		result2 error
	}
	StdoutPipeStub        func() (io.Reader, error)
	stdoutPipeMutex       sync.RWMutex
	stdoutPipeArgsForCall []struct{}
	stdoutPipeReturns     struct {
		result1 io.Reader
		result2 error
	}
	StderrPipeStub        func() (io.Reader, error)
	stderrPipeMutex       sync.RWMutex
	stderrPipeArgsForCall []struct{}
	stderrPipeReturns     struct {
		result1 io.Reader
		result2 error
	}
	SendRequestStub        func(name string, wantReply bool, payload []byte) (bool, error)
	sendRequestMutex       sync.RWMutex
	sendRequestArgsForCall []struct {
		name      string
		wantReply bool
		payload   []byte
	}
	sendRequestReturns struct {
		result1 bool
		result2 error
	}
	RequestPtyStub        func(term string, h, w int, termmodes ssh.TerminalModes) error
	requestPtyMutex       sync.RWMutex
	requestPtyArgsForCall []struct {
		term      string
		h         int
		w         int
		termmodes ssh.TerminalModes
	}
	requestPtyReturns struct {
		result1 error
	}
	ShellStub        func() error
	shellMutex       sync.RWMutex
	shellArgsForCall []struct{}
	shellReturns     struct {
		result1 error
	}
	RunStub        func(string) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 string
	}
	runReturns struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
}

func (fake *FakeSSHSession) StdinPipe() (io.WriteCloser, error) {
	fake.stdinPipeMutex.Lock()
	fake.stdinPipeArgsForCall = append(fake.stdinPipeArgsForCall, struct{}{})
	fake.stdinPipeMutex.Unlock()
	if fake.StdinPipeStub != nil {
		return fake.StdinPipeStub()
	} else {
		return fake.stdinPipeReturns.result1, fake.stdinPipeReturns.result2
	}
}

func (fake *FakeSSHSession) StdinPipeCallCount() int {
	fake.stdinPipeMutex.RLock()
	defer fake.stdinPipeMutex.RUnlock()
	return len(fake.stdinPipeArgsForCall)
}

func (fake *FakeSSHSession) StdinPipeReturns(result1 io.WriteCloser, result2 error) {
	fake.StdinPipeStub = nil
	fake.stdinPipeReturns = struct {
		result1 io.WriteCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHSession) StdoutPipe() (io.Reader, error) {
	fake.stdoutPipeMutex.Lock()
	fake.stdoutPipeArgsForCall = append(fake.stdoutPipeArgsForCall, struct{}{})
	fake.stdoutPipeMutex.Unlock()
	if fake.StdoutPipeStub != nil {
		return fake.StdoutPipeStub()
	} else {
		return fake.stdoutPipeReturns.result1, fake.stdoutPipeReturns.result2
	}
}

func (fake *FakeSSHSession) StdoutPipeCallCount() int {
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	return len(fake.stdoutPipeArgsForCall)
}

func (fake *FakeSSHSession) StdoutPipeReturns(result1 io.Reader, result2 error) {
	fake.StdoutPipeStub = nil
	fake.stdoutPipeReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHSession) StderrPipe() (io.Reader, error) {
	fake.stderrPipeMutex.Lock()
	fake.stderrPipeArgsForCall = append(fake.stderrPipeArgsForCall, struct{}{})
	fake.stderrPipeMutex.Unlock()
	if fake.StderrPipeStub != nil {
		return fake.StderrPipeStub()
	} else {
		return fake.stderrPipeReturns.result1, fake.stderrPipeReturns.result2
	}
}

func (fake *FakeSSHSession) StderrPipeCallCount() int {
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	return len(fake.stderrPipeArgsForCall)
}

func (fake *FakeSSHSession) StderrPipeReturns(result1 io.Reader, result2 error) {
	fake.StderrPipeStub = nil
	fake.stderrPipeReturns = struct {
		result1 io.Reader
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHSession) SendRequest(name string, wantReply bool, payload []byte) (bool, error) {
	fake.sendRequestMutex.Lock()
	fake.sendRequestArgsForCall = append(fake.sendRequestArgsForCall, struct {
		name      string
		wantReply bool
		payload   []byte
	}{name, wantReply, payload})
	fake.sendRequestMutex.Unlock()
	if fake.SendRequestStub != nil {
		return fake.SendRequestStub(name, wantReply, payload)
	} else {
		return fake.sendRequestReturns.result1, fake.sendRequestReturns.result2
	}
}

func (fake *FakeSSHSession) SendRequestCallCount() int {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return len(fake.sendRequestArgsForCall)
}

func (fake *FakeSSHSession) SendRequestArgsForCall(i int) (string, bool, []byte) {
	fake.sendRequestMutex.RLock()
	defer fake.sendRequestMutex.RUnlock()
	return fake.sendRequestArgsForCall[i].name, fake.sendRequestArgsForCall[i].wantReply, fake.sendRequestArgsForCall[i].payload
}

func (fake *FakeSSHSession) SendRequestReturns(result1 bool, result2 error) {
	fake.SendRequestStub = nil
	fake.sendRequestReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSSHSession) RequestPty(term string, h int, w int, termmodes ssh.TerminalModes) error {
	fake.requestPtyMutex.Lock()
	fake.requestPtyArgsForCall = append(fake.requestPtyArgsForCall, struct {
		term      string
		h         int
		w         int
		termmodes ssh.TerminalModes
	}{term, h, w, termmodes})
	fake.requestPtyMutex.Unlock()
	if fake.RequestPtyStub != nil {
		return fake.RequestPtyStub(term, h, w, termmodes)
	} else {
		return fake.requestPtyReturns.result1
	}
}

func (fake *FakeSSHSession) RequestPtyCallCount() int {
	fake.requestPtyMutex.RLock()
	defer fake.requestPtyMutex.RUnlock()
	return len(fake.requestPtyArgsForCall)
}

func (fake *FakeSSHSession) RequestPtyArgsForCall(i int) (string, int, int, ssh.TerminalModes) {
	fake.requestPtyMutex.RLock()
	defer fake.requestPtyMutex.RUnlock()
	return fake.requestPtyArgsForCall[i].term, fake.requestPtyArgsForCall[i].h, fake.requestPtyArgsForCall[i].w, fake.requestPtyArgsForCall[i].termmodes
}

func (fake *FakeSSHSession) RequestPtyReturns(result1 error) {
	fake.RequestPtyStub = nil
	fake.requestPtyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSHSession) Shell() error {
	fake.shellMutex.Lock()
	fake.shellArgsForCall = append(fake.shellArgsForCall, struct{}{})
	fake.shellMutex.Unlock()
	if fake.ShellStub != nil {
		return fake.ShellStub()
	} else {
		return fake.shellReturns.result1
	}
}

func (fake *FakeSSHSession) ShellCallCount() int {
	fake.shellMutex.RLock()
	defer fake.shellMutex.RUnlock()
	return len(fake.shellArgsForCall)
}

func (fake *FakeSSHSession) ShellReturns(result1 error) {
	fake.ShellStub = nil
	fake.shellReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSHSession) Run(arg1 string) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeSSHSession) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeSSHSession) RunArgsForCall(i int) string {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1
}

func (fake *FakeSSHSession) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSHSession) Wait() error {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	} else {
		return fake.waitReturns.result1
	}
}

func (fake *FakeSSHSession) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeSSHSession) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSHSession) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeSSHSession) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSSHSession) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

var _ sshapi.SSHSession = new(FakeSSHSession)
