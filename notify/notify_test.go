package notify

import (
	"testing"

	"github.com/ibiteam/go-core/notify/config"
	"github.com/ibiteam/go-core/notify/driver"
)

func TestSendText(t *testing.T) {
	r := New()

	feishuConfig := config.FeiShuConfig{
		WebhookURL: "",
	}
	r.RegisterDriver(driver.NewFeiShu(feishuConfig))
	_ = r.SendText("新会话通知,请及时接待\n司机姓名:   -\n司机号码：18837159416\n会话时间：2025-10-27 20:05:07")
}

func TestSendTestWithWorkWechat(t *testing.T) {
	r := New()
	workWechatConfig := config.WorkWechatConfig{
		WebhookURL: "",
	}
	r.RegisterDriver(driver.NewWorkWechat(workWechatConfig))
	_ = r.SendText("新会话通知,请及时接待\n司机姓名:   -\n司机号码：18837159416\n会话时间：2025-10-27 20:05:07")
}
