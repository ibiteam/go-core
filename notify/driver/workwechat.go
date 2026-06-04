package driver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ibiteam/go-core/notify/config"
	"github.com/ibiteam/go-core/notify/message"
)

// WorkWechat 企业微信驱动
type WorkWechat struct {
	config config.WorkWechatConfig
}

// NewWorkWechat 创建企业微信驱动实例
func NewWorkWechat(config config.WorkWechatConfig) *WorkWechat {
	return &WorkWechat{
		config: config,
	}
}

// Name 返回驱动名称
func (w *WorkWechat) Name() string {
	return "workwechat"
}

// Send 发送消息到企业微信
func (w *WorkWechat) Send(data interface{}) error {
	return w.SendWithWebhook(data, w.config.WebhookURL)
}

func (w WorkWechat) SendWithWebhook(data interface{}, webhook string) error {
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
		return fmt.Errorf("workwechat api error: status code %d", resp.StatusCode)
	}
	return nil
}

// convertText 转换文本消息为企业微信格式
func (w *WorkWechat) convertText(text *message.Text) (interface{}, error) {
	return map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]string{"content": text.Content},
	}, nil
}
