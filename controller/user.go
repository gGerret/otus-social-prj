package controller

import (
	"fmt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/controller/transformer"
	"github.com/gGerret/otus-social-prj/controller/validator"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type UserController struct {
	ApiController
}

func NewUserController(config *ConfigApi, logger *zap.SugaredLogger) *UserController {
	c := &UserController{}
	c.Init(config, logger)
	c.Name = "UserController"
	return c
}

func (c *UserController) GetCurrentUser(ctx *gin.Context) {
	ec := NewErrHelper(ctx, c.Name, "GetCurrentUser", c.logger)

	user, err := c.GetUserFromContext(ctx)
	if err != nil {
		ec.SetErr(entity.ErrUnauthorized, err)
	} else {
		userTransformer := &transformer.UserTransformer{UserModel: user}
		ctx.JSON(http.StatusOK, userTransformer.Transform())
		c.logger.Debug(userTransformer.Transform())
	}
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	ec := NewErrHelper(ctx, c.Name, "GetUserById", c.logger)
	rep := repository.GetUserRepository()

	user, err := rep.GetByPublicId(ctx.Param("id"))
	if err != nil {
		ec.SetErr(entity.ErrNotFound, err)
	} else {
		userTransformer := &transformer.UserPublicTransformer{UserModel: user}
		ctx.JSON(http.StatusOK, userTransformer.Transform())
		c.logger.Debug(userTransformer.Transform())
	}
}

//Update current user information
func (c *UserController) PutUser(ctx *gin.Context) {
	ec := NewErrHelper(ctx, c.Name, "PutUser", c.logger)
	rep := repository.GetUserRepository()

	var userInfo entity.UserUpdateEntity
	err := ctx.BindJSON(&userInfo)

	if err != nil {
		ec.SetErr(entity.ErrBadRequest, err)
	} else {

		curUser, err := c.GetUserFromContext(ctx)
		if err != nil {
			ec.SetErr(entity.ErrUnauthorized, err)
			return
		}
		trans := &transformer.UserUpdateTransformer{Entity: &userInfo}
		var userModel = trans.Transform().(*model.UserModel)
		userModel.ID = curUser.ID
		c.logger.Debug(fmt.Sprintf("userModel.ID = %d, currentUder.ID = %d", userModel.Id, curUser.Id))
		if rep.UpdateUser(userModel) != nil {
			ec.SetErr(entity.UpdateUserErr, err)
		} else {
			ctx.Status(http.StatusCreated)
			c.logger.Debug(userModel)
		}
	}

}

func (c *UserController) GetUserMoc(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		id = "currentUser"
	}
	ctx.JSON(http.StatusOK, entity.UserEntity{
		UserId:     id,
		FirstName:  "Михаил",
		LastName:   "Ушаков",
		MiddleName: "Николаевич",
		Town:       "Рязань",
		Gender:     "мужской",
		Interests:  []string{"Автомобили", "Рисование", "Программирование"},
		CreatedAt:  time.Now().AddDate(0, -1, 0),
		UpdatedAt:  time.Now().AddDate(0, 0, -11),
	})
}
