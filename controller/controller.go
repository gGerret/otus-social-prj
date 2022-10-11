package controller

import (
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type Controller struct {
	Name string
}

//Базовый фильтр для всех запросов. Идёт первым.
//    * Добвляет случайный идентификатор запроса в контекст запроса для отслеживания
func BaseFilter(ctx *gin.Context) {
	reqId := uuid.Must(uuid.NewV4())
	ctx.Set("reqId", reqId.String())
	ctx.Next()
}

type ApiController struct {
	Controller
	cfg    *ConfigApi
	logger *social.SocialLogger
}

func (c *ApiController) Init(config *ConfigApi, logger *social.SocialLogger) {
	c.logger = logger
	c.cfg = config
}
