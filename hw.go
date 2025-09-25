type client struct {
	cl                 *redis.ClusterClient
	InitSystemsDataKey string
	AttributesDataKey  string
	ExtSystemsDataKey  string
	ResponseKeyPrefix  string
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

	return &client{
		cl:                 cli,
		InitSystemsDataKey: InitSystemDataKey(cfg.Redis.KeyPrefix),
		AttributesDataKey:  AttributeDataKey(cfg.Redis.KeyPrefix),
		ExtSystemsDataKey:  ExtSystemsDataKey(cfg.Redis.KeyPrefix),
		ResponseKeyPrefix:  cfg.Redis.KeyPrefix,
	}, nil
}

func (c client) Close() error {
	if err := c.cl.Close(); err != nil {
		return fmt.Errorf("failed to close redis connection, %w", err)
	}
	return nil
}

func (c client) GetInitSystems(ctx context.Context) (map[string]uint, error) {
	var res map[string]uint
	value, err := c.cl.Get(ctx, c.InitSystemsDataKey).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get value by key, %w", err)
	}
	if err := json.Unmarshal([]byte(value), &res); err != nil {
		return nil, fmt.Errorf("failed to prepare received value, %w", err)
	}
	return res, nil
}
