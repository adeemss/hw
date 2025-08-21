"failed: invalid ID": {
			contextSetup: func() context.Context {
				return context.Background()
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {},
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
				log.On("Errorf", "failed to validate response body, %s", []interface{}{"Key: 'respBodyValidator.ID' Error:Field validation for 'ID' failed on the 'uuid' tag"})
			},
			inputSetup: func(data bDTO.Response) *bDTO.Response {
				data.ID = converter.Pointer("invalid_uuid")
				return &data
			},
			expError: fmt.Errorf("failed to validate response body, %s", "Key: 'respBodyValidator.ID' Error:Field validation for 'ID' failed on the 'uuid' tag"),
		},
Error:      	Not equal: 
        	            	expected: *errors.errorString(&errors.errorString{s:"failed to validate response body, Key: 'respBodyValidator.ID' Error:Field validation for 'ID' failed on the 'uuid' tag"})
        	            	actual  : *fmt.wrapError(&fmt.wrapError{msg:"failed to validate response body, Key: 'respBodyValidator.ID' Error:Field validation for 'ID' failed on the 'uuid' tag", err:validator.ValidationErrors{(*validator.fieldError)(0xc00014e6c0)}})
        	Test:       	TestSaveResponseBody/failed:_invalid_ID
--- FAIL: TestSaveResponseBody/failed:_invalid_ID (0.00s)

Expected :*errors.errorString(&errors.errorString{s:"failed to validate response body, Key: 'respBodyValidator.ID' Error:Field validation for 'ID' failed on the 'uuid' tag"})
Actual   :*fmt.wrapError(&fmt.wrapError{msg:"failed to validate response body, Key: 'respBodyValidator.ID' Error:Field validation for 'ID' failed on the 'uuid' tag", err:validator.ValidationErrors{(*validator.fieldError)(0xc00014e6c0)}})

как пофиксить 
