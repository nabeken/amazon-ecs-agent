// Copyright 2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/acs/update_handler/os (interfaces: FileSystem)

package mock_os

import (
	io "io"
	os "os"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of FileSystem interface
type MockFileSystem struct {
	ctrl     *gomock.Controller
	recorder *_MockFileSystemRecorder
}

// Recorder for MockFileSystem (not exported)
type _MockFileSystemRecorder struct {
	mock *MockFileSystem
}

func NewMockFileSystem(ctrl *gomock.Controller) *MockFileSystem {
	mock := &MockFileSystem{ctrl: ctrl}
	mock.recorder = &_MockFileSystemRecorder{mock}
	return mock
}

func (_m *MockFileSystem) EXPECT() *_MockFileSystemRecorder {
	return _m.recorder
}

func (_m *MockFileSystem) Copy(_param0 io.Writer, _param1 io.Reader) (int64, error) {
	ret := _m.ctrl.Call(_m, "Copy", _param0, _param1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemRecorder) Copy(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Copy", arg0, arg1)
}

func (_m *MockFileSystem) Create(_param0 string) (io.ReadWriteCloser, error) {
	ret := _m.ctrl.Call(_m, "Create", _param0)
	ret0, _ := ret[0].(io.ReadWriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemRecorder) Create(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create", arg0)
}

func (_m *MockFileSystem) Exit(_param0 int) {
	_m.ctrl.Call(_m, "Exit", _param0)
}

func (_mr *_MockFileSystemRecorder) Exit(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Exit", arg0)
}

func (_m *MockFileSystem) MkdirAll(_param0 string, _param1 os.FileMode) error {
	ret := _m.ctrl.Call(_m, "MkdirAll", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFileSystemRecorder) MkdirAll(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MkdirAll", arg0, arg1)
}

func (_m *MockFileSystem) Open(_param0 string) (io.ReadWriteCloser, error) {
	ret := _m.ctrl.Call(_m, "Open", _param0)
	ret0, _ := ret[0].(io.ReadWriteCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemRecorder) Open(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Open", arg0)
}

func (_m *MockFileSystem) ReadAll(_param0 io.Reader) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "ReadAll", _param0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemRecorder) ReadAll(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadAll", arg0)
}

func (_m *MockFileSystem) Remove(_param0 string) {
	_m.ctrl.Call(_m, "Remove", _param0)
}

func (_mr *_MockFileSystemRecorder) Remove(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Remove", arg0)
}

func (_m *MockFileSystem) Rename(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "Rename", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFileSystemRecorder) Rename(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Rename", arg0, arg1)
}

func (_m *MockFileSystem) TeeReader(_param0 io.Reader, _param1 io.Writer) io.Reader {
	ret := _m.ctrl.Call(_m, "TeeReader", _param0, _param1)
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

func (_mr *_MockFileSystemRecorder) TeeReader(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TeeReader", arg0, arg1)
}

func (_m *MockFileSystem) TempFile(_param0 string, _param1 string) (*os.File, error) {
	ret := _m.ctrl.Call(_m, "TempFile", _param0, _param1)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockFileSystemRecorder) TempFile(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TempFile", arg0, arg1)
}

func (_m *MockFileSystem) WriteFile(_param0 string, _param1 []byte, _param2 os.FileMode) error {
	ret := _m.ctrl.Call(_m, "WriteFile", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockFileSystemRecorder) WriteFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "WriteFile", arg0, arg1, arg2)
}
