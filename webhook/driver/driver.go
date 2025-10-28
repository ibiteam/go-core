package driver

import (
	"fmt"

	"github.com/ibiteam/go-core/webhook/message"
)

// Driver 定义驱动接口，每个平台需实现此接口
type Driver interface {
	Name() string                                           // 驱动名称（如"feishu"）
	Send(data interface{}) error                            // 发送消息
	SendWithWebhook(data interface{}, webhook string) error // 根据传入的 webhook 发送消息
	convertText(*message.Text) (interface{}, error)         // 转换文本消息为平台格式
}

// ConvertMessage 自动根据消息类型调用对应转换方法（工具包内部使用）
func ConvertMessage(d Driver, msg interface{}) (interface{}, error) {
	switch m := msg.(type) {
	case *message.Text:
		return d.convertText(m)
	default:
		return nil, fmt.Errorf("unsupported message type: %T", msg)
	}
}
