package notification

// FeiShuWebhookConfig webhook配置
type FeiShuWebhookConfig struct {
	Uri string
}

// NotifierConfig 协程配置
type NotifierConfig struct {
	Quote   chan string
	Workers int
}

// Config 通知服务配置
type Config struct {
	Driver              string              // 驱动类型
	Notifier            *NotifierConfig     // 协程配置
	FeiShuWebhookConfig FeiShuWebhookConfig // 飞书 webhook 配置
}
