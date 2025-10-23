package filesystem

import (
	"mime/multipart"
	"testing"

	"github.com/ibiteam/go-core/filesystem/config"
)

func TestPutFile(t *testing.T) {
	c := config.Config{
		Driver: "oss",
		Oss: &config.OssConfig{
			AccessKey:    "test-key",
			AccessSecret: "test-secret",
			Region:       "cn-hangzhou",
			Bucket:       "test-bucket",
		},
	}

	factory := NewFactory(c)
	driver, driverErr := factory.Disk()
	if driverErr != nil {
		t.Fatalf("Failed to create driver: %v", driverErr)
	}

	fileHeader := &multipart.FileHeader{
		Filename: "test.txt",
		Size:     1024,
		Header:   make(map[string][]string),
	}

	// 注意：实际测试应使用有效的文件和配置
	uri, err := driver.PutFile(fileHeader, "test-dir", "test-file.txt")
	if err != nil {
		t.Logf("Expected error due to invalid configuration: %v", err)
		return
	}

	t.Logf("File uploaded successfully: %s", uri)
}
