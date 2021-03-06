// Automatically generated by MockGen. DO NOT EDIT!
// Source: va/gsb.go

package mock_gsb

import (
	gomock "github.com/golang/mock/gomock"
	safebrowsing "github.com/google/safebrowsing"
)

// Mock of SafeBrowsing interface
type MockSafeBrowsing struct {
	ctrl     *gomock.Controller
	recorder *_MockSafeBrowsingRecorder
}

// Recorder for MockSafeBrowsing (not exported)
type _MockSafeBrowsingRecorder struct {
	mock *MockSafeBrowsing
}

func NewMockSafeBrowsing(ctrl *gomock.Controller) *MockSafeBrowsing {
	mock := &MockSafeBrowsing{ctrl: ctrl}
	mock.recorder = &_MockSafeBrowsingRecorder{mock}
	return mock
}

func (_m *MockSafeBrowsing) EXPECT() *_MockSafeBrowsingRecorder {
	return _m.recorder
}

func (_m *MockSafeBrowsing) IsListed(url string) (string, error) {
	ret := _m.ctrl.Call(_m, "IsListed", url)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSafeBrowsingRecorder) IsListed(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsListed", arg0)
}

// Mock of SafeBrowsingV4 interface
type MockSafeBrowsingV4 struct {
	ctrl     *gomock.Controller
	recorder *_MockSafeBrowsingV4Recorder
}

// Recorder for MockSafeBrowsingV4 (not exported)
type _MockSafeBrowsingV4Recorder struct {
	mock *MockSafeBrowsingV4
}

func NewMockSafeBrowsingV4(ctrl *gomock.Controller) *MockSafeBrowsingV4 {
	mock := &MockSafeBrowsingV4{ctrl: ctrl}
	mock.recorder = &_MockSafeBrowsingV4Recorder{mock}
	return mock
}

func (_m *MockSafeBrowsingV4) EXPECT() *_MockSafeBrowsingV4Recorder {
	return _m.recorder
}

func (_m *MockSafeBrowsingV4) LookupURLs(urls []string) ([][]safebrowsing.URLThreat, error) {
	ret := _m.ctrl.Call(_m, "LookupURLs", urls)
	ret0, _ := ret[0].([][]safebrowsing.URLThreat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSafeBrowsingV4Recorder) LookupURLs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LookupURLs", arg0)
}
