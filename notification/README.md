## 通知处理模块
```go
package main

import (
	"github.com/ibiteam/go-core/notification"
)

// 初始化
func setup() {
	config := notification.Config{
		Driver: "feishu_webhook",
		FeiShuWebhookConfig: notification.FeiShuWebhookConfig{
			Uri: "https://open.feishu.cn/xxxxxxxx",
		},
		WorkerCount: 1,
	}
	notification.Initialize(config)

	// 使用
	notification.Send("hello world")
}
```