package controller

import (
	"fmt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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

	userModel, err := c.GetUserFromContext(ctx)
	if err != nil {
		ec.SetErr(entity.ErrUnauthorized, err)
	} else {
		userEntity := &entity.UserEntity{}
		userEntity.LoadFromModel(userModel)
		ctx.JSON(http.StatusOK, userEntity)
	}
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	ec := NewErrHelper(ctx, c.Name, "GetUserById", c.logger)
	rep := repository.GetUserRepository()

	userModel, err := rep.GetByPublicId(ctx.Param("id"))
	if err != nil {
		ec.SetErr(entity.ErrNotFound, err)
	} else {
		userEntity := &entity.UserEntity{}
		userEntity.LoadFromModel(userModel)
		ctx.JSON(http.StatusOK, userEntity)
		c.logger.Debug(userEntity)
	}
}

//Update current user information
func (c *UserController) PutUser(ctx *gin.Context) {
	ec := NewErrHelper(ctx, c.Name, "PutUser", c.logger)
	rep := repository.GetUserRepository()

	var userUpdateInfo entity.UserUpdateEntity
	err := ctx.BindJSON(&userUpdateInfo)

	if err != nil {
		ec.SetErr(entity.ErrBadRequest, err)
	} else {

		curUser, err := c.GetUserFromContext(ctx)
		if err != nil {
			ec.SetErr(entity.ErrUnauthorized, err)
			return
		}
		var userModel = userUpdateInfo.ToModel()
		userModel.Id = curUser.Id
		c.logger.Debug(fmt.Sprintf("userModel.ID = %d, currentUder.ID = %d", userModel.Id, curUser.Id))
		if rep.UpdateUser(userModel) != nil {
			ec.SetErr(entity.UpdateUserErr, err)
		} else {
			ctx.Status(http.StatusCreated)
			c.logger.Debug(userModel)
		}
	}

}

func GetUserMoc(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		id = "currentUser"
	}
	ctx.JSON(http.StatusOK, entity.CreateUserEntityMoc())
}
