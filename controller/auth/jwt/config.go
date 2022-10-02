package jwt

type ConfigGuard struct {
	Header         string `json:"header"`
	Secret         string `json:"secret"`
	TokenLifeHours int    `json:"tokenLifeHours"`
}
