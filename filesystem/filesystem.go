package filesystem

import (
	"errors"
	"mime/multipart"
	"strings"

	"github.com/ibiteam/go-core/filesystem/config"
	"github.com/ibiteam/go-core/filesystem/driver"
)

type FileDriver interface {
	PutFile(file *multipart.FileHeader, dirPath string, filename string) (string, error)
}

// Factory 存储工厂，用于按需创建驱动实例
type Factory struct {
	config config.Config
}

// NewFactory 创建存储工厂实例
func NewFactory(cfg config.Config) *Factory {
	return &Factory{
		config: cfg,
	}
}

// Disk 按需创建并返回指定磁盘驱动实例
func (f *Factory) Disk() (FileDriver, error) {
	driverName := strings.ToLower(f.config.Driver)
	switch driverName {
	case "oss":
		return driver.NewOss(f.config.Oss)
	case "minio":
		return driver.NewMinio(f.config.Minio)
	default:
		return nil, errors.New("unsupported driver: " + driverName)
	}
}
