package config

type Config struct {
	Server   *ServerCfg
	Database *DatabaseCfg
	Cache    *CacheCfg
}

type ServerCfg struct {
	Port string
}

type DatabaseCfg struct {
	User     string
	Password string
	DbName   string `yaml:"db_name"`
	SslMode  string `yaml:"ssl_mode"`
	Host     string
	Port     string
}

type CacheCfg struct {
	Host     string
	Port     string
	Password string
}
