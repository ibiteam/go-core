package webhook

import (
	"testing"

	"github.com/ibiteam/go-core/webhook/config"
	"github.com/ibiteam/go-core/webhook/driver"
)

func TestSendText(t *testing.T) {
	r := New()

	feishuConfig := config.FeiShuConfig{
		WebhookURL: "https://open.feishu.cn/open-apis/bot/v2/hook/e3f1cce4-947a-4abf-8bc9-ee8efbbce710",
	}
	r.RegisterDriver(driver.NewFeiShu(feishuConfig))
	_ = r.SendText("新会话通知,请及时接待\n司机姓名:   -\n司机号码：18837159416\n会话时间：2025-10-27 20:05:07")
}
