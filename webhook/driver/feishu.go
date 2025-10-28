package driver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ibiteam/go-core/webhook/config"
	"github.com/ibiteam/go-core/webhook/message"
)

// FeiShu 飞书驱动
type FeiShu struct {
	config config.FeiShuConfig
}

// NewFeiShu 创建飞书驱动实例
func NewFeiShu(config config.FeiShuConfig) *FeiShu {
	return &FeiShu{
		config: config,
	}
}

// Name 返回驱动名称
func (f *FeiShu) Name() string {
	return "feishu"
}

// Send 发送消息到飞书
func (f *FeiShu) Send(data interface{}) error {
	return f.SendWithWebhook(data, f.config.WebhookURL)
}

func (f FeiShu) SendWithWebhook(data interface{}, webhook string) error {
	body, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal data failed: %w", err)
	}
	ctx := context.Background()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, webhook, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request failed: %w", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("feishu api error: status code %d", resp.StatusCode)
	}
	return nil
}

// convertText 转换文本消息为飞书格式
func (f *FeiShu) convertText(text *message.Text) (interface{}, error) {
	return map[string]interface{}{
		"msg_type": "text",
		"content":  map[string]string{"text": text.Content},
	}, nil
}
