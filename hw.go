func TestRESTToBillingResource(t *testing.T) {
	outputTime := time.Now()
	testCases := map[string]struct {
		expRes billing.Resource
		input  api.ExtSystemResource
		expErr error
	}{
		"success": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing ID": {
			expRes: billing.Resource{
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing Code": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing Name": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing Method": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing Path": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing Amount": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing PriceType": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing ExtSystemID": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing CacheStoragePeriod": {
			expRes: billing.Resource{
				ID:          Pointer[uint](1),
				Code:        Pointer("test_code"),
				Name:        Pointer("test_name"),
				Path:        Pointer("test_path"),
				Method:      Pointer("POST"),
				PriceType:   Pointer("test_price_type"),
				Amount:      Pointer(23234.43),
				ExtSystemID: Pointer[uint](1),
				CreatedAt:   &outputTime,
				UpdatedAt:   &outputTime,
				IsActive:    Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:          Pointer[uint](1),
				Code:        Pointer("test_code"),
				Name:        Pointer("test_name"),
				Path:        Pointer("test_path"),
				Method:      Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId: Pointer[uint](1),
				CreatedAt:   &outputTime,
				UpdatedAt:   &outputTime,
				Active:      Pointer(true),
			},
		},
		"success: missing CreatedAt": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				UpdatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				UpdatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing UpdatedAt": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				IsActive:           Pointer(true),
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				Active:             Pointer(true),
			},
		},
		"success: missing IsActive": {
			expRes: billing.Resource{
				ID:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer("POST"),
				PriceType:          Pointer("test_price_type"),
				Amount:             Pointer(23234.43),
				ExtSystemID:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
			},
			input: api.ExtSystemResource{
				Id:                 Pointer[uint](1),
				Code:               Pointer("test_code"),
				Name:               Pointer("test_name"),
				Path:               Pointer("test_path"),
				Method:             Pointer[api.ExtSystemResourceMethod]("POST"),
				ExtSystemId:        Pointer[uint](1),
				CacheStoragePeriod: Pointer(1),
				CreatedAt:          &outputTime,
				UpdatedAt:          &outputTime,
			},
		},
		"success: empty struct": {
			expRes: billing.Resource{},
			input:  api.ExtSystemResource{},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			res, err := RESTToBillingResource(tc.input)
			if tc.expErr != nil {
				assert.Equal(t, tc.expErr, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tc.expRes, res)
		})
	}

	GOROOT=C:\Users\bakhyade\go\go1.23.5 #gosetup
GOPATH=C:\Users\bakhyade\go #gosetup
C:\Users\bakhyade\go\go1.23.5\bin\go.exe test -c -o C:\Users\bakhyade\AppData\Local\JetBrains\GoLand2025.1\tmp\GoLand\___TestRESTToBillingResource_in_gitlab_bcc_kz_digital_banking_platform_microservices_billing_dbp_ext_requests_billing_system_admin_server_converter.test.exe gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system-admin/server/converter #gosetup
C:\Users\bakhyade\go\go1.23.5\bin\go.exe tool test2json -t C:\Users\bakhyade\AppData\Local\JetBrains\GoLand2025.1\tmp\GoLand\___TestRESTToBillingResource_in_gitlab_bcc_kz_digital_banking_platform_microservices_billing_dbp_ext_requests_billing_system_admin_server_converter.test.exe -test.v=test2json -test.paniconexit0 -test.run ^\QTestRESTToBillingResource\E$ #gosetup
=== RUN   TestRESTToBillingResource
=== RUN   TestRESTToBillingResource/success:_missing_UpdatedAt
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc000222480), Code:(*string)(0xc0000257d0), Name:(*string)(0xc0000257e0), Method:(*string)(0xc000025800), Path:(*string)(0xc0000257f0), PriceType:(*string)(0xc000025810), Amount:(*float64)(0xc000222488), ExtSystemID:(*uint)(0xc000222490), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002224a0), CacheStoragePeriod:(*int)(0xc000222498), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:<nil>}
        	            	actual  : billing.Resource{ID:(*uint)(0xc0002224a8), Code:(*string)(0xc000025820), Name:(*string)(0xc000025830), Method:(*string)(0xc000025850), Path:(*string)(0xc000025840), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002224b0), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002224c0), CacheStoragePeriod:(*int)(0xc0002224b8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:<nil>}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_UpdatedAt
--- FAIL: TestRESTToBillingResource/success:_missing_UpdatedAt (0.07s)

Expected :billing.Resource{ID:(*uint)(0xc000222480), Code:(*string)(0xc0000257d0), Name:(*string)(0xc0000257e0), Method:(*string)(0xc000025800), Path:(*string)(0xc0000257f0), PriceType:(*string)(0xc000025810), Amount:(*float64)(0xc000222488), ExtSystemID:(*uin ...

Actual   :billing.Resource{ID:(*uint)(0xc0002224a8), Code:(*string)(0xc000025820), Name:(*string)(0xc000025830), Method:(*string)(0xc000025850), Path:(*string)(0xc000025840), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002224b0),  ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_IsActive
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc0002224c8), Code:(*string)(0xc000025860), Name:(*string)(0xc000025870), Method:(*string)(0xc000025890), Path:(*string)(0xc000025880), PriceType:(*string)(0xc0000258a0), Amount:(*float64)(0xc0002224d0), ExtSystemID:(*uint)(0xc0002224d8), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(nil), CacheStoragePeriod:(*int)(0xc0002224e0), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc0002224e8), Code:(*string)(0xc0000258b0), Name:(*string)(0xc0000258c0), Method:(*string)(0xc0000258e0), Path:(*string)(0xc0000258d0), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002224f0), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(nil), CacheStoragePeriod:(*int)(0xc0002224f8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_IsActive
--- FAIL: TestRESTToBillingResource/success:_missing_IsActive (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc0002224c8), Code:(*string)(0xc000025860), Name:(*string)(0xc000025870), Method:(*string)(0xc000025890), Path:(*string)(0xc000025880), PriceType:(*string)(0xc0000258a0), Amount:(*float64)(0xc0002224d0), ExtSystemID:(*uin ...

Actual   :billing.Resource{ID:(*uint)(0xc0002224e8), Code:(*string)(0xc0000258b0), Name:(*string)(0xc0000258c0), Method:(*string)(0xc0000258e0), Path:(*string)(0xc0000258d0), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002224f0),  ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_CreatedAt
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc000222438), Code:(*string)(0xc000025740), Name:(*string)(0xc000025750), Method:(*string)(0xc000025770), Path:(*string)(0xc000025760), PriceType:(*string)(0xc000025780), Amount:(*float64)(0xc000222440), ExtSystemID:(*uint)(0xc000222448), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222458), CacheStoragePeriod:(*int)(0xc000222450), CreatedAt:<nil>, UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc000222460), Code:(*string)(0xc000025790), Name:(*string)(0xc0000257a0), Method:(*string)(0xc0000257c0), Path:(*string)(0xc0000257b0), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222468), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222478), CacheStoragePeriod:(*int)(0xc000222470), CreatedAt:<nil>, UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_CreatedAt
--- FAIL: TestRESTToBillingResource/success:_missing_CreatedAt (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc000222438), Code:(*string)(0xc000025740), Name:(*string)(0xc000025750), Method:(*string)(0xc000025770), Path:(*string)(0xc000025760), PriceType:(*string)(0xc000025780), Amount:(*float64)(0xc000222440), ExtSystemID:(*uin ...

Actual   :billing.Resource{ID:(*uint)(0xc000222460), Code:(*string)(0xc000025790), Name:(*string)(0xc0000257a0), Method:(*string)(0xc0000257c0), Path:(*string)(0xc0000257b0), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222468),  ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc0002221a0), Code:(*string)(0xc000025230), Name:(*string)(0xc000025240), Method:(*string)(0xc000025260), Path:(*string)(0xc000025250), PriceType:(*string)(0xc000025270), Amount:(*float64)(0xc0002221a8), ExtSystemID:(*uint)(0xc0002221b0), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002221c0), CacheStoragePeriod:(*int)(0xc0002221b8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc0002221c8), Code:(*string)(0xc000025280), Name:(*string)(0xc000025290), Method:(*string)(0xc0000252b0), Path:(*string)(0xc0000252a0), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002221d0), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002221e0), CacheStoragePeriod:(*int)(0xc0002221d8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success
--- FAIL: TestRESTToBillingResource/success (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc0002221a0), Code:(*string)(0xc000025230), Name:(*string)(0xc000025240), Method:(*string)(0xc000025260), Path:(*string)(0xc000025250), PriceType:(*string)(0xc000025270), Amount:(*float64)(0xc0002221a8), ExtSystemID:(*uin ...

Actual   :billing.Resource{ID:(*uint)(0xc0002221c8), Code:(*string)(0xc000025280), Name:(*string)(0xc000025290), Method:(*string)(0xc0000252b0), Path:(*string)(0xc0000252a0), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002221d0),  ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_ID
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(nil), Code:(*string)(0xc0000252c0), Name:(*string)(0xc0000252d0), Method:(*string)(0xc0000252f0), Path:(*string)(0xc0000252e0), PriceType:(*string)(0xc000025300), Amount:(*float64)(0xc0002221e8), ExtSystemID:(*uint)(0xc0002221f0), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222200), CacheStoragePeriod:(*int)(0xc0002221f8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(nil), Code:(*string)(0xc000025310), Name:(*string)(0xc000025320), Method:(*string)(0xc000025340), Path:(*string)(0xc000025330), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222208), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222218), CacheStoragePeriod:(*int)(0xc000222210), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_ID
--- FAIL: TestRESTToBillingResource/success:_missing_ID (0.00s)

Expected :billing.Resource{ID:(*uint)(nil), Code:(*string)(0xc0000252c0), Name:(*string)(0xc0000252d0), Method:(*string)(0xc0000252f0), Path:(*string)(0xc0000252e0), PriceType:(*string)(0xc000025300), Amount:(*float64)(0xc0002221e8), ExtSystemID:(*uint)(0xc000 ...

Actual   :billing.Resource{ID:(*uint)(nil), Code:(*string)(0xc000025310), Name:(*string)(0xc000025320), Method:(*string)(0xc000025340), Path:(*string)(0xc000025330), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222208), Subresour ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_Name
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc000222268), Code:(*string)(0xc0000253c0), Name:(*string)(nil), Method:(*string)(0xc0000253e0), Path:(*string)(0xc0000253d0), PriceType:(*string)(0xc0000253f0), Amount:(*float64)(0xc000222270), ExtSystemID:(*uint)(0xc000222278), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222288), CacheStoragePeriod:(*int)(0xc000222280), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc000222290), Code:(*string)(0xc000025400), Name:(*string)(nil), Method:(*string)(0xc000025420), Path:(*string)(0xc000025410), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222298), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002222a8), CacheStoragePeriod:(*int)(0xc0002222a0), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_Name
--- FAIL: TestRESTToBillingResource/success:_missing_Name (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc000222268), Code:(*string)(0xc0000253c0), Name:(*string)(nil), Method:(*string)(0xc0000253e0), Path:(*string)(0xc0000253d0), PriceType:(*string)(0xc0000253f0), Amount:(*float64)(0xc000222270), ExtSystemID:(*uint)(0xc000 ...

Actual   :billing.Resource{ID:(*uint)(0xc000222290), Code:(*string)(0xc000025400), Name:(*string)(nil), Method:(*string)(0xc000025420), Path:(*string)(0xc000025410), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222298), Subresour ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_Path
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc0002222f8), Code:(*string)(0xc0000254a0), Name:(*string)(0xc0000254b0), Method:(*string)(0xc0000254c0), Path:(*string)(nil), PriceType:(*string)(0xc0000254d0), Amount:(*float64)(0xc000222300), ExtSystemID:(*uint)(0xc000222308), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222318), CacheStoragePeriod:(*int)(0xc000222310), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc000222320), Code:(*string)(0xc0000254e0), Name:(*string)(0xc0000254f0), Method:(*string)(0xc000025500), Path:(*string)(nil), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222328), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222338), CacheStoragePeriod:(*int)(0xc000222330), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)(<nil>),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_Path
--- FAIL: TestRESTToBillingResource/success:_missing_Path (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc0002222f8), Code:(*string)(0xc0000254a0), Name:(*string)(0xc0000254b0), Method:(*string)(0xc0000254c0), Path:(*string)(nil), PriceType:(*string)(0xc0000254d0), Amount:(*float64)(0xc000222300), ExtSystemID:(*uint)(0xc000 ...

Actual   :billing.Resource{ID:(*uint)(0xc000222320), Code:(*string)(0xc0000254e0), Name:(*string)(0xc0000254f0), Method:(*string)(0xc000025500), Path:(*string)(nil), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222328), Subresour ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_Amount
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc000222340), Code:(*string)(0xc000025510), Name:(*string)(0xc000025520), Method:(*string)(0xc000025540), Path:(*string)(0xc000025530), PriceType:(*string)(0xc000025550), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222348), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222358), CacheStoragePeriod:(*int)(0xc000222350), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc000222360), Code:(*string)(0xc000025560), Name:(*string)(0xc000025570), Method:(*string)(0xc000025590), Path:(*string)(0xc000025580), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222368), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222378), CacheStoragePeriod:(*int)(0xc000222370), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,3 +6,3 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	+ PriceType: (*string)(<nil>),
        	            	  Amount: (*float64)(<nil>),
        	Test:       	TestRESTToBillingResource/success:_missing_Amount
--- FAIL: TestRESTToBillingResource/success:_missing_Amount (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc000222340), Code:(*string)(0xc000025510), Name:(*string)(0xc000025520), Method:(*string)(0xc000025540), Path:(*string)(0xc000025530), PriceType:(*string)(0xc000025550), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000 ...

Actual   :billing.Resource{ID:(*uint)(0xc000222360), Code:(*string)(0xc000025560), Name:(*string)(0xc000025570), Method:(*string)(0xc000025590), Path:(*string)(0xc000025580), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222368),  ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_PriceType
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc000222380), Code:(*string)(0xc0000255a0), Name:(*string)(0xc0000255b0), Method:(*string)(0xc0000255d0), Path:(*string)(0xc0000255c0), PriceType:(*string)(nil), Amount:(*float64)(0xc000222388), ExtSystemID:(*uint)(0xc000222390), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002223a0), CacheStoragePeriod:(*int)(0xc000222398), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc0002223a8), Code:(*string)(0xc0000255e0), Name:(*string)(0xc0000255f0), Method:(*string)(0xc000025610), Path:(*string)(0xc000025600), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002223b0), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002223c0), CacheStoragePeriod:(*int)(0xc0002223b8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -7,3 +7,3 @@
        	            	  PriceType: (*string)(<nil>),
        	            	- Amount: (*float64)(23234.43),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_PriceType
--- FAIL: TestRESTToBillingResource/success:_missing_PriceType (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc000222380), Code:(*string)(0xc0000255a0), Name:(*string)(0xc0000255b0), Method:(*string)(0xc0000255d0), Path:(*string)(0xc0000255c0), PriceType:(*string)(nil), Amount:(*float64)(0xc000222388), ExtSystemID:(*uint)(0xc000 ...

Actual   :billing.Resource{ID:(*uint)(0xc0002223a8), Code:(*string)(0xc0000255e0), Name:(*string)(0xc0000255f0), Method:(*string)(0xc000025610), Path:(*string)(0xc000025600), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002223b0),  ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_CacheStoragePeriod
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc000222400), Code:(*string)(0xc0000256b0), Name:(*string)(0xc0000256c0), Method:(*string)(0xc0000256e0), Path:(*string)(0xc0000256d0), PriceType:(*string)(0xc0000256f0), Amount:(*float64)(0xc000222408), ExtSystemID:(*uint)(0xc000222410), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222418), CacheStoragePeriod:(*int)(nil), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc000222420), Code:(*string)(0xc000025700), Name:(*string)(0xc000025710), Method:(*string)(0xc000025730), Path:(*string)(0xc000025720), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222428), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222430), CacheStoragePeriod:(*int)(nil), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_CacheStoragePeriod
--- FAIL: TestRESTToBillingResource/success:_missing_CacheStoragePeriod (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc000222400), Code:(*string)(0xc0000256b0), Name:(*string)(0xc0000256c0), Method:(*string)(0xc0000256e0), Path:(*string)(0xc0000256d0), PriceType:(*string)(0xc0000256f0), Amount:(*float64)(0xc000222408), ExtSystemID:(*uin ...

Actual   :billing.Resource{ID:(*uint)(0xc000222420), Code:(*string)(0xc000025700), Name:(*string)(0xc000025710), Method:(*string)(0xc000025730), Path:(*string)(0xc000025720), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222428),  ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_empty_struct
--- PASS: TestRESTToBillingResource/success:_empty_struct (0.00s)
=== RUN   TestRESTToBillingResource/success:_missing_Code
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc000222220), Code:(*string)(nil), Name:(*string)(0xc000025350), Method:(*string)(0xc000025370), Path:(*string)(0xc000025360), PriceType:(*string)(0xc000025380), Amount:(*float64)(0xc000222228), ExtSystemID:(*uint)(0xc000222230), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222240), CacheStoragePeriod:(*int)(0xc000222238), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc000222248), Code:(*string)(nil), Name:(*string)(0xc000025390), Method:(*string)(0xc0000253b0), Path:(*string)(0xc0000253a0), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222250), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc000222260), CacheStoragePeriod:(*int)(0xc000222258), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_Code
--- FAIL: TestRESTToBillingResource/success:_missing_Code (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc000222220), Code:(*string)(nil), Name:(*string)(0xc000025350), Method:(*string)(0xc000025370), Path:(*string)(0xc000025360), PriceType:(*string)(0xc000025380), Amount:(*float64)(0xc000222228), ExtSystemID:(*uint)(0xc000 ...

Actual   :billing.Resource{ID:(*uint)(0xc000222248), Code:(*string)(nil), Name:(*string)(0xc000025390), Method:(*string)(0xc0000253b0), Path:(*string)(0xc0000253a0), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc000222250), Subresour ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_Method
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc0002222b0), Code:(*string)(0xc000025430), Name:(*string)(0xc000025440), Method:(*string)(nil), Path:(*string)(0xc000025450), PriceType:(*string)(0xc000025460), Amount:(*float64)(0xc0002222b8), ExtSystemID:(*uint)(0xc0002222c0), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002222d0), CacheStoragePeriod:(*int)(0xc0002222c8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc0002222d8), Code:(*string)(0xc000025470), Name:(*string)(0xc000025480), Method:(*string)(nil), Path:(*string)(0xc000025490), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002222e0), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002222f0), CacheStoragePeriod:(*int)(0xc0002222e8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(1),
        	Test:       	TestRESTToBillingResource/success:_missing_Method
--- FAIL: TestRESTToBillingResource/success:_missing_Method (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc0002222b0), Code:(*string)(0xc000025430), Name:(*string)(0xc000025440), Method:(*string)(nil), Path:(*string)(0xc000025450), PriceType:(*string)(0xc000025460), Amount:(*float64)(0xc0002222b8), ExtSystemID:(*uint)(0xc000 ...

Actual   :billing.Resource{ID:(*uint)(0xc0002222d8), Code:(*string)(0xc000025470), Name:(*string)(0xc000025480), Method:(*string)(nil), Path:(*string)(0xc000025490), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(0xc0002222e0), Subresour ...

<Click to see difference>


=== RUN   TestRESTToBillingResource/success:_missing_ExtSystemID
    rest_test.go:1502: 
        	Error Trace:	C:/Users/bakhyade/Desktop/projects/mr conf/dbp-ext-requests-billing-system-admin/server/converter/rest_test.go:1502
        	Error:      	Not equal: 
        	            	expected: billing.Resource{ID:(*uint)(0xc0002223c8), Code:(*string)(0xc000025620), Name:(*string)(0xc000025630), Method:(*string)(0xc000025650), Path:(*string)(0xc000025640), PriceType:(*string)(0xc000025660), Amount:(*float64)(0xc0002223d0), ExtSystemID:(*uint)(nil), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002223e0), CacheStoragePeriod:(*int)(0xc0002223d8), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	actual  : billing.Resource{ID:(*uint)(0xc0002223e8), Code:(*string)(0xc000025670), Name:(*string)(0xc000025680), Method:(*string)(0xc0000256a0), Path:(*string)(0xc000025690), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(nil), Subresources:[]billing.Subresource(nil), IsActive:(*bool)(0xc0002223f8), CacheStoragePeriod:(*int)(0xc0002223f0), CreatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local), UpdatedAt:time.Date(2025, time.August, 25, 20, 34, 47, 825774300, time.Local)}
        	            	
        	            	Diff:
        	            	--- Expected
        	            	+++ Actual
        	            	@@ -6,4 +6,4 @@
        	            	  Path: (*string)((len=9) "test_path"),
        	            	- PriceType: (*string)((len=15) "test_price_type"),
        	            	- Amount: (*float64)(23234.43),
        	            	+ PriceType: (*string)(<nil>),
        	            	+ Amount: (*float64)(<nil>),
        	            	  ExtSystemID: (*uint)(<nil>),
        	Test:       	TestRESTToBillingResource/success:_missing_ExtSystemID
--- FAIL: TestRESTToBillingResource/success:_missing_ExtSystemID (0.00s)

Expected :billing.Resource{ID:(*uint)(0xc0002223c8), Code:(*string)(0xc000025620), Name:(*string)(0xc000025630), Method:(*string)(0xc000025650), Path:(*string)(0xc000025640), PriceType:(*string)(0xc000025660), Amount:(*float64)(0xc0002223d0), ExtSystemID:(*uin ...

Actual   :billing.Resource{ID:(*uint)(0xc0002223e8), Code:(*string)(0xc000025670), Name:(*string)(0xc000025680), Method:(*string)(0xc0000256a0), Path:(*string)(0xc000025690), PriceType:(*string)(nil), Amount:(*float64)(nil), ExtSystemID:(*uint)(nil), Subresour ...

<Click to see difference>


--- FAIL: TestRESTToBillingResource (0.08s)

FAIL

Process finished with the exit code 1


}
