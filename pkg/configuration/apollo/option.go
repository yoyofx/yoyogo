package apollo

type Option struct {
	AppID     string `mapstructure:"appid" config:"appid"`
	Endpoint  string `mapstructure:"endpoint" config:"endpoint"`
	Namespace string `mapstructure:"namespace" config:"namespace"`
}

type RemoteProvider struct {
	provider      string
	endpoint      string
	path          string
	secretKeyring string
}

func DefaultRemoteProvider() *RemoteProvider {
	return &RemoteProvider{provider: "apollo", endpoint: "", path: "", secretKeyring: ""}
}

func (rp RemoteProvider) Provider() string {
	return rp.provider
}

func (rp RemoteProvider) Endpoint() string {
	return rp.endpoint
}

func (rp RemoteProvider) Path() string {
	return rp.path
}

func (rp RemoteProvider) SecretKeyring() string {
	return rp.secretKeyring
}
