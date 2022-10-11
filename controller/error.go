package controller

import (
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-gonic/gin"
)

// TODO MultiLang
type ErrorHelper struct {
	ctx    *gin.Context
	logger *social.SocialLogger
}

func NewErrHelper(ct *gin.Context, logger *social.SocialLogger) *ErrorHelper {
	return &ErrorHelper{
		ctx:    ct,
		logger: logger,
	}
}

func (c *ErrorHelper) SetErr(entity entity.ErrorEntity, err ...error) {

	if len(err) > 0 && err[0] != nil {
		entity.Description += " Caused by: " + err[0].Error() + "\n"
		c.logger.Error(entity.Message + entity.Description)
	} else {
		c.logger.Error(entity.Message)
	}
	c.ctx.JSON(entity.HttpCode, entity)
}

func (c *ErrorHelper) SetErrEx(entity entity.ErrorEntityEx, data interface{}) {
	if data != nil {
		c.logger.Error(entity.Message, data)
		entity.Errors = data
	} else {
		c.logger.Error(entity.Message)
	}
	c.ctx.JSON(entity.HttpCode, entity)
}
