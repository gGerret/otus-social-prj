package repository

type ConfigDb struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Hostname string `json:"hostname"`
	Port     uint16 `json:"port"`
	Net      string `json:"net"`
	SslMode  bool   `json:"ssl_mode"`
}
