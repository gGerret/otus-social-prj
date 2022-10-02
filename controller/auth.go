package controller

import (
	"encoding/json"
	"github.com/gGerret/otus-social-prj/controller/auth"
	"github.com/gGerret/otus-social-prj/controller/auth/jwt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
	"net/http"
	"net/http/httputil"
)

const (
	authTypeUser  = "user"
	authTypeGamer = "gamer"
)

type AuthController struct {
	Controller
	cfg    *auth.ConfigAuth
	logger *zap.SugaredLogger
	guard  *jwt.Guard
}

type AuthURIRequest struct {
	Type        string `json:"type"`
	RedirectUri string `json:"redirect_uri"`
}

type AuthURIResponse struct {
	AuthURI string `json:"auth_uri"`
}

type TokenRequest struct {
	Type        string `json:"type"`
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
}

type TokenResponse struct {
	Token     string `json:"token"`
	UserId    string `json:"user_id"`
	IsNewUser bool   `json:"is_new_user"`
}

type ConfiguredAuthResponse struct {
}

func NewAuthController(config *auth.ConfigAuth, logger *zap.SugaredLogger, guard *jwt.Guard) *AuthController {
	c := &AuthController{}
	c.Init(config, logger)
	c.Name = "AuthController"
	c.guard = guard
	return c
}

func (c *AuthController) PostUri(ctx *gin.Context) {
	var a auth.Authenticator
	var authRequest AuthURIRequest
	errHelper := NewErrHelper(ctx, c.Name, "PostUri", c.logger)

	err := ctx.BindJSON(&authRequest)

	if err != nil {
		errHelper.SetErr(entity.ErrInternal, err)
		return
	}

	a, errEnt := auth.GetAuthenticator(authRequest.Type, c.logger)
	if errEnt != nil {
		errHelper.SetErr(*errEnt)
		return
	}

	if authRequest.RedirectUri != "" {
		encoder := json.NewEncoder(ctx.Writer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&AuthURIResponse{
			AuthURI: a.GetAuthURL(authRequest.RedirectUri),
		})
		if err != nil {
			errHelper.SetErr(entity.ErrInternal, err)
		} else {
			ctx.Status(http.StatusOK)
		}
	} else {
		errHelper.SetErr(entity.ErrBadRequest)
	}

}

func (c *AuthController) PostUserToken(ctx *gin.Context) {
	errHelper := NewErrHelper(ctx, c.Name, "PostUserToken", c.logger)
	c.PostToken(ctx, errHelper)
}

func (c *AuthController) PostToken(ctx *gin.Context, errHelper *ErrorHelper) {
	var a auth.Authenticator
	var tokenRequest TokenRequest

	err := ctx.BindJSON(&tokenRequest)
	if err != nil {
		errHelper.SetErr(entity.ErrInternal, err)
		return
	}

	raw, _ := httputil.DumpRequest(ctx.Request, true)
	c.logger.Debugf("PostToken request = %s", string(raw))

	c.logger.Debug(tokenRequest)

	var user *entity.UserEntity

	a, errEnt := auth.GetAuthenticator(tokenRequest.Type, c.logger)
	if errEnt != nil {
		errHelper.SetErr(*errEnt)
		return
	}

	user, err = a.Do(tokenRequest.Code, tokenRequest.RedirectUri)
	if err != nil {
		errHelper.SetErr(entity.ErrInternal, err)
		return
	}

	//Find user in database by VKUserID
	var publicId string
	var authType string

	var isNew = false

	authType = authTypeUser
	rep := repository.GetUserRepository()
	userModel, err := rep.GetByPublicId(user.UserId)
	if err != nil {
		c.logger.Error(err)
	}
	if userModel == nil {
		userModel = user
		err = rep.CreateUserByModel(userModel)
		publicId = userModel.PublicId
		isNew = true
		if err != nil {
			errHelper.SetErr(entity.AuthErrNewUser, err)
			return
		}
	}

	token, err := c.guard.SetToken(publicId, authType)

	if err != nil {
		errHelper.SetErr(entity.AuthErrSetToken, err)
	} else {
		ctx.JSON(http.StatusCreated, &TokenResponse{
			Token:     token,
			UserId:    publicId,
			IsNewUser: isNew,
		})
	}
}
func (c *AuthController) GetToken(ctx *gin.Context) {

	raw, _ := httputil.DumpRequest(ctx.Request, true)
	c.logger.Debugf("GetToken request = %s", string(raw))

	ctx.Status(http.StatusOK)
}

func (c *AuthController) GetConfiguredAuthenticators(ctx *gin.Context) {

}
func (c *AuthController) PostUserPass(ctx *gin.Context) {
	//errHelper := NewErrHelper(ctx, c.Name, "PostUserPass", c.logger)
}
func (c *AuthController) PostUserPassMock(ctx *gin.Context) {
	errHelper := NewErrHelper(ctx, c.Name, "PostUserPassMock", c.logger)

	userModel := &model.UserModel{}
	publicId := uuid.Must(uuid.NewV4())

	userModel.PublicId = publicId.String()
	token, err := c.guard.SetToken(userModel.PublicId, authTypeUser)

	raw, _ := httputil.DumpRequest(ctx.Request, true)
	c.logger.Debugf("PostUserPassMock request = %s", string(raw))

	if err != nil {
		errHelper.SetErr(entity.AuthErrSetToken, err)
	} else {
		ctx.JSON(http.StatusCreated, &TokenResponse{
			Token:  token,
			UserId: userModel.PublicId,
		})
	}
}
