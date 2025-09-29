package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redismock/v9"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/stretchr/testify/require"
	"gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/worker/store/client/redis/constants"
	mMock "gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/worker/store/client/redis/mock"
	"strconv"
	"testing"
	"time"
)

func NewMockClient() (*client, redismock.ClusterClientMock) {
	clusterMock, mock := redismock.NewClusterMock()
	pool := goredis.NewPool(clusterMock)
	rs := redsync.New(pool)
	mutex := mMock.MutexMock{}

	return &client{
		cl:    clusterMock,
		rs:    rs,
		mutex: &mutex,
	}, mock

}
func TestLock(t *testing.T) {
	testCases := map[string]struct {
		setup    func(mock mMock.MutexMock)
		expError error
	}{
		"success": {
			setup: func(mock mMock.MutexMock) {
				mock.On("LockContext", context.Background()).Return(true)
			},
			expError: nil,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			client, mock := NewMockClient()
			mutex := mMock.NewMutexMock(t)
			tc.setup(*mutex)
			err := client.Lock(ctx)
			if tc.expError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
