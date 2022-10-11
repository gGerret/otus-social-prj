package controller

import (
	"fmt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TestController struct {
	ApiController
}

func NewTestController(config *ConfigApi, logger *social.SocialLogger) *TestController {
	c := &TestController{}
	c.Init(config, logger)
	c.Name = "TestController"
	return c
}
func (c *TestController) InitTestDB(ctx *gin.Context) {
	ec := NewErrHelper(ctx, c.logger)
	userRep := repository.GetUserRepository()

	for i := 0; i < 500; i++ {
		userName := fmt.Sprintf("user_%d", i)
		email := fmt.Sprintf("user%d@email.com", i)
		user := &model.UserModel{FirstName: userName, LastName: email}
		err := userRep.CreateByModel(user)
		if err != nil {
			ec.SetErr(entity.ErrInternal, err)
			return
		}
	}
	ctx.Status(http.StatusCreated)
}
