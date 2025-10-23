package driver

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/ibiteam/go-core/filesystem/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	config config.MinioConfig
}

func NewMinio(cfg *config.MinioConfig) (*Minio, error) {
	if cfg == nil {
		return nil, errors.New("Minio config is required")
	}
	if cfg.AccessKey == "" || cfg.AccessSecret == "" {
		return nil, errors.New("Minio access key and secret are required")
	}
	return &Minio{config: *cfg}, nil
}

func (m Minio) PutFile(file *multipart.FileHeader, dirPath string, filename string) (string, error) {
	// 初始化 MinIO 客户端
	minioClient, err := minio.New(m.config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(m.config.AccessKey, m.config.AccessSecret, ""),
		Secure: false, // 不使用 HTTPS
	})
	if err != nil {
		return "", errors.New("MinIO 连接失败: " + err.Error())
	}

	// 打开文件流
	fileStream, openErr := file.Open()
	if openErr != nil {
		return "", errors.New("打开文件失败: " + openErr.Error())
	}
	defer func(fileStream multipart.File) {
		_ = fileStream.Close()
	}(fileStream)

	// 构造完整的文件路径
	newFilePath := fmt.Sprintf("%s/%s", strings.TrimRight(dirPath, "/"), filename)

	// 上传到 MinIO
	_, err = minioClient.PutObject(
		context.Background(),
		m.config.Bucket,
		newFilePath,
		fileStream,
		file.Size,
		minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")},
	)
	if err != nil {
		return "", errors.New("上传失败: " + err.Error())
	}

	// 如果配置了自定义域名
	if m.config.Domain != "" {
		return fmt.Sprintf("http://%s/%s/%s", m.config.Domain, m.config.Bucket, newFilePath), nil
	}

	// 使用默认的 MinIO 地址
	return fmt.Sprintf("http://%s/%s/%s", m.config.Endpoint, m.config.Bucket, newFilePath), nil
}
