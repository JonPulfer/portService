package service

import (
	"testing"

	"github.com/JonPulfer/portsService/internal/portService"
	"github.com/JonPulfer/portsService/internal/portService/repository/mocks"
	"github.com/golang/mock/gomock"
)

func TestPortService(t *testing.T) {
	ctl := gomock.NewController(t)

	type testCase struct {
		name    string
		port    portService.Port
		portSvc Port
		err     error
		init    func(*testCase)
		assert  func(testCase)
	}

	// I would use the mock in tests like so.
	testCases := []testCase{
		{
			name: "store ok",
			init: func(tc *testCase) {
				mockRepos := mocks.NewMockPortRepository(ctl)
				tc.portSvc = NewPortService(mockRepos)

				mockRepos.EXPECT().Store(gomock.Any()).Times(1)
			},
		},
	}

	for _, tc := range testCases {
		tc.init(&tc)
		tc.err = tc.portSvc.Store(tc.port)
		tc.assert(tc)
		ctl.Finish()
	}

}
