package apollo

type Option struct {
	AppID     string `mapstructure:"appid"`
	Endpoint  string `mapstructure:"endpoint"`
	Namespace string `mapstructure:"namespace"`
}
