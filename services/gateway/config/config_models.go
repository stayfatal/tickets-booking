package config

type Config struct {
	Server   *ServerCfg
	Services map[string]ServiceInfo
}

type ServerCfg struct {
	Port string
}

type ServiceInfo struct {
	Host string
	Port string
}
