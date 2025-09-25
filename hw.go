package minio

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gitlab.bcc.kz/digital-banking-platform/microservices/billing/dbp-ext-requests-billing-system/server/config"
)

type client struct {
	cl           *minio.Client
	bucketName   string
	retentionTag string
}

func NewClient(cfg config.Config) (*client, error) {
	cl, err := minio.New(cfg.MinIO.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinIO.AccessKeyID, cfg.MinIO.SecretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		return nil, err
	}

	return &client{
		cl:           cl,
		bucketName:   cfg.MinIO.RequestsBucketName,
		retentionTag: cfg.MinIO.RetentionTag,
	}, nil
}

func (c client) SaveBody(ctx context.Context, fileName, body string, retentionPeriod int) error {
	pl := Payload{
		Body: body,
	}
	b, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(b)
	retention := strconv.Itoa(retentionPeriod)

	_, err = c.cl.PutObject(
		ctx,
		c.bucketName,
		fileName,
		reader,
		int64(len(b)),
		minio.PutObjectOptions{ContentType: "application/json", UserTags: map[string]string{
			c.retentionTag: retention,
		}},
	)
	if err != nil {
		return err
	}

	return nil
}

func (c client) GetBody(ctx context.Context, fileName string) (*string, error) {
	pl := Payload{}
	obj, err := c.cl.GetObject(ctx, c.bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	defer obj.Close()

	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &pl); err != nil {
		return nil, err
	}
	return &pl.Body, nil
}
