package transport

type Config struct {
	Server          string `yaml:"server" json:"server"`
	Port            string `yaml:"port" json:"port"`
	IntervalMs      int64  `yaml:"interval_ms" json:"interval_ms"`
	HeartbeatApi    string `yaml:"heartbeat_api" json:"heartbeat_api"`
	ClientIp        string `yaml:"client_ip" json:"client_ip"`
	AppName         string `yaml:"app_name" json:"app_name"`
	AppType         string `yaml:"app_type" json:"app_type"`
	Hostname        string `yaml:"hostname" json:"hostname"`
	SentinelVersion string `yaml:"sentinel_version" json:"sentinel_version"`
}
