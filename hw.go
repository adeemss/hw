resp = &extproc.ProcessingResponse{
				Response: &extproc.ProcessingResponse_ImmediateResponse{
					ImmediateResponse: &extproc.ImmediateResponse{
						Status: &typev3.HttpStatus{
							Code: typev3.StatusCode(int32(httpStatusCode)),
						},
						Body: body,
					},
				},
			}
extproc "github.com/envoyproxy/go-control-plane/envoy/service/ext_proc/v3" 
как добавить сюда header content-type application/json 
