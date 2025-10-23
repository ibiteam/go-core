## 日志处理模块

#### 三种使用方式：`console`, `file`, `database`

#### 处理model
```
import (
	"github.com/ibiteam/go-core/logger"
)
logger.model.ErrorLog{}
```

#### 初始化日志模块并记录日志
```
import (
	"github.com/ibiteam/go-core/logger"
)
// 初始化日志模块
logger.Initialize(Config{
    OutputMode:   "console",
    ConsoleConfig: ConsoleConfig{Colorful: true},
    GormConfig:    GormConfig{Db: DB},
    FileConfig: FileConfig{
        Filename:   "./logs/app.log",
        MaxSize:    64,
        MaxBackups: 10,
        MaxAge:     10,
        Compress:   false,
        LocalTime:  true,
    },
})
data := map[string]string{
    "key": "value",
}
logger.Info("info", logger.Stack(), logger.DictFields(data))
```