package controller

import (
	"fmt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type TestController struct {
	ApiController
}

func NewTestController(config *ConfigApi, logger *zap.SugaredLogger) *TestController {
	c := &TestController{}
	c.Init(config, logger)
	c.Name = "TestController"
	return c
}
func (c *TestController) InitTestDB(ctx *gin.Context) {
	ec := NewErrHelper(ctx, c.Name, "InitTestDB", c.logger)
	userRep := repository.GetUserRepository()

	for i := 0; i < 500; i++ {
		userName := fmt.Sprintf("user_%d", i)
		email := fmt.Sprintf("user%d@email.com", i)
		user := &model.UserModel{Username: userName, Email: email}
		err := userRep.CreateByModel(user)
		if err != nil {
			ec.SetErr(entity.ErrInternal, err)
			return
		}
	}
	ctx.Status(http.StatusCreated)
}
