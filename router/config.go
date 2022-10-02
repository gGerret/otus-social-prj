package router

import (
	"github.com/gGerret/otus-social-prj/controller"
	"github.com/gGerret/otus-social-prj/controller/auth"
)

type ServerConfig struct {
	BaseURL string                `json:"baseUrl"`
	Port    int                   `json:"port"`
	Api     *controller.ConfigApi `json:"api"`
	Auth    *auth.ConfigAuth      `json:"auth"`
}
