package notification

import (
	"testing"
	"time"
)

func setup() {
	config := Config{
		Driver: "feishu_webhook",
		FeiShuWebhookConfig: FeiShuWebhookConfig{
			Uri: "https://open.feishu.cn/xxxxxxxx",
		},
		Notifier: &NotifierConfig{
			Quote:   make(chan string, 100),
			Workers: 1,
		},
	}
	Initialize(config)
}

func setupSync() {
	config := Config{
		Driver: "feishu_webhook",
		FeiShuWebhookConfig: FeiShuWebhookConfig{
			Uri: "https://open.feishu.cn/xxxxxxxx",
		},
	}
	Initialize(config)
}

func TestSendMessage(t *testing.T) {
	setup()

	Send("测试异步消息,忽略返回异常")

	time.Sleep(3 * time.Second)
}

func TestSendMessageSync(t *testing.T) {
	setupSync()
	Send("测试同步消息,忽略返回异常~")
}
