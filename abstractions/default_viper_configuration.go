package abstractions

import (
	"flag"
	"github.com/jinzhu/copier"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/utils"
	"path"
	"strings"
	"sync"
)

type Configuration struct {
	context   *ConfigurationContext
	configMap map[string]interface{}
	gRWLock   *sync.RWMutex
	config    *viper.Viper
	log       xlog.ILogger
}

func NewConfiguration(configContext *ConfigurationContext) *Configuration {
	log := xlog.GetXLogger("Configuration")
	log.SetCustomLogFormat(nil)

	defaultConfig := viper.New()
	if configContext.enableEnv {
		defaultConfig.AutomaticEnv()
		defaultConfig.SetEnvPrefix("YYG")
	}
	if configContext.enableFlag {
		flag.String("app", "", "application name")
		flag.String("port", "", "application port")
		flag.String("profile", configContext.profile, "application profile")
		flag.String("f", "", "config file path")
		flag.String("conf", ".", "config dir")
		pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
		pflag.Parse()
		_ = defaultConfig.BindPFlags(pflag.CommandLine)
	}

	if pf := defaultConfig.GetString("profile"); pf != "" {
		configContext.profile = pf
	}

	if cf := defaultConfig.GetString("conf"); cf != "" {
		configContext.configDir = cf
	}

	if configFile := defaultConfig.GetString("f"); configFile != "" {
		configContext.configFile = configFile
	}
	configFilePath := configContext.configFile
	if configFilePath == "" {
		configName := configContext.configName + "_" + configContext.profile
		configFilePath = path.Join(configContext.configDir, configName+"."+configContext.ConfigType)
		exists, _ := utils.PathExists(configFilePath)
		if !exists {
			configName = configContext.configName
		}
		defaultConfig.AddConfigPath(configContext.configDir)
		defaultConfig.SetConfigName(configName)
		defaultConfig.SetConfigType(configContext.ConfigType)
	} else {
		defaultConfig.SetConfigFile(configFilePath)
	}

	if err := defaultConfig.ReadInConfig(); err != nil {
		panic(err)
		return nil
	}
	log.Debug(configFilePath)

	configuration := &Configuration{
		context:   configContext,
		config:    defaultConfig,
		gRWLock:   new(sync.RWMutex),
		configMap: make(map[string]interface{}),
		log:       log,
	}

	if configContext.EnableRemote {
		defaultConfig = configContext.RemoteProvider.GetProvider(defaultConfig)
		if defaultConfig.ConfigFileUsed() == "" {
			_ = defaultConfig.BindPFlags(pflag.CommandLine)
			//remote config
			configuration.config = defaultConfig
			configuration.OnWatchRemoteConfigChanged()
			log.Info("remote config is ready , on changed notify listening ......")
		} else {
			log.Error("remote config is not ready , switch local.")
		}
	}

	return configuration
}

func (c *Configuration) OnWatchRemoteConfigChanged() {
	respChan := c.context.RemoteProvider.WatchRemoteConfigOnChannel(c.config)
	go func(rc <-chan bool) {
		for {
			<-rc
			c.RefreshAll()
			c.log.Info("sync remote config")
		}
	}(respChan)
}

func (c *Configuration) Get(name string) interface{} {
	if c.AssertEnvironment(c.config.GetString(name)) {
		c.BindEnvironment(c.config.GetString(name), name)
	}
	return c.config.Get(name)
}

func (c *Configuration) GetString(name string) string {
	if c.AssertEnvironment(c.config.GetString(name)) {
		c.BindEnvironment(c.config.GetString(name), name)
	}
	return c.config.GetString(name)
}

func (c *Configuration) GetBool(name string) bool {
	if c.AssertEnvironment(c.config.GetString(name)) {
		c.BindEnvironment(c.config.GetString(name), name)
	}
	return c.config.GetBool(name)
}

func (c *Configuration) GetInt(name string) int {
	if c.AssertEnvironment(c.config.GetString(name)) {
		c.BindEnvironment(c.config.GetString(name), name)
	}
	return c.config.GetInt(name)
}

func (c *Configuration) GetSection(name string) IConfiguration {
	section := c.config.Sub(name)

	if section != nil {
		return &Configuration{config: section}
	}
	return nil
}

func (c *Configuration) Unmarshal(obj interface{}) {
	err := c.config.Unmarshal(obj)
	if err != nil {
		c.log.Error("unmarshal config is failed, err:", err)
	}
}

func (c *Configuration) GetProfile() string {
	return c.context.profile
}

func (c *Configuration) GetConfDir() string {
	return c.context.configDir
}

func (c *Configuration) GetConfigObject(configTag string, configObject interface{}) {
	c.gRWLock.RLock()
	object := c.configMap[configTag]
	c.gRWLock.RUnlock()
	if object == nil {
		// need lock
		Section := c.GetSection(configTag)
		Section.Unmarshal(configObject)
		object = configObject
		c.gRWLock.Lock()
		c.configMap[configTag] = object
		c.gRWLock.Unlock()
	} else {
		_ = copier.Copy(configObject, object)
	}

}

func (c *Configuration) RefreshAll() {
	c.gRWLock.Lock()
	c.configMap = make(map[string]interface{})
	c.gRWLock.Unlock()
}

func (c *Configuration) RefreshBy(name string) {
	c.gRWLock.Lock()
	delete(c.configMap, name)
	c.gRWLock.Unlock()
}

/**
读取环境变量
*/
func (c *Configuration) BindEnvironment(key string, originalKey string) {
	//fmt.Println("两个参数")
	//fmt.Println(key + "---" + originalKey)
	key = key[2 : len(key)-1]
	envKeyDefaultValue := strings.Split(key, ":")
	if len(envKeyDefaultValue) > 2 {
		panic("can't read environment for illegal key:" + key)
	}
	envKey := envKeyDefaultValue[0]
	//fmt.Println("环境变量key" + envKey)
	viper.BindEnv(envKey)
	envValue := viper.Get(envKey)
	if envValue == nil {
		if len(envKeyDefaultValue) > 1 {
			c.config.Set(originalKey, envKeyDefaultValue[1])
			return
		}
	}
	c.config.Set(originalKey, envValue)
}

func (c *Configuration) AssertEnvironment(key string) bool {
	if len(key) < 2 {
		return false
	}
	prefix := key[0:2]
	last := key[len(key)-1:]
	if !(prefix == "${" && last == "}") {
		return false
	}
	return true
}
