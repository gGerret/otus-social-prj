package auth

import "github.com/gGerret/otus-social-prj/controller/auth/jwt"

type ConfigAuth struct {
	AuthUrl string           `json:"authUrl"`
	Guard   *jwt.ConfigGuard `json:"guard"`
}
