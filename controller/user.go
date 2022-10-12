package controller

import (
	"errors"
	"fmt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/controller/validator"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gGerret/otus-social-prj/utils"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"time"
)

type UserController struct {
	ApiController
}

func NewUserController(config *ConfigApi, logger *social.SocialLogger) *UserController {
	c := &UserController{}
	c.Init(config, logger)
	c.Name = "UserController"
	return c
}

func (c *UserController) GetUserFromContext(ctx *gin.Context) (*model.UserModel, error) {
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

func (c *UserController) RegisterUser(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "RegisterUser")
	ec := NewErrHelper(ctx, localLogger)

	var newUser entity.UserRegisterEntity
	err := ctx.BindJSON(&newUser)
	if err != nil {
		ec.SetErr(entity.ErrBadRequest, err)
	} else {
		v := validator.UserRegisterValidator{Entity: &newUser}
		fe := v.Validate()
		if len(fe) != 0 {
			ec.SetErrEx(entity.DataErrBadUserInfo, fe)
			return
		}
	}

	rep := repository.GetUserRepository()
	userModel := newUser.ToModel()
	userModel.PasswordHash = utils.GeneratePassHash(newUser.Password)
	userModel.PublicId = uuid.Must(uuid.NewV4()).String()
	userModel.CreatedAt = time.Now()

	err = rep.CreateByModel(userModel)
	if err != nil {
		ec.SetErr(entity.RegisterUserErr, err)
		return
	} else {
		localLogger.Infof("User %s successfully registered", userModel.PublicId)
		localLogger.Debug(userModel)
		ctx.Status(http.StatusOK)
	}
}

func (c *UserController) GetCurrentUser(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "GetCurrentUser")
	ec := NewErrHelper(ctx, localLogger)

	user, err := c.GetUserFromContext(ctx)
	if err != nil {
		ec.SetErr(entity.ErrUnauthorized, err)
	} else {
		userEntity := &entity.UserEntity{}
		userEntity.FromModel(user)
		ctx.JSON(http.StatusOK, userEntity)
		localLogger.Debug(userEntity)
	}
}

func (c *UserController) GetUserById(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "GetUserById")
	ec := NewErrHelper(ctx, localLogger)
	rep := repository.GetUserRepository()

	user, err := rep.GetByPublicId(ctx.Param("id"))
	if err != nil {
		ec.SetErr(entity.ErrNotFound, err)
	} else {
		userEntity := &entity.UserPublicEntity{}
		userEntity.FromModel(user)
		ctx.JSON(http.StatusOK, userEntity)
	}
}

//Update current user information
func (c *UserController) PutUser(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "PutUser")
	ec := NewErrHelper(ctx, localLogger)
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

		var userModel = userInfo.ToModel()
		userModel.Id = curUser.Id
		localLogger.Debug(fmt.Sprintf("userModel.Id = %d, currentUder.Id = %d", userModel.Id, curUser.Id))
		err = rep.UpdateUser(userModel)
		if err != nil {
			ec.SetErr(entity.UpdateUserErr, err)
		} else {
			ctx.Status(http.StatusCreated)
			localLogger.Debug(userModel)
		}
	}

}

func (c *UserController) GetUserByFilter(context *gin.Context) {

}

func (c *UserController) GetUserFriends(context *gin.Context) {

}

func (c *UserController) MakeFriendship(context *gin.Context) {

}
