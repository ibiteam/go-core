package driver

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type FeiShuWebhook struct {
	Uri        string
	HttpClient *http.Client
}

func NewFeiShuWebhook(uri string) *FeiShuWebhook {
	return &FeiShuWebhook{
		Uri: uri,
		HttpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (w *FeiShuWebhook) SendMessage(notification string) error {
	// 构造请求体
	payload := map[string]interface{}{
		"msg_type": "text",
		"content": map[string]string{
			"text": "通知：" + notification,
		},
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	// 构建 http request
	req, err := http.NewRequest("POST", w.Uri, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	resp, err := w.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return nil
}
