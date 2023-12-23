package miniohelper

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"time"
)

type MinIOClient struct {
	Client *minio.Client
}

func New(endpoint, accessKey, secretKey string, useSSL bool) (*MinIOClient, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &MinIOClient{Client: minioClient}, nil
}

func (c *MinIOClient) UploadFile(bucketName, objectName, filePath string) (minio.UploadInfo, error) {
	uInfo, err := c.Client.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
	return uInfo, err
}

func (c *MinIOClient) DownloadFile(bucketName, objectName, filePath string) error {
	return c.Client.FGetObject(context.Background(), bucketName, objectName, filePath, minio.GetObjectOptions{})
}

func (c *MinIOClient) DeleteFile(bucketName, objectName string) error {
	return c.Client.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
}

func (c *MinIOClient) BucketExists(bucketName string) (bool, error) {
	return c.Client.BucketExists(context.Background(), bucketName)
}

func (c *MinIOClient) MakeBucket(bucketName string) error {
	return c.Client.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
}

func (c *MinIOClient) RemoveBucket(bucketName string) error {
	return c.Client.RemoveBucket(context.Background(), bucketName)
}

func (c *MinIOClient) ListObjects(bucketName string) <-chan minio.ObjectInfo {
	return c.Client.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{})
}

func (c *MinIOClient) GetDownloadURL(bucketName, objectName string, expires time.Duration, reqParams url.Values) (*url.URL, error) {
	return c.Client.PresignedGetObject(context.Background(), bucketName, objectName, expires, reqParams)
}

func (c *MinIOClient) Ping() error {
	ok := c.Client.IsOnline()
	if !ok {
		return fmt.Errorf("minio client is not online")
	}

	return nil
}
