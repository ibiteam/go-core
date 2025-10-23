package driver

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"github.com/ibiteam/go-core/filesystem/config"
)

type Oss struct {
	config config.OssConfig
}

func NewOss(cfg *config.OssConfig) (*Oss, error) {
	if cfg == nil {
		return nil, errors.New("OSS config is required")
	}
	if cfg.AccessKey == "" || cfg.AccessSecret == "" {
		return nil, errors.New("OSS access key and secret are required")
	}
	return &Oss{config: *cfg}, nil
}

func (o Oss) PutFile(file *multipart.FileHeader, dirPath string, filename string) (string, error) {

	/*
		Go SDK V2 客户端初始化配置说明：

		1. 签名版本：Go SDK V2 默认使用 V4 签名，提供更高的安全性
		2. Region配置：初始化 Client 时，您需要指定阿里云通用 Region ID 作为发起请求地域的标识
		   本示例代码使用华东1（杭州）Region ID：cn-hangzhou
		   如需查询其它 Region ID 请参见：OSS地域和访问域名
		3. Endpoint配置：
		   - 可以通过 Endpoint 参数，自定义服务请求的访问域名
		   - 当不指定时，SDK 默认根据 Region 信息，构造公网访问域名
		   - 例如当 Region 为 'cn-hangzhou' 时，构造出来的访问域名为：'https://oss-cn-hangzhou.aliyuncs.com'
		4. 协议配置：
		   - SDK 构造访问域名时默认采用 HTTPS 协议
		   - 如需采用 HTTP 协议，请在指定域名时指定为 HTTP，例如：'http://oss-cn-hangzhou.aliyuncs.com'
	*/
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(o.config.AccessKey, o.config.AccessSecret)).
		WithRegion(o.config.Region)

	if o.config.Endpoint != "" {
		cfg = cfg.WithEndpoint(o.config.Endpoint)
	}

	client := oss.NewClient(cfg)

	// 打开文件流
	fileStream, openErr := file.Open()
	if openErr != nil {
		return "", errors.New("打开文件失败~")
	}
	defer func(fileStream multipart.File) {
		_ = fileStream.Close()
	}(fileStream)

	newFilePath := fmt.Sprintf("%s/%s", strings.TrimRight(dirPath, "/"), filename)

	uploadRequest := &oss.PutObjectRequest{
		Bucket: oss.Ptr(o.config.Bucket),
		Key:    oss.Ptr(newFilePath),
		Body:   fileStream,
	}
	result, uploadErr := client.PutObject(context.TODO(), uploadRequest)
	if uploadErr != nil {
		return "", errors.New("上传文件失败!")
	}
	if result.StatusCode != http.StatusOK {
		return "", errors.New("上传文件失败！")
	}

	if o.config.Domain != "" {
		return fmt.Sprintf("https://%s/%s", o.config.Domain, newFilePath), nil
	}
	return fmt.Sprintf("https://%s.%s/%s", o.config.Bucket, o.config.Endpoint, newFilePath), nil
}
