package controller

type ConfigApi struct {
	ApiURL      string `json:"apiUrl"`
	Version     string `json:"version"`
	TopPageSize int    `json:"topPageSize"`
	AdminToken  string `json:"adminToken"`
}
