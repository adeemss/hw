package minio

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	min "gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/server/store/client/minio/mock"
	"strconv"
	"testing"
)

func TestSaveBody(t *testing.T) {
	inpuFileName := "testdata/billing.json"
	ctx := context.Background()
	bucket := "bucket"
	retention := "retention"

	body := "test-body"
	pl := Payload{
		Body: body,
	}

	result := minio.UploadInfo{
		Bucket: bucket,
	}
	testCases := map[string]struct {
		inputFileName       string
		inputRetenionPeriod int
		MinioSetup          func(t *testing.T, bs *min.MinIOClientServicesMock)
		ExpErr              error
	}{
		"success": {
			inputFileName:       inpuFileName,
			inputRetenionPeriod: 10,
			MinioSetup: func(t *testing.T, bs *min.MinIOClientServicesMock) {
				expJson, _ := json.Marshal(pl)
				reader := bytes.NewReader(expJson)
				bs.On("PutObject", ctx,
					mock.MatchedBy(func(bucketName string) bool {
						return bucketName == bucket
					}), mock.MatchedBy(func(FileName string) bool {
						return inpuFileName == FileName
					}),
					reader,
					int64(len(expJson)),
					minio.PutObjectOptions{ContentType: "application/json", UserTags: map[string]string{
						retention: strconv.Itoa(10),
					}}).Return(result, nil)
			},
		},
		"failed": {
			inputFileName:       inpuFileName,
			inputRetenionPeriod: 10,
			MinioSetup: func(t *testing.T, bs *min.MinIOClientServicesMock) {
				expJson, _ := json.Marshal(pl)
				reader := bytes.NewReader(expJson)
				bs.On("PutObject", ctx,
					mock.MatchedBy(func(bucketName string) bool {
						return bucketName == bucket
					}), mock.MatchedBy(func(FileName string) bool {
						return inpuFileName == FileName
					}),
					reader,
					int64(len(expJson)),
					minio.PutObjectOptions{ContentType: "application/json", UserTags: map[string]string{
						retention: strconv.Itoa(10),
					}}).Return(minio.UploadInfo{}, assert.AnError)
			},
			ExpErr: assert.AnError,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			mockMinio := min.NewMinIOClientServicesMock(t)
			client := client{
				cl:           mockMinio,
				bucketName:   bucket,
				retentionTag: retention,
			}
			tc.MinioSetup(t, mockMinio)
			err := client.SaveBody(ctx, tc.inputFileName, body, tc.inputRetenionPeriod)
			if tc.ExpErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			mockMinio.AssertExpectations(t)
		})
	}
}
