package controller

import (
	"errors"
	"fmt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-gonic/gin"
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
		localLogger.Debug(fmt.Sprintf("userModel.ID = %d, currentUder.ID = %d", userModel.Id, curUser.Id))
		err = rep.UpdateUser(userModel)
		if err != nil {
			ec.SetErr(entity.UpdateUserErr, err)
		} else {
			ctx.Status(http.StatusCreated)
			localLogger.Debug(userModel)
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
