package scheduler

import (
	"github.com/xxl-job/xxl-job-executor-go"
	"time"
)

type ExecutorOptions struct {
	ServerAddr   string        `mapstructure:"serverAddr"`
	AccessToken  string        `mapstructure:"accessToken"`
	Timeout      time.Duration `mapstructure:"timeout"`
	ExecutorIp   string        `mapstructure:"ip"`
	ExecutorPort string        `mapstructure:"port"`
	RegistryKey  string        // application name
	logger       xxl.Logger    //日志处理
}

// BuildExecutor 构造执行器
func (ops *ExecutorOptions) BuildExecutor() xxl.Executor {
	opsMap := make([]xxl.Option, 0)
	if ops.ServerAddr != "" {
		opsMap = append(opsMap, xxl.ServerAddr(ops.ServerAddr))
	}
	if ops.AccessToken != "" {
		opsMap = append(opsMap, xxl.AccessToken(ops.AccessToken))
	}
	if ops.ExecutorIp != "" {
		opsMap = append(opsMap, xxl.ExecutorIp(ops.ExecutorIp))
	}
	if ops.ExecutorPort != "" {
		opsMap = append(opsMap, xxl.ExecutorPort(ops.ExecutorPort))
	}
	if ops.RegistryKey != "" {
		opsMap = append(opsMap, xxl.RegistryKey(ops.RegistryKey))
	}
	/*if ops.LogDir!="" {
		opsMap=append(opsMap,xxl.ops.LogDir))
	}*/
	if ops.logger != nil {
		opsMap = append(opsMap, xxl.SetLogger(ops.logger))
	}
	return xxl.NewExecutor(opsMap...)
}
