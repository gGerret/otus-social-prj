package controller

import (
	"github.com/gGerret/otus-social-prj/controller/auth"
	"github.com/gGerret/otus-social-prj/controller/auth/jwt"
	"github.com/gGerret/otus-social-prj/controller/entity"
	"github.com/gGerret/otus-social-prj/social"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
	"net/http/httputil"
	"uyg-backend/models"
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
	//errHelper := NewErrHelper(ctx, c.Name, "PostUserPass", c.logger)
}
func (c *AuthController) PostUserPassMock(ctx *gin.Context) {
	localLogger := c.logger.ContextLogger(ctx.GetString("reqId"), "PostUserPassMock")
	errHelper := NewErrHelper(ctx, localLogger)

	userModel := &models.UserModel{}
	publicId := uuid.Must(uuid.NewV4())

	userModel.PublicId = publicId.String()
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
