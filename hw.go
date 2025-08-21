func (p *Provider) GetCacheStoragePeriodByResponseID(ctx context.Context, respID string) (*int, error) {
	var err error
	var cachePeriod *int
	log := p.Logger.New(ctx)

	resp, err := p.DBClient.GetResponseByID(respID)
	if err != nil {
		if err == srvErrs.ErrNotFound {
			return nil, err
		}
		log.Errorf("failed to get request by response ID from DB, %s", err.Error())
		return nil, fmt.Errorf("failed to get request by response ID from DB, %w", err)
	}
	cachePeriod, err = p.GetCacheStoragePeriodByRequestID(ctx, *resp.RequestID)
	if err != nil {
		log.Errorf("failed to get cache storage period by request ID from DB, %s", err.Error())
		return nil, fmt.Errorf("failed to get cache storage period by request ID from DB, %w", err)
	}
	return cachePeriod, nil
}
func TestGetCacheStoragePeriodByResponseID(t *testing.T) {
	respID := "7b0e0d62-697c-4e57-88fb-884b91bcd4d6"
	expectedResponse := &bDTO.Response{
		RequestID: converter.Pointer("3759c05e-1596-4c3e-91e7-8db82ad20239"),
	}
	testCases := map[string]struct {
		contextSetup  func() context.Context
		loggerSetup   func(ctx context.Context, log *lMock.MockContextLogger)
		dbClientSetup func(cl *cMock.DBClientMock)
		ExpRes        *int
		expError      error
	}{
		"success": {
			contextSetup: func() context.Context {
				return context.Background()
			},
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
			},

			dbClientSetup: func(cl *cMock.DBClientMock) {
				cl.On("GetResponseByID", mock.MatchedBy(func(responseID string) bool {
					return assert.Equal(t, respID, responseID)
				})).Return(expectedResponse, nil)
			},
			ExpRes: converter.Pointer(1),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			log := &lMock.MockContextLogger{}
			cl := cMock.NewDBClientMock(t)
			ctx := tc.contextSetup()
			tc.loggerSetup(ctx, log)
			tc.dbClientSetup(cl)

			sp := Provider{
				Logger:   log,
				DBClient: cl,
			}

			resp, err := sp.GetCacheStoragePeriodByResponseID(ctx, respID)
			if err != nil {
				assert.Equal(t, tc.expError, err)
			}

			assert.Equal(t, resp, tc.ExpRes)

			mock.AssertExpectationsForObjects(t, cl, log)
		})
	}
}
