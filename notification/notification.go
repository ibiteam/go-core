package notification

import (
	"sync"

	"github.com/ibiteam/go-core/notification/driver"
)

type NotificationInterface interface {
	SendMessage(message string) error
}

var (
	notification NotificationInterface
	once         sync.Once
	notifier     *NotifierConfig
)

// Initialize 初始化通知服务
func Initialize(cfg Config) {
	once.Do(func() {
		notifier = cfg.Notifier
		if notifier != nil {
			// 启动工作协程
			startWorkers()
		}

		switch cfg.Driver {
		case "feishu_webhook":
			notification = driver.NewFeiShuWebhook(cfg.FeiShuWebhookConfig.Uri)
		}
	})
}

func Send(message string) {
	if notification != nil {
		if notifier == nil {
			_ = sendSync(message)
		} else {
			select {
			case notifier.Quote <- message:
			default:
			}
		}
	}
}

// SendSync 同步发送消息
func sendSync(message string) error {
	if notification != nil {
		return notification.SendMessage(message)
	}
	return nil
}

// startWorkers 启动工作协程
func startWorkers() {
	for i := 0; i < notifier.Workers; i++ {
		go func() {
			for message := range notifier.Quote {
				_ = sendSync(message)
			}
		}()
	}
}
