package controller

import (
	"errors"
	"github.com/gGerret/otus-social-prj/controller/auth"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Controller struct {
	Name string
}

type ApiController struct {
	Controller
	cfg    *ConfigApi
	logger *zap.SugaredLogger
}

func (c *ApiController) Init(config *ConfigApi, logger *zap.SugaredLogger) {
	c.logger = logger
	c.cfg = config
}

func (c *AuthController) Init(config *auth.ConfigAuth, logger *zap.SugaredLogger) {
	c.logger = logger
	c.cfg = config
}

func (c *ApiController) GetUserFromContext(ctx *gin.Context) (*model.UserModel, error) {
	u, exists := ctx.Get("User")

	if !exists || u == nil {
		return nil, errors.New("there is no user information in context")
	}

	switch u.(type) {
	case *model.UserModel:
		return u.(*model.UserModel), nil
	default:
		return nil, errors.New("unknown type in context instead of UserModel")
	}

}
