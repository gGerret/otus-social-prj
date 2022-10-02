package controller

import (
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TODO MultiLang
type ErrorHelper struct {
	ctx        *gin.Context
	ctrlName   string
	actionName string
	logger     *zap.SugaredLogger
}

func NewErrHelper(ct *gin.Context, controllerName, actionName string, logger *zap.SugaredLogger) *ErrorHelper {
	return &ErrorHelper{
		ctx:        ct,
		ctrlName:   controllerName,
		actionName: actionName,
		logger:     logger,
	}
}

func (c *ErrorHelper) SetErr(entity entity.ErrorEntity, err ...error) {

	if len(err) > 0 && err[0] != nil {
		entity.Description += " Caused by: " + err[0].Error() + "\n"
		c.logger.Error("[" + c.ctrlName + "." + c.actionName + "] " + entity.Message + entity.Description)
	} else {
		c.logger.Error("[" + c.ctrlName + "." + c.actionName + "] " + entity.Message)
	}
	c.ctx.JSON(entity.HttpCode, entity)
}

func (c *ErrorHelper) SetErrEx(entity entity.ErrorEntityEx, data interface{}) {
	if data != nil {
		c.logger.Error("["+c.ctrlName+"."+c.actionName+"] "+entity.Message, data)
		entity.Errors = data
	} else {
		c.logger.Error("[" + c.ctrlName + "." + c.actionName + "] " + entity.Message)
	}
	c.ctx.JSON(entity.HttpCode, entity)
}
