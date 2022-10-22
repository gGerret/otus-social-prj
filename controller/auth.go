package controller

import (
	"github.com/gGerret/otus-social-prj/controller/auth"
	"github.com/gGerret/otus-social-prj/controller/auth/jwt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/repository"
	"github.com/gGerret/otus-social-prj/repository/model"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gGerret/otus-social-prj/utils"

	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

const (
	authTypeUser = "user"
)

type AuthController struct {
	Controller
	cfg    *auth.ConfigAuth
	logger *social.SocialLogger
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

func NewAuthController(config *auth.ConfigAuth, logger *social.SocialLogger, guard *jwt.Guard) *AuthController {
	c := &AuthController{}
	c.Init(config, logger)
	c.Name = "AuthController"
	c.guard = guard
	return c
}

func (c *AuthController) Init(config *auth.ConfigAuth, logger *social.SocialLogger) {
	c.logger = logger
	c.cfg = config
}

func (c *AuthController) PostUserPass(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "PostUserPass")
	ec := NewErrHelper(ctx, localLogger)

	var userLogin entity.UserBaseLoginEntity
	err := ctx.BindJSON(&userLogin)
	if err != nil {
		ec.SetErr(entity.ErrBadRequest, err)
		return
	}


	userRepo := repository.GetUserRepository()
	user, err := userRepo.GetByEmail(userLogin.Email)
	if err != nil {
		ec.SetErr(entity.AuthErrUserNotFound)
		return
	}

	loginPassHash := utils.GeneratePassHash(userLogin.Password)
	//Сравниваем хеши паролей
	if loginPassHash != user.PasswordHash {
		ec.SetErr(entity.AuthErrUserNotFound)
		return
	}
	token, err := c.guard.SetToken(user.PublicId, authTypeUser)
	if err != nil {
		ec.SetErr(entity.AuthErrSetToken, err)
	} else {
		ctx.JSON(http.StatusCreated, &TokenResponse{
			Token:  token,
			UserId: user.PublicId,
		})
	}
}
func (c *AuthController) PostUserPassMock(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "PostUserPassMock")
	errHelper := NewErrHelper(ctx, localLogger)

	userModel := &model.UserModel{}
	//publicId := uuid.Must(uuid.NewV4())

	userModel.PublicId = "53722bba-4030-453f-8578-dc1d3941069c" //publicId.String()
	token, err := c.guard.SetToken(userModel.PublicId, authTypeUser)

	raw, _ := httputil.DumpRequest(ctx.Request, true)
	localLogger.Debugf("PostUserPassMock request = %s", string(raw))

	if err != nil {
		errHelper.SetErr(entity.AuthErrSetToken, err)
	} else {
		ctx.JSON(http.StatusCreated, &TokenResponse{
			Token:  token,
			UserId: userModel.PublicId,
		})
	}
}
