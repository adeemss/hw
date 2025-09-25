func (c client) SetInitSystems(ctx context.Context, initSystems []blDTO.InitSystem) error {
	initSysVals := make(map[string]uint)
	for _, initSystem := range initSystems {
		converter.ParseInitSystem(initSystem, initSysVals)
	}

	value, err := json.Marshal(initSysVals)
	if err != nil {
		return fmt.Errorf("failed to prepare for set value, %w", err)
	}
	if err := c.cl.Set(ctx, c.InitSystemsDataKey, value, 0).Err(); err != nil {
		return fmt.Errorf("failed to set InitSystems directory, %w", err)
	}
	return nil
}
