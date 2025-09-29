package redis

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redsync/redsync/v4"
	grd "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/worker/config"
	wrErrs "gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/worker/errors"
	"gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/worker/store/client/redis/constants"
)

type Mutex interface {
	UnlockContext(ctx context.Context) (bool, error)
	LockContext(ctx context.Context) error
}

type MutexClient struct {
	cl *redsync.Mutex
}

func (cl *MutexClient) UnlockContext(ctx context.Context) (bool, error) {
	return cl.cl.UnlockContext(ctx)
}
func (cl *MutexClient) LockContext(ctx context.Context) error {
	return cl.cl.LockContext(ctx)
}

type client struct {
	cl    *redis.ClusterClient
	rs    *redsync.Redsync
	mutex Mutex
}

func NewClient(ctx context.Context, cfg config.Config) (*client, error) {
	addrs := strings.Split(cfg.Redis.Addresses, ";")
	cli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: cfg.Redis.Password,
	})

	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis server: %w", err)
	}

	pool := grd.NewPool(cli)
	if pool == nil {
		return nil, errors.New("empty pool returned")
	}
	rs := redsync.New(pool)

	if cfg.Redis.DefaultLockTTL <= 0 {
		return nil, errors.New("ttl must be > 0")
	}

	m := rs.NewMutex(
		LockerKey(constants.LockerKey),
		redsync.WithExpiry(cfg.Redis.DefaultLockTTL),
		redsync.WithTries(cfg.Redis.LockAttemptsCount),
		redsync.WithRetryDelay(cfg.Redis.LockRetryDelay),
	)

	return &client{
		cl: cli,
		rs: rs,
		mutex: &MutexClient{
			cl: m,
		},
	}, nil
}

func (c client) Close() error {
	if err := c.cl.Close(); err != nil {
		return fmt.Errorf("failed to close redis connection, %w", err)
	}
	return nil
}

func (c client) GetCleanOperLogExecTime(ctx context.Context) (time.Time, bool, error) {
	s, err := c.cl.Get(ctx, ExecTimeKey(constants.ExecTimeKey)).Int64()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return time.Time{}, false, nil
		}
		return time.Time{}, false, err
	}

	return time.Unix(s, 0).UTC(), true, nil
}

func (c client) SetCleanOperLogExecTime(ctx context.Context, runAt time.Time) error {
	if err := c.cl.Set(ctx, ExecTimeKey(constants.ExecTimeKey), runAt.UTC().Unix(), 0).Err(); err != nil {
		return err
	}
	return nil
}

func (c *client) Lock(ctx context.Context) error {
	if err := c.mutex.LockContext(ctx); err != nil {
		return fmt.Errorf("failed to lock context, %w, %w", err, wrErrs.BusyLock)
	}
	return nil
}

func (c *client) Unlock(ctx context.Context) error {
	if ok, err := c.mutex.UnlockContext(ctx); !ok || err != nil {
		_ = err
	}
	return nil
}
