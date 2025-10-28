package webhook

import (
	"fmt"
	"sync"

	"github.com/ibiteam/go-core/webhook/driver"
	"github.com/ibiteam/go-core/webhook/message"
)

// Notifier 通知器，用户的主要交互对象
type Notifier struct {
	drivers map[string]driver.Driver // 注册的驱动
	mu      sync.RWMutex             // 添加读写锁保护
}

// New 创建通知器实例
func New() *Notifier {
	return &Notifier{
		drivers: make(map[string]driver.Driver),
	}
}

// RegisterDriver 注册一个或多个驱动（如飞书、企业微信）
func (n *Notifier) RegisterDriver(drivers ...driver.Driver) {
	n.mu.Lock()
	defer n.mu.Unlock()
	for _, d := range drivers {
		n.drivers[d.Name()] = d
	}
}

// SendText 发送文本消息
// content: 文本内容
func (n *Notifier) SendText(content string) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.send(message.NewText(content))
}

func (n *Notifier) SendTextWithWebhook(content string, webhook string) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.sendWithWebhook(message.NewText(content), webhook)
}

// 内部发送逻辑：转换消息格式并调用驱动发送
func (n *Notifier) send(msg interface{}) error {
	for _, d := range n.drivers {
		// 转换消息为驱动所需格式
		convertedMsg, err := driver.ConvertMessage(d, msg)
		if err != nil {
			return fmt.Errorf("convert message failed: %w", err)
		}
		// 调用驱动发送
		if sendErr := d.Send(convertedMsg); sendErr != nil {
			return fmt.Errorf("send failed: %w", sendErr)
		}
	}
	return nil
}

func (n *Notifier) sendWithWebhook(msg interface{}, webhook string) error {
	for _, d := range n.drivers {
		// 转换消息为驱动所需格式
		convertedMsg, err := driver.ConvertMessage(d, msg)
		if err != nil {
			return fmt.Errorf("convert message failed: %w", err)
		}
		// 调用驱动发送
		if sendErr := d.SendWithWebhook(convertedMsg, webhook); sendErr != nil {
			return fmt.Errorf("send failed: %w", sendErr)
		}
	}
	return nil
}
