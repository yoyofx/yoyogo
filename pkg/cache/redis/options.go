package redis

type Options struct {
	Addr     string
	Addrs    []string
	Password string
	DB       int
}
