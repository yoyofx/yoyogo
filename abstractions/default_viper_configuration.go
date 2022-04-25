package abstractions

import (
	"errors"
	"flag"
	"github.com/jinzhu/copier"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/utils"
	"path"
	"reflect"
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
	configuration.Initialize()
	return configuration
}

func (c *Configuration) Initialize() {
	c.context.decoderConfigOption = func(config *mapstructure.DecoderConfig) {
		config.TagName = "config"
		config.DecodeHook = func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			if f.Kind() != reflect.String && t.Kind() != reflect.String {
				return data, nil
			}
			fromStr := data.(string)
			if c.assertDSL(fromStr) {
				return c.bindEnvDSL(fromStr, "")
			}
			return data, nil
		}
	}
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
	if c.assertDSL(c.config.GetString(name)) {
		_, _ = c.bindEnvDSL(c.config.GetString(name), name)
	}
	return c.config.Get(name)
}

func (c *Configuration) GetString(name string) string {
	if c.assertDSL(c.config.GetString(name)) {
		_, _ = c.bindEnvDSL(c.config.GetString(name), name)
	}
	return c.config.GetString(name)
}

func (c *Configuration) GetBool(name string) bool {
	if c.assertDSL(c.config.GetString(name)) {
		_, _ = c.bindEnvDSL(c.config.GetString(name), name)
	}
	return c.config.GetBool(name)
}

func (c *Configuration) GetInt(name string) int {
	if c.assertDSL(c.config.GetString(name)) {
		_, _ = c.bindEnvDSL(c.config.GetString(name), name)
	}
	return c.config.GetInt(name)
}

func (c *Configuration) GetSection(name string) IConfiguration {
	section := c.config.Sub(name)
	if section != nil {
		return &Configuration{config: section, log: c.log, context: c.context, configMap: c.configMap}
	}
	return nil
}

func (c *Configuration) Unmarshal(obj interface{}) {
	err := c.config.Unmarshal(obj, c.context.decoderConfigOption)
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
		_ = copier.CopyWithOption(configObject, object, copier.Option{IgnoreEmpty: true, DeepCopy: true})
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
bindEnvDSL 读取DSL:环境变量 -> ${ENV:DEFAULT}
*/
func (c *Configuration) bindEnvDSL(dslKey string, originalKey string) (interface{}, error) {
	dslKey = dslKey[2 : len(dslKey)-1]
	envKeyDefaultValue := strings.Split(dslKey, ":")
	if len(envKeyDefaultValue) > 2 {
		return nil, errors.New("can't read environment for illegal key:" + dslKey)
	}
	envKey := envKeyDefaultValue[0] // ${ENV:DEFAULT}  [0] is key of the env
	c.gRWLock.Lock()
	defer c.gRWLock.Unlock()
	// written lock
	_ = viper.BindEnv(envKey)     // binding environment for the env key
	envValue := viper.Get(envKey) // get value of env
	dslValue := envValue          // dsl value is env value by default
	// if env value is nil and dsl parameters length > 1
	if envValue == nil && len(envKeyDefaultValue) > 1 {
		dslValue = envKeyDefaultValue[1]
	}
	if originalKey != "" {
		c.config.Set(originalKey, dslValue)
	}
	return dslValue, nil
}

func (c *Configuration) assertDSL(key string) bool {
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
