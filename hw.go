func TestCreateResource(t *testing.T) {
	outputTime := time.Now()
	var inputResource billing.Resource
	outputResource := &billing.Resource{
		ID:                 converter.Pointer[uint](1),
		Code:               converter.Pointer("test_code"),
		Name:               converter.Pointer("test_name"),
		Method:             converter.Pointer("POST"),
		ExtSystemID:        converter.Pointer[uint](1),
		CacheStoragePeriod: converter.Pointer(1),
		CreatedAt:          outputTime,
		UpdatedAt:          outputTime,
		IsActive:           converter.Pointer(true),
	}
	testCases := map[string]struct {
		inputSetup         func(ins billing.Resource) *billing.Resource
		storeProviderSetup func(ctx context.Context, sp *sMock.StoreServiceMock)
		loggerSetup        func(ctx context.Context, log *lMock.MockContextLogger)
		expError           error
	}{
		"success": {
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
			},
			inputSetup: func(ins billing.Resource) *billing.Resource {
				return &ins
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {
				sp.On("SaveResource", ctx, mock.MatchedBy(func(ins *billing.Resource) bool {
					ins.ID = converter.Pointer[uint](1)
					ins.CreatedAt = outputTime
					ins.UpdatedAt = outputTime
					ins.IsActive = converter.Pointer(true)
					return assert.Equal(t, inputResource.Code, ins.Code) &&
						assert.Equal(t, inputResource.Name, ins.Name) &&
						assert.Equal(t, inputResource.Method, ins.Method) &&
						assert.Equal(t, inputResource.Path, ins.Path) &&
						assert.Equal(t, inputResource.ExtSystemID, ins.ExtSystemID) &&
						assert.Equal(t, inputResource.CacheStoragePeriod, ins.CacheStoragePeriod)
				})).Return(nil)
			},
			expError: nil,
		},
		"failed: missing Code": {
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
				log.On("Errorf", "failed to validate Resource, %s", []interface{}{"Key: 'resourceValidator.Code' Error:Field validation for 'Code' failed on the 'required' tag"})
			},
			inputSetup: func(ins billing.Resource) *billing.Resource {
				ins.Code = nil
				return &ins
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {},
			expError:           fmt.Errorf("failed to validate Resource, %s", "Key: 'resourceValidator.Code' Error:Field validation for 'Code' failed on the 'required' tag"),
		},
		"failed: missing Name": {
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
				log.On("Errorf", "failed to validate Resource, %s", []interface{}{"Key: 'resourceValidator.Name' Error:Field validation for 'Name' failed on the 'required' tag"})
			},
			inputSetup: func(ins billing.Resource) *billing.Resource {
				ins.Name = nil
				return &ins
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {},
			expError:           fmt.Errorf("failed to validate Resource, %s", "Key: 'resourceValidator.Name' Error:Field validation for 'Name' failed on the 'required' tag"),
		},
		"failed: missing Method": {
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
				log.On("Errorf", "failed to validate Resource, %s", []interface{}{"Key: 'resourceValidator.Method' Error:Field validation for 'Method' failed on the 'required' tag"})
			},
			inputSetup: func(ins billing.Resource) *billing.Resource {
				ins.Method = nil
				return &ins
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {},
			expError:           fmt.Errorf("failed to validate Resource, %s", "Key: 'resourceValidator.Method' Error:Field validation for 'Method' failed on the 'required' tag"),
		},
		"failed: missing Path": {
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
				log.On("Errorf", "failed to validate Resource, %s", []interface{}{"Key: 'resourceValidator.Path' Error:Field validation for 'Path' failed on the 'required' tag"})
			},
			inputSetup: func(ins billing.Resource) *billing.Resource {
				ins.Path = nil
				return &ins
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {},
			expError:           fmt.Errorf("failed to validate Resource, %s", "Key: 'resourceValidator.Path' Error:Field validation for 'Path' failed on the 'required' tag"),
		},
		"failed: missing CacheStoragePeriod": {
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
				log.On("Errorf", "failed to validate Resource, %s", []interface{}{"Key: 'resourceValidator.CacheStoragePeriod' Error:Field validation for 'CacheStoragePeriod' failed on the 'required' tag"})
			},
			inputSetup: func(ins billing.Resource) *billing.Resource {
				ins.CacheStoragePeriod = nil
				return &ins
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {
			},
			expError: fmt.Errorf("failed to validate Resource, %s", "Key: 'resourceValidator.CacheStoragePeriod' Error:Field validation for 'CacheStoragePeriod' failed on the 'required' tag"),
		},
		"failed: missing ExtSystemID": {
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
				log.On("Errorf", "failed to validate Resource, %s", []interface{}{"Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"})
			},
			inputSetup: func(ins billing.Resource) *billing.Resource {
				ins.ExtSystemID = nil
				return &ins
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {
			},
			expError: fmt.Errorf("failed to validate Resource, %s", "Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"),
		},
		"failed: SaveResource": {
			loggerSetup: func(ctx context.Context, log *lMock.MockContextLogger) {
				log.On("New", ctx).Return(log)
				log.On("Errorf", "failed to save resource, %s", []interface{}{assert.AnError.Error()})
			},
			inputSetup: func(ins billing.Resource) *billing.Resource {
				return &ins
			},
			storeProviderSetup: func(ctx context.Context, sp *sMock.StoreServiceMock) {
				sp.On("SaveResource", ctx, mock.MatchedBy(func(ins *billing.Resource) bool {
					return assert.Equal(t, inputResource, *ins)
				})).Return(assert.AnError)
			},
			expError: fmt.Errorf("failed to save resource, %w", assert.AnError),
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			inputResource = billing.Resource{
				Code:               converter.Pointer("test_code"),
				Name:               converter.Pointer("test_name"),
				Path:               converter.Pointer("test_path"),
				Method:             converter.Pointer("POST"),
				ExtSystemID:        converter.Pointer[uint](1),
				CacheStoragePeriod: converter.Pointer(1),
			}
			ctx := context.Background()
			log := &lMock.MockContextLogger{}
			sp := sMock.NewStoreServiceMock(t)
			tc.loggerSetup(ctx, log)
			tc.storeProviderSetup(ctx, sp)

			ins := tc.inputSetup(inputResource)

			bp := BillingProvider{
				StoreProvider: sp,
				Logger:        log,
			}
			err := bp.CreateResource(ctx, ins)
			if tc.expError != nil {
				assert.Equal(t, tc.expError.Error(), err.Error())
			} else {
				assert.Equal(t, outputResource, ins)
			}

			mock.AssertExpectationsForObjects(t, sp, log)
		})
	}
GOROOT=C:\Users\bakhyade\go\go1.23.5 #gosetup
GOPATH=C:\Users\bakhyade\go #gosetup
C:\Users\bakhyade\go\go1.23.5\bin\go.exe test -c -o C:\Users\bakhyade\AppData\Local\JetBrains\GoLand2025.1\tmp\GoLand\___TestCreateResource_in_gitlab_bcc_kz_digital_banking_platform_microservices_billing_dbp_ext_requests_billing_system_server_logic.test.exe gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/server/logic #gosetup
C:\Users\bakhyade\go\go1.23.5\bin\go.exe tool test2json -t C:\Users\bakhyade\AppData\Local\JetBrains\GoLand2025.1\tmp\GoLand\___TestCreateResource_in_gitlab_bcc_kz_digital_banking_platform_microservices_billing_dbp_ext_requests_billing_system_server_logic.test.exe -test.v=test2json -test.paniconexit0 -test.run ^\QTestCreateResource\E$ #gosetup
=== RUN   TestCreateResource
=== RUN   TestCreateResource/failed:_missing_ExtSystemID
--- FAIL: TestCreateResource/failed:_missing_ExtSystemID (0.00s)

--- FAIL: TestCreateResource (0.00s)
panic: 
	
	mock: Unexpected Method Call
	-----------------------------
	
	Errorf(string,[]interface {})
			0: "failed to validate Resource, %s"
			1: []interface {}{"Key: 'resourceValidator.PriceType' Error:Field validation for 'PriceType' failed on the 'required' tag\nKey: 'resourceValidator.Amount' Error:Field validation for 'Amount' failed on the 'required' tag\nKey: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"}
	
	The closest call I have is: 
	
	Errorf(string,[]interface {})
			0: "failed to validate Resource, %s"
			1: []interface {}{"Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"}
	
	Difference found in argument 1:
	
	--- Expected
	+++ Actual
	@@ -1,3 +1,3 @@
	 ([]interface {}) (len=1) {
	- (string) (len=106) "Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"
	+ (string) (len=306) "Key: 'resourceValidator.PriceType' Error:Field validation for 'PriceType' failed on the 'required' tag\nKey: 'resourceValidator.Amount' Error:Field validation for 'Amount' failed on the 'required' tag\nKey: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"
	 }
	
	Diff: 0: PASS:  (string=failed to validate Resource, %s) == (string=failed to validate Resource, %s)
		1: FAIL:  ([]interface {}=[Key: 'resourceValidator.PriceType' Error:Field validation for 'PriceType' failed on the 'required' tag
	Key: 'resourceValidator.Amount' Error:Field validation for 'Amount' failed on the 'required' tag
	Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag]) != ([]interface {}=[Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag])
	at: [C:/Users/bakhyade/Desktop/projects/billing/dbp-ext-requests-billing-system/server/logic/manager.go:199 C:/Users/bakhyade/Desktop/projects/billing/dbp-ext-requests-billing-system/server/logic/manager_test.go:1019]
	 [recovered]
	panic: 
	
	mock: Unexpected Method Call
	-----------------------------
	
	Errorf(string,[]interface {})
			0: "failed to validate Resource, %s"
			1: []interface {}{"Key: 'resourceValidator.PriceType' Error:Field validation for 'PriceType' failed on the 'required' tag\nKey: 'resourceValidator.Amount' Error:Field validation for 'Amount' failed on the 'required' tag\nKey: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"}
	
	The closest call I have is: 
	
	Errorf(string,[]interface {})
			0: "failed to validate Resource, %s"
			1: []interface {}{"Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"}
	
	Difference found in argument 1:
	
	--- Expected
	+++ Actual
	@@ -1,3 +1,3 @@
	 ([]interface {}) (len=1) {
	- (string) (len=106) "Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"
	+ (string) (len=306) "Key: 'resourceValidator.PriceType' Error:Field validation for 'PriceType' failed on the 'required' tag\nKey: 'resourceValidator.Amount' Error:Field validation for 'Amount' failed on the 'required' tag\nKey: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag"
	 }
	
	Diff: 0: PASS:  (string=failed to validate Resource, %s) == (string=failed to validate Resource, %s)
		1: FAIL:  ([]interface {}=[Key: 'resourceValidator.PriceType' Error:Field validation for 'PriceType' failed on the 'required' tag
	Key: 'resourceValidator.Amount' Error:Field validation for 'Amount' failed on the 'required' tag
	Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag]) != ([]interface {}=[Key: 'resourceValidator.ExtSystemID' Error:Field validation for 'ExtSystemID' failed on the 'required' tag])
	at: [C:/Users/bakhyade/Desktop/projects/billing/dbp-ext-requests-billing-system/server/logic/manager.go:199 C:/Users/bakhyade/Desktop/projects/billing/dbp-ext-requests-billing-system/server/logic/manager_test.go:1019]
	

goroutine 8 [running]:
testing.tRunner.func1.2({0x1002160, 0xc000236be0})
	C:/Users/bakhyade/go/go1.23.5/src/testing/testing.go:1632 +0x225
testing.tRunner.func1()
	C:/Users/bakhyade/go/go1.23.5/src/testing/testing.go:1635 +0x359
panic({0x1002160?, 0xc000236be0?})
	C:/Users/bakhyade/go/go1.23.5/src/runtime/panic.go:785 +0x132
github.com/stretchr/testify/mock.(*Mock).fail(0xc00014e960, {0x11304a0?, 0x2?}, {0xc00014edc0?, 0x2?, 0x2?})
	C:/Users/bakhyade/go/pkg/mod/github.com/stretchr/testify@v1.10.0/mock/mock.go:349 +0x12d
github.com/stretchr/testify/mock.(*Mock).MethodCalled(0xc00014e960, {0x1317425, 0x6}, {0xc00005bf20, 0x2, 0x2})
	C:/Users/bakhyade/go/pkg/mod/github.com/stretchr/testify@v1.10.0/mock/mock.go:509 +0x5d7
github.com/stretchr/testify/mock.(*Mock).Called(0xc00014e960, {0xc00005bf20, 0x2, 0x2})
	C:/Users/bakhyade/go/pkg/mod/github.com/stretchr/testify@v1.10.0/mock/mock.go:481 +0x125
gitlab.bcc.kz/digital-banking-platform/packages/go/logs/context-logger/ctxlogger/mock.(*MockContextLogger).Errorf(0xc00014e960, {0x1117a44, 0x1f}, {0xc0002368f0, 0x1, 0x1})
	C:/Users/bakhyade/go/pkg/mod/gitlab.bcc.kz/digital-banking-platform/packages/go/logs/context-logger.git@v0.1.0/ctxlogger/mock/mock.go:66 +0xd1
gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/server/logic.BillingProvider.CreateResource({{0x0, 0x0}, {0x1209db0, 0xc00014e9b0}, {0x0, 0x0}, {0x12071e8, 0xc00014e960}}, {0x11feca0, 0x16d2860}, ...)
	C:/Users/bakhyade/Desktop/projects/billing/dbp-ext-requests-billing-system/server/logic/manager.go:199 +0x158
gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/server/logic.TestCreateResource.func25(0xc000184ea0)
	C:/Users/bakhyade/Desktop/projects/billing/dbp-ext-requests-billing-system/server/logic/manager_test.go:1019 +0x379
testing.tRunner(0xc000184ea0, 0xc000069480)
	C:/Users/bakhyade/go/go1.23.5/src/testing/testing.go:1690 +0xcb
created by testing.(*T).Run in goroutine 7
	C:/Users/bakhyade/go/go1.23.5/src/testing/testing.go:1743 +0x377


Process finished with the exit code 1


}
