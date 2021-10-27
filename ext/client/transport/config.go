package transport

type Config struct {
	Server       string `yaml:"server" json:"server"`
	Port         string `yaml:"port" json:"port"`
	IntervalMs   int64  `yaml:"interval_ms" json:"interval_ms"`
	ClientIp     string `yaml:"client_ip" json:"client_ip"`
	HeartbeatApi string `yaml:"heartbeat_api" json:"heartbeat_api"`
}




