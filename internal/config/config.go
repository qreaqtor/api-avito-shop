package config

type Config struct {
	Env  string `yaml:"env"`
	Port uint   `yaml:"app_port" env-required:"true"`

	Postgres PostgresConfig `yaml:"postgres"`
	Auth     AuthConfig     `yaml:"auth"`
}
