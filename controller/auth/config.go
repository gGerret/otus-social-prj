package auth

import "github.com/gGerret/otus-social-prj/controller/auth/jwt"

type ConfigVk struct {
	Name         string `json:"name"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	AuthURL      string `json:"authUrl"`
	TokenURL     string `json:"tokenUrl"`
	ApiURL       string `json:"apiUrl"`
}

type ConfigGoogle struct {
	Name         string `json:"name"`
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	AuthURL      string `json:"authUrl"`
	TokenURL     string `json:"tokenUrl"`
	ApiURL       string `json:"apiUrl"`
}

type ConfigAuth struct {
	AuthUrl string           `json:"authUrl"`
	Guard   *jwt.ConfigGuard `json:"guard"`
	Vk      *ConfigVk        `json:"vk"`
	Google  *ConfigGoogle    `json:"google"`
}
