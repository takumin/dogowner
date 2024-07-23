package config

type Config struct {
	LogLevel  string
	LogFormat string
	Reviewdog Reviewdog
}

type Reviewdog struct {
	ConfigPath string
}

func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, o := range opts {
		o.Apply(c)
	}
	return c
}
